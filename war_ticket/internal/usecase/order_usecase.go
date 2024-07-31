package usecase

import (
	"context"
	"war_ticket/internal/domain"
	"war_ticket/internal/domain/dto"
	errr "war_ticket/internal/err"
	"war_ticket/internal/interfaces"
	"war_ticket/internal/middleware"
	"war_ticket/internal/repository"
	"war_ticket/internal/repository/sqlc"
)

type OrderUsecaseImpl struct {
	orderRepository  repository.OrderRepository
	ticketRepository repository.TicketRepository
	sqlcQueries      sqlc.Querier
}

type OrderUsecase interface {
	CreateOrder(ctx context.Context, value *dto.OrderRequest) (*domain.Order, error)
	interfaces.Getter[domain.Order]
}

func NewOrderUsecase(
	or repository.OrderRepository,
	tr repository.TicketRepository,
	sqlcQueries sqlc.Querier,
) OrderUsecase {
	return &OrderUsecaseImpl{
		orderRepository:  or,
		ticketRepository: tr,
		sqlcQueries:      sqlcQueries,
	}
}

// GetAll implements OrderUsecase.
func (o *OrderUsecaseImpl) GetAll() []domain.Order {
	result, _ := o.sqlcQueries.ListOrdersWithTickets(context.Background())
	return result
}

// Save implements OrderUsecase.
func (o *OrderUsecaseImpl) CreateOrder(ctx context.Context, value *dto.OrderRequest) (*domain.Order, error) {

	var tickets []domain.Ticket
	var totalPrice float64

	for _, v := range value.Tickets {
		var subTotal float64

		ticket, err := o.sqlcQueries.GetTicket(ctx, int32(v.TicketID))

		if err != nil {
			return nil, err
		}

		if ticket.Stock < v.Quantity {
			return nil, errr.ErrTicketOutOfStock
		}

		ticket.Stock = ticket.Stock - v.Quantity

		err = o.sqlcQueries.UpdateStock(ctx, sqlc.UpdateStockParams{
			ID:    int32(v.TicketID),
			Stock: int32(ticket.Stock),
		})

		if err != nil {
			return nil, err
		}

		ticket.Stock = v.Quantity
		tickets = append(tickets, ticket)

		subTotal = ticket.Price * float64(v.Quantity)
		totalPrice += subTotal
	}

	user, ok := ctx.Value(middleware.ContextUserKey).(*domain.User)

	if !ok {
		return nil, errr.ErrUserContextEmpty
	}

	order, err := o.sqlcQueries.CreateOrder(ctx, sqlc.CreateOrderParams{
		Customer:   value.Name,
		Username:   user.Username,
		TotalPrice: totalPrice,
	})

	if err != nil {
		return nil, err
	}

	for _, v := range tickets {
		err = o.sqlcQueries.CreateOrderTicket(ctx, sqlc.CreateOrderTicketParams{
			OrderID:  int32(order.ID),
			TicketID: int32(v.ID),
		})

		if err != nil {
			return nil, err
		}
	}

	return &domain.Order{
		ID:         int(order.ID),
		Customer:   order.Customer,
		Username:   order.Username,
		TotalPrice: order.TotalPrice,
		Tickets:    tickets,
		Common: domain.Common{
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
		},
	}, nil
}
