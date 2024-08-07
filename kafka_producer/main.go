package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{"localhost:29092"}
	topic := "aw"

	// Create a new Sarama configuration
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 5                    // Retry up to 5 times to produce the message
	config.Producer.Return.Successes = true

	// Create a new synchronous producer
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start Sarama producer: %v", err)
	}
	defer producer.Close()

	// Sending messages to multiple partitions
	for i := 10; i < 20; i++ {
		partition := int32(rand.Intn(3)) // Assuming the topic has 3 partitions: 0, 1, and 2
		msg := &sarama.ProducerMessage{
			Topic: topic,
			// Partition: 1,
			Key:   sarama.StringEncoder(fmt.Sprintf("%d", 3)),
			Value: sarama.StringEncoder(fmt.Sprintf("message-%d", i)),
		}

		log.Printf("Key = %s, Value = %s", msg.Key, msg.Value)
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("Failed to send message to partition %d: %v", partition, err)
		} else {
			log.Printf("Message sent to partition %d at offset %d", partition, offset)
		}

		time.Sleep(500 * time.Millisecond) // To simulate delay
	}
}
