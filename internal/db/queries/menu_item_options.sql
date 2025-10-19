-- name: CreateMenuItemOption :one
INSERT INTO menu_item_options (id, menu_item_id, name, price)
VALUES (gen_random_uuid(), $1, $2, $3)
RETURNING *;

-- name: GetOptionsForMenuItem :many
SELECT * FROM menu_item_options WHERE menu_item_id = $1 ORDER BY name;

-- name: DeleteMenuItemOption :exec
DELETE FROM menu_item_options WHERE id = $1;
