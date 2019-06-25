package scan

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/negbie/logp"
)

// formHeader will do some header manipulations.
func (c *CSVInput) formHeader(headerColumn []string) ([]string, error) {
	var hasTs bool
	for k := range headerColumn {
		// Rename following headers.
		switch headerColumn[k] {
		case c.timeColumn:
			// Check if we found our configured timestamp column
			// and rename it to time
			hasTs = true
			headerColumn[k] = "time"

		case "source_carrier_free_time":
			headerColumn[k] = "setup_duration"
		case "source_customer_free_time":
			headerColumn[k] = "end_time"
		case "source_reseller_free_time":
			headerColumn[k] = "dialog_duration"

		case "destination_carrier_free_time":
			headerColumn[k] = "direction"
		case "destination_customer_free_time":
			headerColumn[k] = "prefix"
		case "destination_reseller_free_time":
			headerColumn[k] = "owner"
		}
	}

	// CSV file has no configured timestamp column
	if !hasTs {
		return nil, fmt.Errorf("CSVTimeColumn %s does not match any header", c.timeColumn)
	}

	return headerColumn, nil
}

// formRecords will do some records manipulations.
func (c *CSVInput) formRecords(headerColumn, column []string) []string {
	var (
		initTime      int
		startTime     time.Time
		startTimeUnix int
		duration      int
		srcID         string
		dstID         string
		prefixNum     string
		prefix        string
		setupTime     int
	)
	for k := range headerColumn {
		switch headerColumn[k] {
		case "source_user_id":
			srcID = column[k]

		case "destination_user_id":
			dstID = column[k]

		case "destination_user":
			prefixNum = column[k]

		case "destination_user_in":
			prefix = strings.TrimSuffix(prefixNum, column[k])

		case "init_time":
			t, err := time.Parse(c.timeFormat, column[k])
			if err != nil {
				logp.Err("%v", err)
				continue
			}
			initTime = int(t.Unix())
			//column[k] = strconv.Itoa(initTime)

		case "time":
			t, err := time.Parse(c.timeFormat, column[k])
			if err != nil {
				logp.Err("%v", err)
				continue
			}
			startTime = t
			startTimeUnix = int(t.Unix())
			//column[k] = strconv.Itoa(startTime)

		case "duration":
			f, err := strconv.ParseFloat(column[k], 32)
			if err != nil {
				logp.Err("%v", err)
				continue
			}
			duration = int(math.Round(f))
			column[k] = strconv.Itoa(duration)

		case "setup_duration":
			setupTime = startTimeUnix - initTime
			column[k] = strconv.Itoa(setupTime)

		case "end_time":
			//column[k] = strconv.Itoa(startTimeUnix + duration)
			column[k] = startTime.Add(time.Duration(duration) * time.Second).Format(c.timeFormat)

		case "dialog_duration":
			column[k] = strconv.Itoa(setupTime + duration)

		case "prefix":
			column[k] = ""
			if _, err := strconv.ParseUint(prefix, 0, 64); err != nil {
				column[k] = prefix
			}

		case "direction":
			if srcID == "0" {
				column[k] = "inc"
			}
			if dstID == "0" {
				column[k] = "out"
			}
			if srcID == "0" && dstID == "0" {
				column[k] = "tra"
			}
			if srcID != "0" && dstID != "0" {
				column[k] = "int"
			}

		case "owner":
			column[k] = c.owner
		}
	}

	return column
}

// End will close the watcher.
func (c *CSVInput) End() {
	close(c.shut)
}
