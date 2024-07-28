package usecase

import (
	"war_ticket/internal/domain"
	"war_ticket/internal/domain/dto"
	errr "war_ticket/internal/err"
	"war_ticket/internal/interfaces"
	"war_ticket/internal/repository"
)

type OrderUsecaseImpl struct {
	orderRepository  repository.OrderRepository
	ticketRepository repository.TicketRepository
}

type OrderUsecase interface {
	CreateOrder(value *dto.OrderRequest) (*domain.Order, error)
	interfaces.Getter[domain.Order]
}

func NewOrderUsecase(
	or repository.OrderRepository,
	tr repository.TicketRepository,
) OrderUsecase {
	return &OrderUsecaseImpl{
		orderRepository:  or,
		ticketRepository: tr,
	}
}

// GetAll implements OrderUsecase.
func (o *OrderUsecaseImpl) GetAll() []domain.Order {
	return o.orderRepository.GetAll()
}

// Save implements OrderUsecase.
func (o *OrderUsecaseImpl) CreateOrder(value *dto.OrderRequest) (*domain.Order, error) {

	var tickets []domain.Ticket

	if len(value.Tickets) < 1 {
		return nil, errr.ErrTicketsRequestEmpty
	}

	for _, v := range value.Tickets {
		ticket, err := o.ticketRepository.FindByID(v.TicketID)

		if err != nil {
			return nil, err
		}

		if ticket.Stock < v.Quantity {
			return nil, errr.ErrTicketOutOfStock
		}

		ticket.Stock = ticket.Stock - v.Quantity

		_, err = o.ticketRepository.Update(ticket)
		if err != nil {
			return nil, err
		}

		ticket.Stock = v.Quantity
		tickets = append(tickets, *ticket)
	}

	result, err := o.orderRepository.Save(
		&domain.Order{
			Customer: value.Name,
			Tickets:  tickets,
		},
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}
