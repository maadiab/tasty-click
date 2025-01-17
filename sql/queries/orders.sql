-- name: GetOrders :many
SELECT * FROM orders;

-- name: GetOneOrder :one
SELECT * FROM orders WHERE id=$1;

-- name: CreateOrder :exec
INSERT INTO orders (customer_id,created_at,updated_at,status) VALUES ($1,NOW(),NOW(),$2)
RETURNING *;

-- name: UpdateOrder :exec
UPDATE orders SET updated_at = NOW(), status=$1 WHERE id =$2;

-- name: DeleteOrder :exec
DELETE FROM orders WHERE id =$1;
