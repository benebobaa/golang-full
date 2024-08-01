package usecase

import (
	"context"
	"database/sql"
	"war_ticket/internal/domain"
	"war_ticket/internal/domain/dto"
	errr "war_ticket/internal/err"
	"war_ticket/internal/interfaces"
	"war_ticket/internal/repository"
	"war_ticket/internal/repository/sqlc"
)

type OrderUsecaseImpl struct {
	orderRepository  repository.OrderRepository
	ticketRepository repository.TicketRepository
	DB               *sql.DB
	sqlcQueries      *sqlc.Queries
}

type OrderUsecase interface {
	CreateOrder(ctx context.Context, value *dto.OrderRequest, user *domain.User) (*domain.Order, error)
	interfaces.Getter[domain.Order]
}

func NewOrderUsecase(
	or repository.OrderRepository,
	tr repository.TicketRepository,
	db *sql.DB,
	sqlcQueries *sqlc.Queries,
) OrderUsecase {
	return &OrderUsecaseImpl{
		orderRepository:  or,
		ticketRepository: tr,
		DB:               db,
		sqlcQueries:      sqlcQueries,
	}
}

// GetAll implements OrderUsecase.
func (o *OrderUsecaseImpl) GetAll() []domain.Order {
	result, _ := o.sqlcQueries.ListOrdersWithTickets(context.Background())
	return result
}

// Save implements OrderUsecase.
func (o *OrderUsecaseImpl) CreateOrder(ctx context.Context, value *dto.OrderRequest, user *domain.User) (*domain.Order, error) {

	var err error
	var tickets []domain.Ticket
	var totalPrice float64

	tx, err := o.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})

	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	for _, v := range value.Tickets {
		var subTotal float64

		ticket, err := o.sqlcQueries.WithTx(tx).GetTicket(ctx, int32(v.TicketID))

		if err != nil {
			return nil, err
		}

		if ticket.Stock < v.Quantity {
			return nil, errr.ErrTicketOutOfStock
		}

		//ticket.Stock = ticket.Stock - v.Quantity

		err = o.sqlcQueries.WithTx(tx).UpdateStock(ctx, sqlc.UpdateStockParams{
			ID:    int32(ticket.ID),
			Stock: int32(v.Quantity),
		})

		if err != nil {
			return nil, err
		}

		ticket.Stock = v.Quantity
		tickets = append(tickets, ticket)

		subTotal = ticket.Price * float64(v.Quantity)
		totalPrice += subTotal
	}

	order, err := o.sqlcQueries.WithTx(tx).CreateOrder(ctx, sqlc.CreateOrderParams{
		Customer:   value.Name,
		Username:   user.Username,
		TotalPrice: totalPrice,
	})

	if err != nil {
		return nil, err
	}

	for _, v := range tickets {
		err = o.sqlcQueries.WithTx(tx).CreateOrderTicket(ctx, sqlc.CreateOrderTicketParams{
			OrderID:  int32(order.ID),
			TicketID: int32(v.ID),
		})

		if err != nil {
			return nil, err
		}
	}

	tx.Commit()

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
