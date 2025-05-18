-- name: CreateProductCharacteristic :exec
INSERT INTO product_characteristics (
    product_id,
    description,
    weight,
    quantity_in_package,
    shelf_life,
    storage_conditions,
    nutrition
) VALUES (
    @product_id,
    @description,
    @weight,
    @quantity_in_package,
    @shelf_life,
    @storage_conditions,
    ROW(@proteins::int, @fats::int, @carbohydrates::int, @calories::int)::nutritional_info
);

-- name: GetProductCharacteristics :many
SELECT * FROM product_characteristics;

-- name: GetProductCharacteristicByID :one
SELECT * FROM product_characteristics WHERE product_id = $1;

-- name: UpdateProductCharacteristic :exec
UPDATE product_characteristics
SET description = $1,
    weight = $2,
    quantity_in_package = $3,
    shelf_life = $4,
    storage_conditions = $5,
    nutrition = ROW($6, $7, $8, $9)::nutritional_info
WHERE product_id = $10;

-- name: DeleteProductCharacteristicByID :exec
DELETE FROM product_characteristics WHERE product_id = $1;
