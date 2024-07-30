package db_repo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"war_ticket/internal/domain/dto"
)

type TicketRepositoryImpl struct {
	DB *sql.DB
}

type TicketRepository interface {
	ListTicketsWithEvents() ([]dto.TicketEventResponse, error)
}

func NewTicketRepository(db *sql.DB) TicketRepository {
	return &TicketRepositoryImpl{
		DB: db,
	}
}

// ListTicketsWithEvents implements TicketRepository.
func (t *TicketRepositoryImpl) ListTicketsWithEvents() ([]dto.TicketEventResponse, error) {
	query := `SELECT * FROM ticket_event_view`
	rows, err := t.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying ticket_event_view: %w", err)
	}
	defer rows.Close()

	var results []dto.TicketEventResponse

	for rows.Next() {
		var response dto.TicketEventResponse
		var ticketsJSON []byte

		err := rows.Scan(
			&response.Event.ID,
			&response.Event.Name,
			&response.Event.Location,
			&response.Event.CreatedAt,
			&response.Event.UpdatedAt,
			&ticketsJSON,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		err = json.Unmarshal(ticketsJSON, &response.Tickets)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling tickets JSON: %w", err)
		}

		results = append(results, response)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return results, nil
}
