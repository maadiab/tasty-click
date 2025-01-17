-- name: GetReciept :many
SELECT order_foods.order_id, order_foods.quantity,foods.name
 from order_foods
 JOIN foods
 ON order_foods.food_id = foods.id
WHERE order_foods.order_id =$1;

