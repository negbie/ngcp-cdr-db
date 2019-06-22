package database

import (
	"database/sql"
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	_ "github.com/lib/pq"
	"github.com/negbie/logp"
	"github.com/negbie/ngcp-cdr-db/config"
	"github.com/robfig/cron"
)

type DBObject struct {
	db          *sql.DB
	copyOpts    string
	columnCount int64
	rowCount    int64
	splitChar   string
	delimStr    string
	tableName   string
}

func (dbo *DBObject) setup() error {
	if err := dbo.dbOpen(); err != nil {
		logp.Err("%v", err)
		return err
	}
	// Postgres copy command parameters.
	dbo.copyOpts = config.Setting.CSVCopyOpts
	// Define how the CSV columns are seperated.
	dbo.splitChar = config.Setting.CSVSplitChar
	// Quote delimiter string.
	dbo.delimStr = fmt.Sprintf("'%s'", dbo.splitChar)
	// Get Postgres table shema and name.
	dbo.tableName = getFullTableName()

	// Need to cover Postgres copy tab delimiter.
	if dbo.splitChar == "\\t" {
		dbo.delimStr = "E" + dbo.delimStr
	}
	// Need to cover Go's strings split function.
	if dbo.splitChar == "\\t" {
		dbo.splitChar = "\t"
	}

	// Start stats reporting function.
	go dbo.report()
	go dbo.rotate(config.Setting.CDRDBRotate, config.Setting.CDRDBTable)

	return nil
}

func (dbo *DBObject) dbOpen() error {
	// Connect to Postgres.
	conn, err := sql.Open("postgres", getConnectString())
	if err != nil {
		return err
	}
	logp.Info("try to open database connections...")

	if err := conn.Ping(); err != nil {
		conn.Close()
		return err
	}
	logp.Info("cdr database connection successfully established")
	conn.SetMaxOpenConns(20)
	conn.SetMaxIdleConns(10)
	dbo.db = conn
	return nil
}

func getConnectString() string {
	return fmt.Sprintf("sslmode=disable connect_timeout=2 host=%s port=%d dbname=%s user=%s password=%s",
		config.Setting.CDRDBHost,
		config.Setting.CDRDBPort,
		config.Setting.CDRDBName,
		config.Setting.CDRDBUser,
		config.Setting.CDRDBPass)
}

func getFullTableName() string {
	return fmt.Sprintf("\"%s\".\"%s\"", config.Setting.CDRDBSchema, config.Setting.CDRDBTable)
}

func (dbo *DBObject) insert(rows chan []string) {
	// Recover if Postgres panics.
	defer func() {
		if r := recover(); r != nil {
			logp.Err("recover %v", r)
		}
	}()

	cc := int64(0)
	for r := range rows {
		start := time.Now()

		// Start SQL transaction.
		tx, err := dbo.db.Begin()
		if err != nil || tx == nil {
			logp.Err("%v", err)
			select {
			case rows <- r:
			default:
				logp.Warn("overflowing CSV queue channel")
			}
			continue
		}

		// Prepare Postgres copy command.
		copyCmd := fmt.Sprintf("COPY %s(%s) FROM STDIN WITH DELIMITER %s %s", dbo.tableName, r[0], dbo.delimStr, dbo.copyOpts)
		stmt, err := tx.Prepare(copyCmd)
		if err != nil {
			logp.Err("%v", err)
			err := tx.Rollback()
			if err != nil {
				logp.Err("%v", err)
			}
			continue
		}

		// Read CSV line by line and add the row
		// to the SQL transaction.
		for _, line := range r {
			// For some reason this is only needed for tab splitting
			if dbo.splitChar == "\t" {
				sp := strings.Split(line, dbo.splitChar)
				cc += int64(len(sp))
				args := make([]interface{}, len(sp))
				for i, v := range sp {
					args[i] = v
				}
				logp.Debug("db", "insert row %s", line)
				_, err = stmt.Exec(args...)
			} else {
				logp.Debug("db", "insert row %s", line)
				_, err = stmt.Exec(line)
			}
			if err != nil {
				logp.Err("%v", err)
			}
		}
		// Raise stats counter.
		atomic.AddInt64(&dbo.columnCount, cc)
		atomic.AddInt64(&dbo.rowCount, int64(len(r)))
		cc = 0

		err = stmt.Close()
		if err != nil {
			logp.Err("%v", err)
		}

		// Commit SQL transaction.
		err = tx.Commit()
		if err != nil {
			logp.Err("%v", err)
		}

		end := time.Now().Sub(start)
		logp.Info("insert took %v, row size %d, row rate %f/sec\n", end, len(r), float64(len(r))/float64(end.Seconds()))
	}
}

// report periodically prints the write rate in number of rows per second
func (dbo *DBObject) report() {
	start := time.Now()
	prevTime := start
	prevRowCount := int64(0)

	for now := range time.NewTicker(60 * time.Second).C {
		rCount := atomic.LoadInt64(&dbo.rowCount)

		took := now.Sub(prevTime)
		rowrate := float64(rCount-prevRowCount) / float64(took.Seconds())
		overallRowrate := float64(rCount) / float64(now.Sub(start).Seconds())
		totalTook := now.Sub(start)

		logp.Info("at %v, row rate %f/sec (period), row rate %f/sec (since start), rows total %E (since start)\n",
			totalTook-(totalTook%time.Second), rowrate, overallRowrate, float64(rCount))

		prevRowCount = rCount
		prevTime = now
	}
}

func (dbo *DBObject) rotate(r, t string) {
	if r != "" && t != "" {
		dc := cron.New()
		dc.AddFunc("0 30 03 * * *", func() {
			if err := dbo.dropChunks(r, t); err != nil {
				logp.Err("%v", err)
			}
			logp.Info("next rotate will be at %v\n", time.Now().Add(time.Hour*24+1))
		})
		dc.Start()
	}
}

func (dbo *DBObject) dropChunks(r, t string) error {
	q := fmt.Sprintf("SELECT drop_chunks(interval '%s', '%s');", r, t)
	_, err := dbo.db.Exec(q)
	return err
}
