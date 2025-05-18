-- name: CreateProduct :one
INSERT INTO products (name, brand, category, price, rating)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetProducts :many
SELECT * FROM products;

-- name: GetProductByID :one
SELECT * FROM products WHERE product_id = $1;

-- name: GetProductsByCategory :many
SELECT * FROM products WHERE category = $1;

-- name: UpdateProduct :exec
UPDATE products
SET name = $1,
    brand = $2,
    category = $3,
    price = $4,
    rating = $5
WHERE product_id = $6;

-- name: DeleteProduct :exec
DELETE FROM products WHERE product_id = $1;
