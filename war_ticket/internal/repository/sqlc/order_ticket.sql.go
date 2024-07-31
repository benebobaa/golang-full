// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: order_ticket.sql

package sqlc

import (
	"context"
)

const createOrderTicket = `-- name: CreateOrderTicket :exec
INSERT INTO order_tickets(order_id, ticket_id)
VALUES ($1,$2)
`

type CreateOrderTicketParams struct {
	OrderID  int32 `json:"order_id"`
	TicketID int32 `json:"ticket_id"`
}

func (q *Queries) CreateOrderTicket(ctx context.Context, arg CreateOrderTicketParams) error {
	_, err := q.db.ExecContext(ctx, createOrderTicket, arg.OrderID, arg.TicketID)
	return err
}
