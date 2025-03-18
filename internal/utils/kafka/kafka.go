package kafka

import (
	"context"
	"ecommerce_go/global"
	"ecommerce_go/internal/database"
	"ecommerce_go/internal/utils/redis"
	constant "ecommerce_go/pkg"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type InsertAvailabilityKafka struct {
	RoomID string      `json:"room_id"`
	Days   []time.Time `json:"days"`
}
type UpdateAvailabilityKafka struct {
	RoomID      string    `json:"room_id"`
	CheckIn     time.Time `json:"check_in"`
	CheckOut    time.Time `json:"check_out"`
	IsAvailable bool      `json:"is_available"`
}

func SendInsertAvailibility(c context.Context, body *InsertAvailabilityKafka, callback func(error)) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		callback(fmt.Errorf("%s: %w", constant.JSONEncodeError, err))
	}
	msg := kafka.Message{
		Key:   []byte(constant.KeyInsertAvailability),
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

func SendUpdateAvailibility(c context.Context, body *UpdateAvailabilityKafka, callback func(error)) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		callback(fmt.Errorf("%s: %w", constant.JSONEncodeError, err))
	}
	msg := kafka.Message{
		Key:   []byte(constant.KeyInsertBooking),
		Value: []byte(jsonBody),
	}

	callback(global.KafkaProducer.WriteMessages(c, msg))
}

func ActionUpdateAvailbility(c context.Context, body []byte, sqlc *database.Queries, callback func(error)) {
	var parseBody UpdateAvailabilityKafka
	if err := json.Unmarshal(body, &parseBody); err != nil {
		callback(err)
		return
	}
	record := database.BulkUpdateAvailabilityStruct{
		RoomID:      parseBody.RoomID,
		CheckIn:     parseBody.CheckIn,
		CheckOut:    parseBody.CheckOut,
		IsAvailable: parseBody.IsAvailable,
	}

	err := sqlc.BulkUpdateAvailable(c, record)
	callback(err)
}

func SendInsertBooking(c context.Context, body *database.CreateBookingParams, callback func(error)) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		callback(fmt.Errorf("%s: %w", constant.JSONEncodeError, err))
	}
	msg := kafka.Message{
		Key:   []byte(constant.KeyUpdateAvailability),
		Value: []byte(jsonBody),
	}

	redis.SoftLock(body.RoomID.String, body.UserID, body.CheckIn, body.CheckOut)

	callback(global.KafkaProducer.WriteMessages(c, msg))
}

func ActionInsertBooking(c context.Context, body []byte, sqlc *database.Queries, callback func(error)) {
	var parseBody database.CreateBookingParams
	if err := json.Unmarshal(body, &parseBody); err != nil {
		callback(err)
		return
	}

	x := database.CheckRoomAvailabilityParams{
		RoomID:   parseBody.RoomID.String,
		CheckIn:  parseBody.CheckIn,
		CheckOut: parseBody.CheckOut,
	}
	isAvailable, err := sqlc.CheckRoomAvailability(c, x)
	if err != nil {
		callback(err)
		return
	}

	if isAvailable {
		_, err := sqlc.CreateBooking(c, parseBody)
		if err != nil {
			global.Logger.Error("Error when creating Booking: ", zap.Error(err))
			callback(err)
			return
		}

		global.Logger.Info("Created Booking Success with id: ", zap.String("ID:", parseBody.ID))
		callback(err)
	}
}
