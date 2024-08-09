package usecase

import (
	"context"
	"order-svc/internal/delivery/kafka"
	"order-svc/internal/dto"
	"order-svc/internal/repository/sqlc"
)

type OrderUsecase struct {
	queries           sqlc.Querier
	orchestraProducer *kafka.KafkaProducer
}

func NewOrderUsecase(queries sqlc.Querier, producer *kafka.KafkaProducer) *OrderUsecase {
	return &OrderUsecase{queries: queries, orchestraProducer: producer}
}

func (oc *OrderUsecase) CreateOrder(ctx context.Context, order *dto.OrderRequest) (*sqlc.Order, error) {

	orderCreated, err := oc.queries.CreateOrder(ctx, sqlc.CreateOrderParams{
		CustomerID:  order.CustomerID,
		Username:    order.Username,
		ProductName: order.ProductName,
	})

	if err != nil {
		return nil, err
	}

	return &orderCreated, nil
}
