package repository

import (
	"time"
	"war_ticket/internal/domain"
	"war_ticket/internal/err"
	"war_ticket/internal/interfaces"
)

type TicketEventRepositoryImpl struct {
	TicketEvents map[int]domain.TicketEvent
	lastID       int
	now          time.Time
}

type TicketEventRepository interface {
	interfaces.Saver[domain.TicketEvent]
}

func NewTicketEventRepository() TicketEventRepository {
	return &TicketEventRepositoryImpl{
		TicketEvents: make(map[int]domain.TicketEvent),
		lastID:       0,
		now:          time.Now(),
	}
}

// Save implements TicketEventRepository.
func (t *TicketEventRepositoryImpl) Save(value *domain.TicketEvent) (*domain.TicketEvent, error) {
	t.lastID++

	value.ID = t.lastID
	value.CreatedAt = t.now.Format(time.DateTime)
	value.UpdatedAt = t.now.Format(time.DateTime)

	_, ok := t.TicketEvents[value.ID]

	if ok {
		return nil, err.ErrDuplicateID
	}

	t.TicketEvents[value.ID] = *value

	return value, nil
}
