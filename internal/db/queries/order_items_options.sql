-- name: AddOrderItemOption :exec
INSERT INTO order_item_options (id, order_item_id, option_id, price_difference)
VALUES (gen_random_uuid(), $1, $2, $3);

-- name: GetOrderItemOptions :many
SELECT * FROM order_item_options WHERE order_item_id = $1;
