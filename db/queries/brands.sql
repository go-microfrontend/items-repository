-- name: CreateBrand :exec
INSERT INTO brands (name) VALUES ($1);

-- name: GetBrands :many
SELECT * FROM brands;

-- name: GetBrandByName :one
SELECT * FROM brands WHERE name = $1;

-- name: UpdateBrand :exec
UPDATE brands SET name = $1 WHERE name = $2;

-- name: DeleteBrand :exec
DELETE FROM brands WHERE name = $1;
