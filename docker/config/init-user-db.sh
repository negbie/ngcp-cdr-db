#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL

    CREATE DATABASE vmess;
    CREATE USER admin WITH PASSWORD 'admin';
    CREATE USER grafana WITH PASSWORD 'grafana';
    GRANT ALL PRIVILEGES ON DATABASE vmess TO admin;

    \c vmess;

    CREATE TABLE ngcp_cdr (
        cdr_id TEXT DEFAULT '',
        update_time TIMESTAMP NOT NULL,
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
        init_time TIMESTAMP NOT NULL,
        "time" TIMESTAMP NOT NULL,
        duration INTEGER DEFAULT 0,
        call_id TEXT DEFAULT '',
        rating_status TEXT DEFAULT '',
        rated_at TIMESTAMP NOT NULL,
        source_carrier_cost TEXT DEFAULT '',
        source_customer_cost TEXT DEFAULT '',
        source_carrier_zone TEXT DEFAULT '',
        source_customer_zone TEXT DEFAULT '',
        source_carrier_detail TEXT DEFAULT '',
        source_customer_detail TEXT DEFAULT '',
        setup_duration INTEGER DEFAULT 0,
        end_time TIMESTAMP NOT NULL,
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
        dialog_duration INTEGER DEFAULT 0,
        destination_reseller_cost TEXT DEFAULT '',
        destination_reseller_zone TEXT DEFAULT '',
        destination_reseller_detail TEXT DEFAULT '',
        "owner" TEXT DEFAULT ''
    );


    CREATE INDEX IF NOT EXISTS ngcp_cdr_source_user_id ON ngcp_cdr ("source_user_id");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_source_subscriber_id ON ngcp_cdr ("source_subscriber_id");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_source_external_contract_id ON ngcp_cdr ("source_external_contract_id");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_source_account_id ON ngcp_cdr ("source_account_id");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_source_user ON ngcp_cdr ("source_user");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_source_domain ON ngcp_cdr ("source_domain");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_source_cli ON ngcp_cdr ("source_cli");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_source_ip ON ngcp_cdr ("source_ip");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_user_id ON ngcp_cdr ("destination_user_id");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_subscriber_id ON ngcp_cdr ("destination_subscriber_id");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_external_contract_id ON ngcp_cdr ("destination_external_contract_id");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_account_id ON ngcp_cdr ("destination_account_id");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_user ON ngcp_cdr ("destination_user");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_domain ON ngcp_cdr ("destination_domain");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_user_in ON ngcp_cdr ("destination_user_in");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_domain_in ON ngcp_cdr ("destination_domain_in");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_destination_user_dialed ON ngcp_cdr ("destination_user_dialed");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_call_code ON ngcp_cdr ("call_code");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_init_time ON ngcp_cdr ("init_time");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_time ON ngcp_cdr ("time");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_duration ON ngcp_cdr ("duration");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_call_id ON ngcp_cdr ("call_id");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_setup_duration ON ngcp_cdr ("setup_duration");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_end_time ON ngcp_cdr ("end_time");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_direction ON ngcp_cdr ("direction");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_prefix ON ngcp_cdr ("prefix");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_dialog_duration ON ngcp_cdr ("dialog_duration");
    CREATE INDEX IF NOT EXISTS ngcp_cdr_owner ON ngcp_cdr ("owner");


    CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;
    SELECT create_hypertable('ngcp_cdr', 'time', chunk_time_interval => interval '1 day');
    GRANT ALL PRIVILEGES ON TABLE ngcp_cdr TO admin;
    GRANT ALL PRIVILEGES ON TABLE ngcp_cdr TO grafana;

EOSQL