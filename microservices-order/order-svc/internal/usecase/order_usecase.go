package usecase

import (
	"context"
	"order-svc/internal/delivery/kafka"
	"order-svc/internal/dto"
	"order-svc/internal/dto/event"
	"order-svc/internal/repository/sqlc"
	"time"

	"github.com/google/uuid"
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

	orderCreatedEvent, err := event.GlobalEvent[sqlc.Order]{
		EventID:   uuid.New().String(),
		EventType: "order",
		Timestamp: time.Now(),
		Source:    "order-svc",
		Action:    "create",
		Status:    "success",
		Payload:   orderCreated,
	}.ToJSON()

	if err != nil {
		return nil, err
	}

	err = oc.orchestraProducer.SendMessage(uuid.New().String(), orderCreatedEvent)

	if err != nil {
		return nil, err
	}

	return &orderCreated, nil
}
