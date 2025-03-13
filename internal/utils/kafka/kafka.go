package kafka

import (
	"context"
	"ecommerce_go/global"
	"ecommerce_go/internal/database"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type InsertAvailabilityKafka struct {
	RoomID string      `json:"room_id"`
	Days   []time.Time `json:"days"`
}

func SendInsertAvailibility(c context.Context, body *InsertAvailabilityKafka, callback func(error)) {

	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("JSON encoding error: %v", err)
		callback(fmt.Errorf("failed to encode Kafka message: %w", err))
	}
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	callback(global.KafkaProducer.WriteMessages(c, msg))
}

func ActionInsertAvailbility(c context.Context, body []byte, sqlc *database.Queries, callback func(error)) {
	var parseBody InsertAvailabilityKafka
	if err := json.Unmarshal(body, &parseBody); err != nil {
		callback(err)
		return
	}
	roomID := parseBody.RoomID
	records := make([]database.BulkInsertAvailabilityStruct, len(parseBody.Days))
	for i, v := range parseBody.Days {
		records[i] = database.BulkInsertAvailabilityStruct{
			ID:          uuid.New().String(),
			RoomID:      roomID,
			Date:        v,
			IsAvailable: true,
		}
	}

	err := sqlc.BulkInsertAvailability(c, records)
	callback(err)
}
