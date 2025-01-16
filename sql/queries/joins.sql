-- name: GetReciept :many
SELECT foods_orders.order_id, foods_orders.quantity,foods.name
 from foods_orders
 JOIN foods
 ON foods_orders.food_id = foods.id
WHERE foods_orders.order_id =$1;

