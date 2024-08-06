// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: ticket_event.sql

package sqlc

import (
	"context"
)

const createTicketEvent = `-- name: CreateTicketEvent :exec
INSERT INTO ticket_events(event_id, ticket_id)
VALUES ($1, $2)
`

type CreateTicketEventParams struct {
	EventID  int32 `json:"event_id"`
	TicketID int32 `json:"ticket_id"`
}

func (q *Queries) CreateTicketEvent(ctx context.Context, arg CreateTicketEventParams) error {
	_, err := q.db.ExecContext(ctx, createTicketEvent, arg.EventID, arg.TicketID)
	return err
}

const listTicketsWithEvents = `-- name: ListTicketsWithEvents :many
SELECT id, event_id, ticket_id, created_at, updated_at FROM ticket_events
`

func (q *Queries) ListTicketsWithEvents(ctx context.Context) ([]TicketEvent, error) {
	rows, err := q.db.QueryContext(ctx, listTicketsWithEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TicketEvent{}
	for rows.Next() {
		var i TicketEvent
		if err := rows.Scan(
			&i.ID,
			&i.EventID,
			&i.TicketID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
