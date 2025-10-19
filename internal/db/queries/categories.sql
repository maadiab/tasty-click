-- name: CreateMenuCategory :one
INSERT INTO categories (id, name)
VALUES (gen_random_uuid(), $1)
RETURNING *;

-- name: GetAllCategories :many
SELECT * FROM categories ORDER BY name;

-- name: DeleteMenuCategory :exec
DELETE FROM categories WHERE id = $1;
