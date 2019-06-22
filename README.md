# ngcp-cdr-db
Is a simple program which scans a configurable folder for new sipwise cdr's and sends them as a batch to timescaledb.

## Init
```bash
cd docker  
sudo docker-compose up -d  
```
Download ngcp-cdr-db from releases or compile it. Run it with the flags you need.

## Usage
go to localhost:8088  
login with superset/superset  

## Flags
```bash
Use with ./ngcp-cdr-db
  -cdrdbhost
        Change value of CDRDBHost. (default localhost)
  -cdrdbname
        Change value of CDRDBName. (default vmess)
  -cdrdbpass
        Change value of CDRDBPass. (default root)
  -cdrdbport
        Change value of CDRDBPort. (default 5432)
  -cdrdbrotate
        Change value of CDRDBRotate. (default 3 months)
  -cdrdbschema
        Change value of CDRDBSchema. (default public)
  -cdrdbtable
        Change value of CDRDBTable. (default ngcp_cdr)
  -cdrdbuser
        Change value of CDRDBUser. (default root)
  -configfile
        Change value of ConfigFile. (default ./ngcp-cdr-db.toml)
  -csvbatchsize
        Change value of CSVBatchSize. (default 500)
  -csvcopyopts
        Change value of CSVCopyOpts. (default CSV HEADER)
  -csvqueuesize
        Change value of CSVQueueSize. (default 20000)
  -csvsplitchar
        Change value of CSVSplitChar. (default ,)
  -csvtimecolumn
        Change value of CSVTimeColumn. (default start_time)
  -csvtimeformat
        Change value of CSVTimeFormat. (default 2006-01-02 15:04:05.999)
  -dryrun
        Change value of DryRun. (default false)
  -logdbg
        Change value of LogDbg.
  -loglvl
        Change value of LogLvl. (default info)
  -logstd
        Change value of LogStd. (default false)
  -logsys
        Change value of LogSys. (default false)
  -version
        Change value of Version. (default false)
  -watchfolder
        Change value of WatchFolder. (default ./example/rated)
  -watchmaxevent
        Change value of WatchMaxEvent. (default 30)
  -watchrecursive
        Change value of WatchRecursive. (default false)
  -watchtime
        Change value of WatchTime. (default 5s)

```