package kafka

import (
	"log"
	"orchestra-svc/internal/dto/event"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
)

type MessageHandler struct {
	userProducer    *KafkaProducer
	productProducer *KafkaProducer
}

func NewMessageHandler(userProducer *KafkaProducer, productProducer *KafkaProducer) *MessageHandler {
	return &MessageHandler{userProducer: userProducer, productProducer: productProducer}
}

func (h MessageHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h MessageHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h MessageHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for msg := range claim.Messages() {

		log.Println("event: ", string(msg.Value))

		event := event.GlobalEvent[event.Order]{}

		err := event.FromJSON(msg.Value)

		if err != nil {
			log.Printf("Error unmarshalling event: %v", err)
		}

		log.Println("event: ", event)

		if event.Source == "order-svc" {
			log.Println("order event")
			eventBytes, err := event.ToJSON()
			if err != nil {
				log.Printf("Error marshalling event: %s", err.Error())
			}
			err = h.userProducer.SendMessage(uuid.New().String(), eventBytes)

			if err != nil {
				log.Printf("Error sending message: %v", err)
			}
		}

		if event.Source == "user-svc" {

			err := h.productProducer.SendMessage(uuid.New().String(), msg.Value)

			if err != nil {
				log.Printf("Error sending message: %v", err.Error())
			}
		}

		sess.MarkMessage(msg, "")
	}
	return nil
}
