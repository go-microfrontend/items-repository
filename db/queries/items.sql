-- name: CreateItem :one
Insert Into items (
  name,
  description,
  type,
  weight_in_grams
)
Values ($1, $2, $3, $4)
Returning id;

-- name: GetItemByID :one
Select * From items
Where id = $1 Limit 1;

-- name: GetItems :many
Select * From items
Order By Name
Limit $1 Offset $2;
