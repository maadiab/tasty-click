-- name: GetFoods :many
SELECT * FROM foods;

-- name: GetOneFood :one
SELECT * FROM foods WHERE id =$1;

-- name: CreateFood :exec
INSERT INTO foods (name,price,picture,category)
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: UpdateFood :exec
UPDATE foods SET name=$1,price=$2,picture=$3,category=$4 WHERE id =$5;

-- name: DeleteFood :exec
DELETE FROM foods WHERE id =$1;
