package repository

import (
	"sync"
	"time"
	"war_ticket/internal/domain"
	errr "war_ticket/internal/err"
	"war_ticket/internal/interfaces"
)

type OrderRepositoryImpl struct {
	Orders map[int]domain.Order
	lastID int
	mutex  sync.RWMutex
}

type OrderRepository interface {
	interfaces.Getter[domain.Order]
	interfaces.Saver[domain.Order]
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{
		Orders: make(map[int]domain.Order),
		lastID: 0,
		mutex:  sync.RWMutex{},
	}
}

// GetAll implements OrderRepository.
func (o *OrderRepositoryImpl) GetAll() []domain.Order {
	var orders []domain.Order

	for _, v := range o.Orders {
		orders = append(orders, v)
	}

	return orders
}

// Save implements OrderRepository.
func (o *OrderRepositoryImpl) Save(value *domain.Order) (*domain.Order, error) {
	// o.mutex.Lock()
	// defer o.mutex.Unlock()
	o.lastID++

	value.ID = o.lastID
	value.CreatedAt = time.Now().Format(time.DateTime)
	value.UpdatedAt = time.Now().Format(time.DateTime)

	_, ok := o.Orders[value.ID]

	if ok {
		return nil, errr.ErrDuplicateID
	}

	o.Orders[value.ID] = *value

	return value, nil
}