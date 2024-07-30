// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: ticket.sql

package sqlc

import (
	"context"
)

const createTicket = `-- name: CreateTicket :one
INSERT INTO tickets(name, stock, price)
VALUES ($1, $2, $3) RETURNING id, name, stock, price, created_at, updated_at
`

type CreateTicketParams struct {
	Name  string  `json:"name"`
	Stock int32   `json:"stock"`
	Price float64 `json:"price"`
}

func (q *Queries) CreateTicket(ctx context.Context, arg CreateTicketParams) (Ticket, error) {
	row := q.db.QueryRowContext(ctx, createTicket, arg.Name, arg.Stock, arg.Price)
	var i Ticket
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Stock,
		&i.Price,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTickets = `-- name: ListTickets :many
SELECT id, name, stock, price, created_at, updated_at FROM tickets
`

func (q *Queries) ListTickets(ctx context.Context) ([]Ticket, error) {
	rows, err := q.db.QueryContext(ctx, listTickets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Ticket{}
	for rows.Next() {
		var i Ticket
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Stock,
			&i.Price,
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
