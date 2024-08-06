-- name: ListTickets :many
SELECT * FROM tickets;

-- name: CreateTicket :one
INSERT INTO tickets(name, stock, price)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateStock :exec
UPDATE tickets SET stock = stock - $1 WHERE id = $2;

-- name: GetTicket :one
SELECT * FROM tickets WHERE id = $1;