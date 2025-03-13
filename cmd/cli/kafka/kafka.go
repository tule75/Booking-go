package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	broker := "localhost:9092" // Change to "localhost:9092" if outside Docker
	topic := "bookinggo"

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})

	defer writer.Close()

	fmt.Println("🚀 Sending 100 messages to Kafka...")

	for i := 1; i <= 100; i++ {
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("key-%d", i)),
			Value: []byte(fmt.Sprintf(`{"type": "test", "id": %d, "content": "Message %d"}`, i, i)),
		}

		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Fatalf("❌ Failed to send message: %v", err)
		}
		fmt.Printf("✅ Sent message %d\n", i)
		time.Sleep(10 * time.Millisecond) // Prevent overloading Kafka
	}

	fmt.Println("🎯 Successfully sent 100 messages to Kafka!")
}
