package config

import (
	"time"
)

var Setting Configuration

type Configuration struct {
	CDRDBHost      string        `default:"localhost"`
	CDRDBPort      int           `default:"5432"`
	CDRDBName      string        `default:"vmess"`
	CDRDBUser      string        `default:"ngcp-cdr"`
	CDRDBPass      string        `default:"ngcp-cdr"`
	CDRDBTable     string        `default:"ngcp_cdr"`
	CDRDBSchema    string        `default:"public"`
	CSVTimeColumn  string        `default:"start_time"`
	CSVTimeFormat  string        `default:"2006-01-02 15:04:05.999"`
	CSVCopyOpts    string        `default:"CSV HEADER"`
	CSVSplitChar   string        `default:","`
	CSVHeader      string        `default:"cdr_id, update_time, source_user_id, source_provider_id, source_external_subscriber_id, source_subscriber_id, source_external_contract_id, source_account_id, source_user, source_domain, source_cli, source_clir, source_ip, destination_user_id, destination_provider_id, destination_external_subscriber_id, destination_subscriber_id, destination_external_contract_id, destination_account_id, destination_user, destination_domain, destination_user_in, destination_domain_in, destination_user_dialed, peer_auth_user, peer_auth_realm, call_type, call_status, call_code, init_time, start_time, duration, call_id, rating_status, rated_at, source_carrier_cost, source_customer_cost, source_carrier_zone, source_customer_zone, source_carrier_detail, source_customer_detail, source_carrier_free_time, source_customer_free_time, destination_carrier_cost, destination_customer_cost, destination_carrier_zone, destination_customer_zone, destination_carrier_detail, destination_customer_detail, destination_carrier_free_time, destination_customer_free_time, source_reseller_cost, source_reseller_zone, source_reseller_detail, source_reseller_free_time, destination_reseller_cost, destination_reseller_zone, destination_reseller_detail, destination_reseller_free_time"`
	CSVBatchSize   int           `default:"200"`
	CSVQueueSize   int           `default:"10000"`
	WatchFolder    string        `default:"./example/rated"`
	WatchRecursive bool          `default:"false"`
	WatchMaxEvent  int           `default:"20"`
	WatchTime      time.Duration `default:"2000ms"`
	LogDbg         string        `default:""`
	LogLvl         string        `default:"info"`
	LogStd         bool          `default:"false"`
	LogSys         bool          `default:"false"`
	ConfigFile     string        `default:"./ngcp-cdr-db.toml"`
	DryRun         bool          `default:"true"`
	Version        bool          `default:"false"`
}
