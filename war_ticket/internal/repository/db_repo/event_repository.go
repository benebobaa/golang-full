package db_repo

import (
	"context"
	"database/sql"
	"war_ticket/internal/domain"
	"war_ticket/internal/interfaces"
)

type EventRepositoryImpl struct {
	DB *sql.DB
}

type EventRepository interface {
	interfaces.Getter[domain.Event]
	interfaces.Saver[domain.Event]
	interfaces.Finder[domain.Event]
}

func NewEventRepository(db *sql.DB) EventRepository {
	return &EventRepositoryImpl{
		DB: db,
	}
}

// GetAll implements EventRepository.
func (e *EventRepositoryImpl) GetAll() []domain.Event {
	var events []domain.Event

	query := `SELECT * FROM events`

	rows, err := e.DB.Query(query)

	if err != nil {
		return events
	}

	var event domain.Event
	for rows.Next() {
		rows.Scan(&event.ID, &event.Name, &event.Location, &event.CreatedAt, &event.UpdatedAt)

		events = append(events, event)
	}

	return events
}

// Save implements EventRepository.
func (e *EventRepositoryImpl) Save(ctx context.Context, value *domain.Event) (*domain.Event, error) {

	query := `INSERT INTO events(name,location) VALUES ($1, $2) RETURNING id`

	err := e.DB.QueryRowContext(ctx, query, value.Name, value.Location).Scan(&value.ID)

	if err != nil {
		return nil, err
	}

	return value, nil
}

// FindByID implements EventRepository.
func (e *EventRepositoryImpl) FindByID(id int) (*domain.Event, error) {
	var event domain.Event

	query := `SELECT * FROM events WHERE id = $1`

	row := e.DB.QueryRow(query, id)

	err := row.Scan(&event.ID, &event.Name, &event.Location, &event.CreatedAt, &event.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &event, nil
}
