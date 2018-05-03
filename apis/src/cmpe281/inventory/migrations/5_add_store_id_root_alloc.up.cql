DROP TABLE IF EXISTS starbucks.allocations;

CREATE TABLE IF NOT EXISTS starbucks.allocations (
  store_id UUID,
  user_id UUID,
  id UUID,
  status text,
  expires timestamp,
  products frozen<set<product_details>>,
  PRIMARY KEY((store_id, user_id), id)
);