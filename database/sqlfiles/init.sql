CREATE DATABASE vmess;
CREATE USER ngcp-cdr WITH PASSWORD 'ngcp-cdr';
GRANT ALL PRIVILEGES ON DATABASE vmess TO ngcp-cdr;

\c vmess;

CREATE TABLE ngcp-cdr (
    cdr_id TEXT DEFAULT '',
    update_time TEXT DEFAULT '',
    source_user_id TEXT DEFAULT '',
    source_provider_id TEXT DEFAULT '',
    source_external_subscriber_id TEXT DEFAULT '',
    source_subscriber_id TEXT DEFAULT '',
    source_external_contract_id TEXT DEFAULT '',
    source_account_id TEXT DEFAULT '',
    source_user TEXT DEFAULT '',
    source_domain TEXT DEFAULT '',
    source_cli TEXT DEFAULT '',
    source_clir TEXT DEFAULT '',
    source_ip TEXT DEFAULT '',
    destination_user_id TEXT DEFAULT '',
    destination_provider_id TEXT DEFAULT '',
    destination_external_subscriber_id TEXT DEFAULT '',
    destination_subscriber_id TEXT DEFAULT '',
    destination_external_contract_id TEXT DEFAULT '',
    destination_account_id TEXT DEFAULT '',
    destination_user TEXT DEFAULT '',
    destination_domain TEXT DEFAULT '',
    destination_user_in TEXT DEFAULT '',
    destination_domain_in TEXT DEFAULT '',
    destination_user_dialed TEXT DEFAULT '',
    peer_auth_user TEXT DEFAULT '',
    peer_auth_realm TEXT DEFAULT '',
    call_type TEXT DEFAULT '',
    call_status TEXT DEFAULT '',
    call_code TEXT DEFAULT '',
    init_time TEXT DEFAULT '',
    time TEXT DEFAULT '',
    duration TEXT DEFAULT '',
    call_id TEXT DEFAULT '',
    rating_status TEXT DEFAULT '',
    rated_at TEXT DEFAULT '',
    source_carrier_cost TEXT DEFAULT '',
    source_customer_cost TEXT DEFAULT '',
    source_carrier_zone TEXT DEFAULT '',
    source_customer_zone TEXT DEFAULT '',
    source_carrier_detail TEXT DEFAULT '',
    source_customer_detail TEXT DEFAULT '',
    setup_time TEXT DEFAULT '',
    end_time TEXT DEFAULT '',
    destination_carrier_cost TEXT DEFAULT '',
    destination_customer_cost TEXT DEFAULT '',
    destination_carrier_zone TEXT DEFAULT '',
    destination_customer_zone TEXT DEFAULT '',
    destination_carrier_detail TEXT DEFAULT '',
    destination_customer_detail TEXT DEFAULT '',
    direction TEXT DEFAULT '',
    prefix TEXT DEFAULT '',
    source_reseller_cost TEXT DEFAULT '',
    source_reseller_zone TEXT DEFAULT '',
    source_reseller_detail TEXT DEFAULT '',
    dialog_time TEXT DEFAULT '',
    destination_reseller_cost TEXT DEFAULT '',
    destination_reseller_zone TEXT DEFAULT '',
    destination_reseller_detail TEXT DEFAULT '',
    trunk TEXT DEFAULT '',
);


