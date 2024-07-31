-- name: CreateOrder :one
INSERT INTO orders(customer, username, total_price)
VALUES ($1, $2, $3) RETURNING *;