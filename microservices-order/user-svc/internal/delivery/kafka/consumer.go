package kafka

import (
	"context"
	"fmt"
	"time"
	"user-svc/internal/dto/event"
	"user-svc/internal/usecase"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
)

type KafkaConsumer struct {
	consumer sarama.ConsumerGroup
	topics   []string
	handler  sarama.ConsumerGroupHandler
}

type MessageHandler struct {
	userUsecase *usecase.Usecase
	producer    *KafkaProducer
}

func (h MessageHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h MessageHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h MessageHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for msg := range claim.Messages() {

		event, err := event.FromJSON(msg.Value)
		if err != nil {
			fmt.Printf("Error unmarshalling event: %v\n", err)
			continue
		}

		request, err := event.GetPayloadAsUserValidateRequest()
		if err != nil {
			fmt.Printf("Error converting payload: %v\n", err)
			continue
		}

		user, err := h.userUsecase.ValidateUser(request)
		if err != nil {
			fmt.Printf("Error validating user: %v\n", err)
			continue
		}

		event.EventType = "user_validation"
		event.Action = "validate"
		event.Timestamp = time.Now()
		event.Source = "user-svc"
		event.Payload = user

		if user.Error != "" {
			event.Status = "error"
		} else {
			event.Status = "success"
		}

		fmt.Printf("User validated: %v\n", user)

		eventBytes, err := event.ToJSON()

		err = h.producer.SendMessage(uuid.New().String(), eventBytes)

		if err != nil {
			fmt.Printf("Error sending message: %v\n", err)
		}

		sess.MarkMessage(msg, "")
	}
	return nil
}

func NewKafkaConsumer(
	brokers []string, groupID string,
	topics []string, usecase *usecase.Usecase,
	producer *KafkaProducer,
) (*KafkaConsumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		consumer: consumer,
		topics:   topics,
		handler: MessageHandler{
			userUsecase: usecase,
			producer:    producer,
		},
	}, nil
}

func (kc *KafkaConsumer) Consume(ctx context.Context) error {
	for {
		err := kc.consumer.Consume(ctx, kc.topics, kc.handler)
		if err != nil {
			return err
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

func (kc *KafkaConsumer) Close() error {
	return kc.consumer.Close()
}
