-- name: CreateUser :one
INSERT INTO users (id, name, phone, email, password_hash, role, token, created_at, updated_at)
VALUES (gen_random_uuid(), $1, $2, $3, $4, COALESCE($5, 'customer'), NULL, NOW(), NOW())
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByPhone :one
SELECT * FROM users WHERE phone = $1;

-- name: UpdateUserToken :exec
UPDATE users
SET token = $2, last_login = NOW(), updated_at = NOW()
WHERE id = $1;

-- name: ClearUserToken :exec
UPDATE users
SET token = NULL, updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserProfile :exec
UPDATE users
SET name = $2, email = $3, updated_at = NOW()
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
