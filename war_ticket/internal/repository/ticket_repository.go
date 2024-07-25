package repository

import (
	"time"
	"war_ticket/internal/domain"
	"war_ticket/internal/err"
	"war_ticket/internal/interfaces"
)

type TicketRepositoryImpl struct {
	Tickets map[int]domain.Ticket
	lastID  int
	now     time.Time
}

type TicketRepository interface {
	interfaces.Saver[domain.Ticket]
	interfaces.Getaller[domain.Ticket]
}

func NewTicketRepository() TicketRepository {
	return &TicketRepositoryImpl{
		Tickets: make(map[int]domain.Ticket),
		lastID:  0,
		now:     time.Now(),
	}
}

// GetAll implements TicketRepository.
func (t *TicketRepositoryImpl) GetAll() []domain.Ticket {
	panic("unimplemented")
}

// Save implements TicketRepository.
func (t *TicketRepositoryImpl) Save(value *domain.Ticket) (*domain.Ticket, error) {
	t.lastID++

	value.ID = t.lastID
	value.CreatedAt = t.now.Format(time.DateTime)
	value.UpdatedAt = t.now.Format(time.DateTime)

	_, ok := t.Tickets[value.ID]

	if ok {
		return nil, err.ErrDuplicateID
	}

	t.Tickets[value.ID] = *value

	return value, nil

}
