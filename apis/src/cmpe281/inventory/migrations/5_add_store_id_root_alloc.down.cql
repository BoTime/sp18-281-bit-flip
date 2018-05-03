DROP TABLE IF EXISTS starbucks.allocations;

CREATE TABLE IF NOT EXISTS starbucks.allocations (
  user_id UUID,
  id UUID,
  status text,
  expires timestamp,
  products frozen<set<product_details>>,
  PRIMARY KEY(user_id, id)
);