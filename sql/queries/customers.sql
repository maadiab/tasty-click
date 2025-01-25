-- name: AddCustomer :exec
INSERT INTO customers (name,mobile,email,password) VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: GetCustomer :one
SELECT * FROM customers WHERE id =$1;

-- name: GetAllCustomers :many
SELECT * FROM customers;

-- name: UpdateCustomer :exec
UPDATE customers SET name=$1,mobile=$2,email=$3 WHERE id=$4;

-- name: DeleteCustomr :exec
DELETE FROM customers WHERE id=$1;

