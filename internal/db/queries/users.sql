-- name: CreateUser :one
INSERT INTO users (name, phone, email, password_hash, role, address, created_at, updated_at)
VALUES ($1, $2, $3, $4, COALESCE($5, 'customer'), $6, NOW(), NOW())
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByPhone :one
SELECT * FROM users WHERE phone = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: UpdateUserProfile :exec
UPDATE users
SET name = $2, email = $3, phone = $4, address = $5, updated_at = NOW()
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
