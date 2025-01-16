-- name: GetUsers :many
SELECT * FROM users;

-- name: GetOneUser :one
SELECT FROM users WHERE id =$1;

-- name: CreateUser :exec
INSERT INTO users (name,username,mobile,password,type)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users SET name =$1,username=$2,mobile=$3,type=$4 WHERE id =$5;

-- name: DeleteUser :exec
DELETE FROM users WHERE id =$1;
