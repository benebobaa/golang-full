-- name: CreateUser :one
INSERT INTO users(name, email, password)
VALUES ($1, $2, $3) RETURNING *;

-- name: FindByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: CountByEmail :one
SELECT COUNT(*) FROM users WHERE email = $1;
