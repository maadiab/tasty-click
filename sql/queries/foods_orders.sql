-- name: CreateFoodsOrder :exec
INSERT INTO foods_orders (order_id,customer_id,food_id,quantity)
VALUES ($1,$2,$3,$4)
RETURNING *;



