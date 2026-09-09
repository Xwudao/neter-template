-- name: ListUsers :many
SELECT id, username, password, role FROM users ORDER BY id;

-- name: GetUser :one
SELECT id, username, password, role FROM users WHERE id = $1;

-- name: GetUserByUsername :one
SELECT id, username, password, role FROM users WHERE username = $1;

-- name: CreateUser :one
INSERT INTO users (username, password, role)
VALUES ($1, $2, $3)
RETURNING id, username, password, role;

-- name: DeleteUser :execrows
DELETE FROM users WHERE id = $1;
