CREATE KEYSPACE IF NOT EXISTS starbucks WITH REPLICATION = {
  'class': 'NetworkTopologyStrategy',
  'dc1': '3',
  'dc2': '3'
};

CREATE TABLE IF NOT EXISTS starbucks.stores (
  id UUID PRIMARY KEY,
  name text
);

CREATE TABLE IF NOT EXISTS starbucks.products (
  store_id UUID,
  id UUID,
  name text,
  quantity text,
  size text,
  PRIMARY KEY(store_id, id)
);

CREATE TABLE IF NOT EXISTS starbucks.allocations (
  id UUID PRIMARY KEY,
  status text,
  expires timestamp
);


