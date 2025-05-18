-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE nutritional_info AS (
    proteins      INT,
    fats          INT,
    carbohydrates INT,
    calories      INT
);

CREATE TABLE categories (
    name VARCHAR(255) PRIMARY KEY,
    label VARCHAR(255)
);

CREATE TABLE brands (
    name VARCHAR(255) PRIMARY KEY
);

CREATE TABLE products (
    product_id  UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    brand       VARCHAR(255) REFERENCES brands(name) ON DELETE SET NULL,
    category    VARCHAR(255) REFERENCES categories(name) ON DELETE SET NULL,
    price       BIGINT NOT NULL,
    rating      REAL
  );

CREATE TABLE product_characteristics (
    product_id          UUID PRIMARY KEY REFERENCES products(product_id) ON DELETE CASCADE,
    description         TEXT,
    weight              INT,
    quantity_in_package INT,
    shelf_life          INTERVAL,
    storage_conditions  VARCHAR(255),
    nutrition           nutritional_info
);

-- +goose Down
DROP TABLE product_characteristics;
DROP TABLE products;
DROP TABLE brands;
DROP TABLE categories;
DROP TYPE nutritional_info;
DROP EXTENSION "uuid-ossp";
