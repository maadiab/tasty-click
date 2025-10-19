-- name: AddOrderItem :one
INSERT INTO order_items (id, order_id, menu_item_id, quantity, price)
VALUES (gen_random_uuid(), $1, $2, $3, $4)
RETURNING *;

-- name: GetOrderItems :many
SELECT * FROM order_items WHERE order_id = $1;

-- name: DeleteOrderItem :exec
DELETE FROM order_items WHERE id = $1;
