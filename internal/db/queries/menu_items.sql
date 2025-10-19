-- name: CreateMenuItem :one
INSERT INTO menu_items (id, category_id, name, description, price, image_url, available)
VALUES (gen_random_uuid(), $1, $2, $3, $4, $5, TRUE)
RETURNING *;

-- name: GetMenuItemsByCategory :many
SELECT * FROM menu_items WHERE category_id = $1 ORDER BY name;

-- name: GetMenuItemByID :one
SELECT * FROM menu_items WHERE id = $1;

-- name: UpdateMenuItemAvailability :exec
UPDATE menu_items SET available = $2 WHERE id = $1;

-- name: DeleteMenuItem :exec
DELETE FROM menu_items WHERE id = $1;
