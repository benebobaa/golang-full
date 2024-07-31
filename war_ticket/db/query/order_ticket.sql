-- name: CreateOrderTicket :exec
INSERT INTO order_tickets(order_id, ticket_id)
VALUES ($1,$2);