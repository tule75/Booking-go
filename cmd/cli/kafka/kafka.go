package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

var KafkaProducer *kafka.Writer

const (
	kafkaURL   = "localhost:19094"
	kafkaTopic = "user_topic_1"
)

// Producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
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

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg

	return &s
}

func ActionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))

	body := make(map[string]interface{})
	body["action"] = "action"
	body["type"] = s

	jsonBody, _ := json.Marshal(body)

	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	err := KafkaProducer.WriteMessages(c, msg)

	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "action Successfully",
	})
}

func RegisterConsumerATCS(id int) {
	kafkaGroupId := "consumer-group-"

	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer %v Hong Phien ATC::", id)

	for {
		m, err := reader.ReadMessage(context.Background())

		if err != nil {
			fmt.Printf("Error reading message from consumer: %v with Error: %v", id, err)
		}
		fmt.Printf("Consumer %v, hong topic %v, partition %v, offset %v, time %d %s = %s\n", id, m.Topic,
			m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))

	}
}

func main() {
	r := gin.Default()
	KafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer KafkaProducer.Close()

	r.POST("action/stock", ActionStock)

	go RegisterConsumerATCS(1)
	go RegisterConsumerATCS(2)

	r.Run(":8888")
}
