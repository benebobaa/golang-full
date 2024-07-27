package repository

import (
	"time"
	"war_ticket/internal/domain"
	"war_ticket/internal/err"
	"war_ticket/internal/interfaces"
)

type EventRepositoryImpl struct {
	Events map[int]domain.Event
	lastID int
}

type EventRepository interface {
	interfaces.Getter[domain.Event]
	interfaces.Saver[domain.Event]
	interfaces.Finder[domain.Event]
}

func NewEventRepository() EventRepository {
	return &EventRepositoryImpl{
		Events: make(map[int]domain.Event),
		lastID: 0,
	}
}

// GetAll implements EventRepository.
func (e *EventRepositoryImpl) GetAll() []domain.Event {
	var events []domain.Event

	for _, v := range e.Events {
		events = append(events, v)
	}

	return events
}

// Save implements EventRepository.
func (e *EventRepositoryImpl) Save(value *domain.Event) (*domain.Event, error) {
	e.lastID++

	value.ID = e.lastID
	value.CreatedAt = time.Now().Format(time.DateTime)
	value.UpdatedAt = time.Now().Format(time.DateTime)

	_, ok := e.Events[value.ID]

	if ok {
		return nil, errr.ErrDuplicateID
	}

	e.Events[value.ID] = *value

	return value, nil
}

// FindByID implements EventRepository.
func (e *EventRepositoryImpl) FindByID(id int) (*domain.Event, error) {

	value, ok := e.Events[id]

	if !ok {
		return nil, errr.ErrNotFound
	}

	return &value, nil
}
