-- name: CreateItem :one
Insert Into items (
  name,
  description,
  type,
  weight_in_grams,
  amount
)
Values ($1, $2, $3, $4, $5)
Returning id;

-- name: GetItemByID :one
Select * From items
Where id = $1 Limit 1;

-- name: GetItems :many
Select * From items
Order By Name
Limit $1 Offset $2;

-- name: GetItemsByType :many
Select * From items
Where Type = $1
Order By Name
Limit $2 Offset $3;
