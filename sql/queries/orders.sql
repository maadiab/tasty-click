-- name: GetOrders :many
SELECT * FROM orders;

-- name: GetOneOrder :one
SELECT FROM orders WHERE id=$1;

-- name: CreateOrder :exec
INSERT INTO orders (created_at,updated_at) VALUES (NOW(),NOW())
RETURNING *;

-- name: UpdateOrder :exec
UPDATE orders SET updated_at = NOW() WHERE id =$1;

-- name: DeleteOrder :exec
DELETE FROM orders WHERE id =$1;
