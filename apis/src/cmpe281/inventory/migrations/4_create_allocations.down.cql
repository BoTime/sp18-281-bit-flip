DROP TABLE IF EXISTS starbucks.allocations;

CREATE TABLE IF NOT EXISTS starbucks.allocations (
  id UUID PRIMARY KEY,
  status text,
  expires timestamp
);

DROP TYPE IF EXISTS starbucks.product_details;