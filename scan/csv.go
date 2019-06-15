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
			headerColumn[k] = "setup_time"
		case "source_customer_free_time":
			headerColumn[k] = "end_time"
		case "source_reseller_free_time":
			headerColumn[k] = "dialog_time"

		case "destination_carrier_free_time":
			headerColumn[k] = "direction"
		case "destination_customer_free_time":
			headerColumn[k] = "prefix"
		case "destination_reseller_free_time":
			headerColumn[k] = "trunk"
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
		initTime  int
		startTime int
		duration  int
		direction string
		prefixNum string
		prefix    string
		incTrunk  string
		outTrunk  string
		setupTime int
	)
	for k := range headerColumn {
		switch headerColumn[k] {
		case "update_time":
			t, err := time.Parse("2006-01-02 15:04:05", column[k])
			if err != nil {
				logp.Err("%v", err)
				continue
			}
			column[k] = strconv.Itoa(int(t.Unix()))

		case "source_user_id":
			if column[k] == "0" {
				direction = "inc"
			}

		case "source_domain":
			incTrunk = column[k]

		case "destination_user_id":
			if column[k] == "0" {
				direction = "out"
			}

		case "destination_domain":
			outTrunk = column[k]

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
			column[k] = strconv.Itoa(initTime)

		case "time":
			t, err := time.Parse(c.timeFormat, column[k])
			if err != nil {
				logp.Err("%v", err)
				continue
			}
			startTime = int(t.Unix())
			column[k] = strconv.Itoa(startTime)

		case "duration":
			f, err := strconv.ParseFloat(column[k], 32)
			if err != nil {
				logp.Err("%v", err)
				continue
			}
			duration = int(math.Round(f))
			column[k] = strconv.Itoa(duration)

		case "rated_at":
			t, err := time.Parse("2006-01-02 15:04:05", column[k])
			if err != nil {
				logp.Err("%v", err)
				continue
			}
			column[k] = strconv.Itoa(int(t.Unix()))

		case "setup_time":
			setupTime = startTime - initTime
			column[k] = strconv.Itoa(setupTime)

		case "end_time":
			column[k] = strconv.Itoa(startTime + duration)

		case "dialog_time":
			column[k] = strconv.Itoa(setupTime + duration)

		case "prefix":
			column[k] = ""
			if _, err := strconv.ParseUint(prefix, 0, 64); err != nil {
				column[k] = prefix
			}

		case "direction":
			column[k] = direction

		case "trunk":
			column[k] = incTrunk
			if direction == "out" {
				column[k] = outTrunk
			}
		}
	}

	return column
}

// End will close the watcher.
func (c *CSVInput) End() {
	close(c.shut)
}
