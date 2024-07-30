-- name: CreateTicketEvent :exec
INSERT INTO ticket_events(event_id, ticket_id)
VALUES ($1, $2);

-- name: ListTicketsWithEvents :many
SELECT * FROM ticket_events;