CREATE INDEX IF NOT EXISTS ngcp_cdr_cdr_id ON ngcp_cdr ("cdr_id");
CREATE INDEX IF NOT EXISTS ngcp_cdr_update_time ON ngcp_cdr ("update_time");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_user_id ON ngcp_cdr ("source_user_id");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_provider_id ON ngcp_cdr ("source_provider_id");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_external_subscriber_id ON ngcp_cdr ("source_external_subscriber_id");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_subscriber_id ON ngcp_cdr ("source_subscriber_id");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_external_contract_id ON ngcp_cdr ("source_external_contract_id");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_account_id ON ngcp_cdr ("source_account_id");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_user ON ngcp_cdr ("source_user");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_domain ON ngcp_cdr ("source_domain");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_cli ON ngcp_cdr ("source_cli");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_clir ON ngcp_cdr ("source_clir");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_ip ON ngcp_cdr ("source_ip");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_user_id ON ngcp_cdr ("destination_user_id");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_provider_id ON ngcp_cdr ("destination_provider_id");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_external_subscriber_id ON ngcp_cdr ("destination_external_subscriber_id");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_subscriber_id ON ngcp_cdr ("destination_subscriber_id");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_external_contract_id ON ngcp_cdr ("destination_external_contract_id");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_account_id ON ngcp_cdr ("destination_account_id");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_user ON ngcp_cdr ("destination_user");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_domain ON ngcp_cdr ("destination_domain");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_user_in ON ngcp_cdr ("destination_user_in");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_domain_in ON ngcp_cdr ("destination_domain_in");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_user_dialed ON ngcp_cdr ("destination_user_dialed");
CREATE INDEX IF NOT EXISTS ngcp_cdr_peer_auth_user ON ngcp_cdr ("peer_auth_user");
CREATE INDEX IF NOT EXISTS ngcp_cdr_peer_auth_realm ON ngcp_cdr ("peer_auth_realm");
CREATE INDEX IF NOT EXISTS ngcp_cdr_call_type ON ngcp_cdr ("call_type");
CREATE INDEX IF NOT EXISTS ngcp_cdr_call_status ON ngcp_cdr ("call_status");
CREATE INDEX IF NOT EXISTS ngcp_cdr_call_code ON ngcp_cdr ("call_code");
CREATE INDEX IF NOT EXISTS ngcp_cdr_init_time ON ngcp_cdr ("init_time");
CREATE INDEX IF NOT EXISTS ngcp_cdr_time ON ngcp_cdr ("time");
CREATE INDEX IF NOT EXISTS ngcp_cdr_duration ON ngcp_cdr ("duration");
CREATE INDEX IF NOT EXISTS ngcp_cdr_call_id ON ngcp_cdr ("call_id");
CREATE INDEX IF NOT EXISTS ngcp_cdr_rating_status ON ngcp_cdr ("rating_status");
CREATE INDEX IF NOT EXISTS ngcp_cdr_rated_at ON ngcp_cdr ("rated_at");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_carrier_cost ON ngcp_cdr ("source_carrier_cost");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_customer_cost ON ngcp_cdr ("source_customer_cost");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_carrier_zone ON ngcp_cdr ("source_carrier_zone");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_customer_zone ON ngcp_cdr ("source_customer_zone");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_carrier_detail ON ngcp_cdr ("source_carrier_detail");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_customer_detail ON ngcp_cdr ("source_customer_detail");
CREATE INDEX IF NOT EXISTS ngcp_cdr_setup_time ON ngcp_cdr ("setup_time");
CREATE INDEX IF NOT EXISTS ngcp_cdr_end_time ON ngcp_cdr ("end_time");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_carrier_cost ON ngcp_cdr ("destination_carrier_cost");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_customer_cost ON ngcp_cdr ("destination_customer_cost");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_carrier_zone ON ngcp_cdr ("destination_carrier_zone");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_customer_zone ON ngcp_cdr ("destination_customer_zone");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_carrier_detail ON ngcp_cdr ("destination_carrier_detail");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_customer_detail ON ngcp_cdr ("destination_customer_detail");
CREATE INDEX IF NOT EXISTS ngcp_cdr_direction ON ngcp_cdr ("direction");
CREATE INDEX IF NOT EXISTS ngcp_cdr_prefix ON ngcp_cdr ("prefix");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_reseller_cost ON ngcp_cdr ("source_reseller_cost");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_reseller_zone ON ngcp_cdr ("source_reseller_zone");
CREATE INDEX IF NOT EXISTS ngcp_cdr_source_reseller_detail ON ngcp_cdr ("source_reseller_detail");
CREATE INDEX IF NOT EXISTS ngcp_cdr_dialog_time ON ngcp_cdr ("dialog_time");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_reseller_cost ON ngcp_cdr ("destination_reseller_cost");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_reseller_zone ON ngcp_cdr ("destination_reseller_zone");
CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_reseller_detail ON ngcp_cdr ("destination_reseller_detail");
CREATE INDEX IF NOT EXISTS ngcp_cdr_trunk ON ngcp_cdr ("trunk");


CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;
SELECT create_hypertable('ngcp_cdr', 'time', 'trunk');
GRANT ALL PRIVILEGES ON TABLE ngcp_cdr TO ngcp-cdr;
GRANT ALL PRIVILEGES ON TABLE ngcp_cdr TO grafana;
