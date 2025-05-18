-- name: CreateCategory :exec
INSERT INTO categories (name) VALUES ($1);

-- name: GetCategories :many
SELECT * FROM categories;

-- name: GetCategoryByName :one
SELECT * FROM categories WHERE name = $1;

-- name: UpdateCategory :exec
UPDATE categories SET name = $1 WHERE name = $2;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE name = $1;
