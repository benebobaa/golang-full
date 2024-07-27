package repository

import (
	"time"
	"war_ticket/internal/domain"
	errr "war_ticket/internal/err"
	"war_ticket/internal/interfaces"
)

type TicketEventRepositoryImpl struct {
	TicketEvents map[int]domain.TicketEvent
	lastID       int
}

type TicketEventRepository interface {
	interfaces.Saver[domain.TicketEvent]
	interfaces.Getter[domain.TicketEvent]
}

func NewTicketEventRepository() TicketEventRepository {
	return &TicketEventRepositoryImpl{
		TicketEvents: make(map[int]domain.TicketEvent),
		lastID:       0,
	}
}

// Save implements TicketEventRepository.
func (t *TicketEventRepositoryImpl) Save(value *domain.TicketEvent) (*domain.TicketEvent, error) {
	t.lastID++

	value.ID = t.lastID
	value.CreatedAt = time.Now().Format(time.DateTime)
	value.UpdatedAt = time.Now().Format(time.DateTime)

	_, ok := t.TicketEvents[value.ID]

	if ok {
		return nil, errr.ErrDuplicateID
	}

	t.TicketEvents[value.ID] = *value

	return value, nil
}

// GetAll implements TicketEventRepository.
func (t *TicketEventRepositoryImpl) GetAll() []domain.TicketEvent {
	var ticketEvents []domain.TicketEvent

	for _, v := range t.TicketEvents {
		ticketEvents = append(ticketEvents, v)
	}

	return ticketEvents
}
