package initialize

import (
	"ecommerce_go/global"
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
		Partition:      10e3, //10KB
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
