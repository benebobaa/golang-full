-- name: ListTickets :many
SELECT * FROM tickets;

-- name: CreateTicket :one
INSERT INTO tickets(name, stock, price)
VALUES ($1, $2, $3) RETURNING *; 

