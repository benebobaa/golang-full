-- name: CreateOrder :one
INSERT INTO orders (customer_id, username, product_name, status) 
VALUES ($1, $2, $3, 'PENDING') RETURNING *;
