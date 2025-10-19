-- name: CreateOrder :one
INSERT INTO orders (id, user_id, driver_id, status, total_amount, delivery_address, created_at, updated_at)
VALUES (gen_random_uuid(), $1, NULL, 'pending', $2, $3, NOW(), NOW())
RETURNING *;

-- name: GetOrderByID :one
SELECT * FROM orders WHERE id = $1;

-- name: GetOrdersByUser :many
SELECT * FROM orders WHERE user_id = $1 ORDER BY created_at DESC;

-- name: AssignDriverToOrder :exec
UPDATE orders SET driver_id = $2, status = 'delivering', updated_at = NOW() WHERE id = $1;

-- name: UpdateOrderStatus :exec
UPDATE orders SET status = $2, updated_at = NOW() WHERE id = $1;

-- name: DeleteOrder :exec
DELETE FROM orders WHERE id = $1;
