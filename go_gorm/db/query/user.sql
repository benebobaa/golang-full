
-- name: CreateUser :one
INSERT INTO users (username, email, password_hash, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW()) RETURNING *;