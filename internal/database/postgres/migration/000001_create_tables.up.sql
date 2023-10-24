BEGIN;

CREATE TABLE IF NOT EXISTS product(
    id serial PRIMARY KEY,
    product_name VARCHAR(50) UNIQUE NOT NULL,
    price FLOAT,
    quantity INT
);

COMMIT;