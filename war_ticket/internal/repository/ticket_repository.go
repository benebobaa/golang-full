package repository

import (
	"sync"
	"time"
	"war_ticket/internal/domain"
	errr "war_ticket/internal/err"
	"war_ticket/internal/interfaces"
)

type TicketRepositoryImpl struct {
	Tickets map[int]domain.Ticket
	lastID  int
	mutex   sync.RWMutex
}
type TicketRepository interface {
	interfaces.Saver[domain.Ticket]
	interfaces.Getter[domain.Ticket]
	interfaces.Finder[domain.Ticket]
	interfaces.Updater[domain.Ticket]
}

func NewTicketRepository() TicketRepository {
	return &TicketRepositoryImpl{
		Tickets: make(map[int]domain.Ticket),
		lastID:  0,
		mutex:   sync.RWMutex{},
	}
}

// GetAll implements TicketRepository.
func (t *TicketRepositoryImpl) GetAll() []domain.Ticket {
	var tickets []domain.Ticket

	for _, v := range t.Tickets {
		tickets = append(tickets, v)
	}

	return tickets
}

// Save implements TicketRepository.
func (t *TicketRepositoryImpl) Save(value *domain.Ticket) (*domain.Ticket, error) {
	t.lastID++

	value.ID = t.lastID
	value.CreatedAt = time.Now().Format(time.DateTime)
	value.UpdatedAt = time.Now().Format(time.DateTime)

	_, ok := t.Tickets[value.ID]

	if ok {
		return nil, errr.ErrDuplicateID
	}

	t.Tickets[value.ID] = *value

	return value, nil

}

// FindByID implements TicketRepository.
func (t *TicketRepositoryImpl) FindByID(id int) (*domain.Ticket, error) {
	// t.mutex.Lock()
	// defer t.mutex.Unlock()
	value, ok := t.Tickets[id]

	if !ok {
		return nil, errr.ErrNotFound
	}

	return &value, nil
}

// Update implements TicketRepository.
func (t *TicketRepositoryImpl) Update(value *domain.Ticket) (*domain.Ticket, error) {
	// t.mutex.Lock()
	// defer t.mutex.Unlock()
	_, ok := t.Tickets[value.ID]

	if !ok {
		return nil, errr.ErrNotFound
	}

	t.Tickets[value.ID] = *value

	return value, nil
}
