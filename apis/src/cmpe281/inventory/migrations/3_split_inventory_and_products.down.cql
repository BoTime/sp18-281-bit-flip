DROP TABLE IF EXISTS starbucks.inventory;

DROP INDEX IF EXISTS starbucks.product_name_idx;

DROP TABLE IF EXISTS starbucks.products;

CREATE TABLE IF NOT EXISTS starbucks.products (
  store_id UUID,
  id UUID,
  name text,
  quantity text,
  size text,
  PRIMARY KEY(store_id, id)
);

CREATE INDEX IF NOT EXISTS product_name_idx ON starbucks.products (name);