# Create Keyspace
CREATE KEYSPACE starbucks WITH replication = {
    'class': 'NetworkTopologyStrategy',
    'dc1': '4'
  } AND durable_writes = true;

# Create Type for Billing Details
CREATE TYPE starbucks.billing_details (
    first_name text,
    last_name text,
    line1 text,
    line2 text,
    city text,
    state text,
    zipcode text
);

# Create Type for Card Details
CREATE TYPE starbucks.card_details (
    number text,
    expiration date
);

# Create Payments Table
CREATE TABLE starbucks.payments (
    user_id uuid,
    payment_id timeuuid,
    billing_details frozen<billing_details>,
    card_details frozen<card_details>,
    amount double,
    status text,
    PRIMARY KEY (user_id, payment_id)
) WITH CLUSTERING ORDER BY (payment_id ASC)
    AND bloom_filter_fp_chance = 0.01
    AND caching = {'keys': 'ALL', 'rows_per_partition': 'NONE'}
    AND comment = 'Starbucks Online Ordering Payment Records'
    AND compaction = {'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy', 'max_threshold': '32', 'min_threshold': '4'}
    AND compression = {'chunk_length_in_kb': '64', 'class': 'org.apache.cassandra.io.compress.LZ4Compressor'}
    AND crc_check_chance = 1.0
    AND dclocal_read_repair_chance = 0.1
    AND default_time_to_live = 0
    AND gc_grace_seconds = 864000
    AND max_index_interval = 2048
    AND memtable_flush_period_in_ms = 0
    AND min_index_interval = 128
    AND read_repair_chance = 0.0
    AND speculative_retry = '99PERCENTILE';

# Create Index for Payment ID lookups
CREATE INDEX payments_id_idx ON starbucks.payments (payment_id);
