package scan

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/negbie/logp"
	watch "github.com/negbie/ngcp-cdr-db"
	"github.com/negbie/ngcp-cdr-db/config"
	"github.com/negbie/ngcp-cdr-db/database"
)

// CSVInput struct
type CSVInput struct {
	csvChan    chan []string
	shut       chan struct{}
	batchSize  int
	timeColumn string
	timeFormat string
}

// NewCSV creates a new CSVInput object
func NewCSV() *CSVInput {
	return &CSVInput{
		csvChan:    make(chan []string, config.Setting.CSVQueueSize),
		shut:       make(chan struct{}),
		batchSize:  config.Setting.CSVBatchSize,
		timeColumn: config.Setting.CSVTimeColumn,
		timeFormat: config.Setting.CSVTimeFormat,
	}
}

// Run will start the watcher and scan the CSV file
func (c *CSVInput) Run() {
	if !config.Setting.DryRun {
		go func() {
			d := database.New("CSVImport")
			d.Chan = c.csvChan
			if err := d.Run(); err != nil {
				logp.Critical("%v", err)
			}
		}()
	}

	// scan reads a CSV file line by line, does some manipulations
	// and sends the result to the database channel
	scan := func(filePath string) {
		var (
			err       error
			csvFile   *os.File
			reader    *csv.Reader
			linesRead int64    // Line counter
			headers   []string // Header column array
			header    string   // Full comma seperated header line
		)

		scanStart := time.Now()

		// Open CSV file.
		csvFile, err = os.Open(filePath)
		if err != nil {
			logp.Err("%v", err)
			return
		}
		defer csvFile.Close()

		in, err := ioutil.ReadFile(filePath)
		if err != nil {
			logp.Err("%v", err)
			return
		}

		out := bytes.Replace(in, []byte("'"), []byte(""), -1)
		b := bytes.NewReader(out)

		//reader = csv.NewReader(bufio.NewReader(b))
		reader = csv.NewReader(b)
		var rec [][]string
		for {
			line, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				if err, ok := err.(*csv.ParseError); !ok || err.Err != csv.ErrFieldCount {
					logp.Err("%v", err)
					return
				}
			}
			rec = append(rec, line)
		}
		rawHeader := strings.Split(cutSpace(config.Setting.CSVHeader), ",")
		rows := make([]string, 1, c.batchSize)

		for i := 0; i < len(rec); i++ {
			if i == 0 {
				headers, err = c.formHeader(rawHeader)
				if err != nil {
					logp.Err("%v", err)
					break
				}
				header = strings.Join(headers, ",")
				logp.Debug("csvheader", "header: %s", header)
				// Set header as first csv row.
				rows[0] = header
				continue
			}

			if len(rec[i]) != len(headers) {
				if len(rec[i]) > 1 {
					logp.Err("row length(%d) != header length(%d)", len(rec[i]), len(headers))
				}
				continue
			}

			records := c.formRecords(headers, rec[i])
			record := strings.Join(records, ",")
			rows = append(rows, record)

			//printSchema(headers, records)

			linesRead++
			if len(rows) >= c.batchSize {
				// Reached max batch size, send to channel and reset
				logp.Debug("scan", "reached max batch size %d\n", len(rows))
				logp.Debug("scan", "sending following rows to channel %s\n", rows)
				select {
				case c.csvChan <- rows:
				default:
					logp.Warn("overflowing CSV queue channel")
				}
				rows = make([]string, 1, c.batchSize)
				// Use header as first csv row.
				rows[0] = header
			}
		}

		// Finished reading input, make sure last batch goes out.
		if len(rows) > 1 {
			logp.Debug("scan", "last batch row size %d\n", len(rows))
			logp.Debug("scan", "sending following rows to channel %#s\n", rows)
			select {
			case c.csvChan <- rows:
			default:
				logp.Warn("overflowing CSV queue channel")
			}
		}

		scanTook := time.Now().Sub(scanStart)

		if linesRead == 1 {
			logp.Info("scan completed with header only so we do nothing\n")
		} else {
			logp.Info("scan took %v for %d cdr rowslines\n", scanTook, linesRead)
		}
	}

	go watch.Start(scan, c.shut)
}

// cutspace will cut all spaces from string.
func cutSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func printSchema(header, rows []string) {
	fmt.Println("##########################")
	for k := range header {
		fmt.Println(header[k] + "=" + rows[k])
	}
	fmt.Println("##########################")
	for k := range header {
		fmt.Printf("CREATE INDEX IF NOT EXISTS %s_%s ON %s (\"%s\");\n", config.Setting.CDRDBTable, header[k], config.Setting.CDRDBTable, header[k])
	}
	fmt.Println("##########################")
	for k := range header {
		fmt.Printf("%s TEXT DEFAULT '',\n", header[k])
	}
}
