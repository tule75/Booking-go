package initialize

import (
	"context"
	"ecommerce_go/global"
	"ecommerce_go/internal/database"
	utilsKafka "ecommerce_go/internal/utils/kafka"
	constant "ecommerce_go/pkg"
	"fmt"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

// Producer
func getKafkaProducer(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// Consumer
func getKafkaReader(kafkaURL, topic, GroupID string) *kafka.Reader {
	broker := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        broker,
		GroupID:        GroupID,
		Topic:          topic,
		MaxBytes:       10e6, //10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		global.Logger.Fatal("Fail to close Producer", zap.Error(err))
	}
}

func InitKafka() {
	fmt.Println(global.Config.Kafka.Addr)
	global.KafkaProducer = getKafkaProducer(global.Config.Kafka.Addr, global.Config.Kafka.Topic)

	fmt.Println("Init Kafka success")

}

// RegisterConsumer runs the Kafka consumer
func RegisterConsumer(ctx context.Context, groupID string, sqlc *database.Queries) {
	reader := getKafkaReader(global.Config.Kafka.Addr, global.Config.Kafka.Topic, groupID)

	defer func() {
		reader.Close()
		global.Logger.Info("Kafka consumer stopped", zap.String("GroupID", groupID))
	}()

	for {
		select {
		case <-ctx.Done():
			global.Logger.Info("Kafka consumer shutting down...", zap.String("GroupID", groupID))
			return
		default:
			// Read next Kafka message
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				global.Logger.Error("Error reading message", zap.String("GroupID", groupID), zap.Error(err))
				time.Sleep(2 * time.Second) // Retry delay before reading again
				continue
			}

			// Process message in a separate goroutine (non-blocking)
			go func(m kafka.Message) {
				key := string(m.Key)
				switch key {
				case constant.KeyInsertAvailability:
					utilsKafka.ActionInsertAvailbility(ctx, m.Value, sqlc, func(err error) {
						if err != nil {
							global.Logger.Error(fmt.Sprintf(constant.ConsumerFailure, "Insert Availability"), zap.Error(err))
						} else {
							global.Logger.Info(fmt.Sprintf(constant.ConsumerSuccess, "Insert Availability"))
							// Commit Kafka message to prevent reprocessing
							if commitErr := reader.CommitMessages(ctx, m); commitErr != nil {
								global.Logger.Error(constant.FailCommitKafkaMessage, zap.Error(commitErr))
							}
						}
					})
				case constant.KeyInsertBooking:
					utilsKafka.ActionInsertBooking(ctx, m.Value, sqlc, func(err error) {
						if err != nil {
							global.Logger.Error(fmt.Sprintf(constant.ConsumerFailure, "Insert Booking"), zap.Error(err))
						} else {
							global.Logger.Info(fmt.Sprintf(constant.ConsumerSuccess, "Insert Booking"))
							// Commit Kafka message to prevent reprocessing
							if commitErr := reader.CommitMessages(ctx, m); commitErr != nil {
								global.Logger.Error(constant.FailCommitKafkaMessage, zap.Error(commitErr))
							}
						}
					})
				case constant.KeyUpdateAvailability:
					utilsKafka.ActionInsertAvailbility(ctx, m.Value, sqlc, func(err error) {
						if err != nil {
							global.Logger.Error(fmt.Sprintf(constant.ConsumerFailure, "update Availability"), zap.Error(err))
						} else {
							global.Logger.Info(fmt.Sprintf(constant.ConsumerSuccess, "update Availability"))
							// Commit Kafka message to prevent reprocessing
							if commitErr := reader.CommitMessages(ctx, m); commitErr != nil {
								global.Logger.Error(constant.FailCommitKafkaMessage, zap.Error(commitErr))
							}
						}
					})
				default:
					global.Logger.Warn("Received unknown Kafka message", zap.String("Key", key), zap.String("Value", string(m.Value)))
				}
			}(m)
		}
	}
}
