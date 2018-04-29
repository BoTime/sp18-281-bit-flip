DROP INDEX IF EXISTS starbucks.products_name_idx;

DROP TABLE IF EXISTS starbucks.products;

CREATE TABLE IF NOT EXISTS starbucks.products (
  id UUID,
  name text,
  size text,
  PRIMARY KEY(id)
);

CREATE INDEX IF NOT EXISTS product_name_idx ON starbucks.products (name);

CREATE TABLE IF NOT EXISTS starbucks.inventory (
  store_id UUID,
  id UUID,
  name text,
  quantity text,
  size text,
  PRIMARY KEY(store_id, id)
);