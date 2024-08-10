package messaging

import (
	"fmt"
	"orchestra-svc/internal/usecase"

	"github.com/IBM/sarama"
)

type MessageHandler struct {
	usecase *usecase.OrderUsecase
}

func NewMessageHandler(u *usecase.OrderUsecase) *MessageHandler {
	return &MessageHandler{usecase: u}
}

func (h MessageHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h MessageHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h MessageHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for msg := range claim.Messages() {

		fmt.Println("Received message", string(msg.Value))
		sess.MarkMessage(msg, "")
	}
	return nil
}
