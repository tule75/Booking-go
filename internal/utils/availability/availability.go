package availability

import (
	"context"
	"ecommerce_go/global"
	"ecommerce_go/internal/database"
	"ecommerce_go/internal/utils/kafka"
	constant "ecommerce_go/pkg"
	"sync"
	"time"

	"go.uber.org/zap"
)

// Generate availability cho các phòng mới từ nay đến cuối 3 tháng sau
func GenerateAvailabilityForNewRooms(ctx context.Context, sqlc *database.Queries, roomID string, wait *sync.WaitGroup) error {
	defer wait.Done()
	now := time.Now()
	endDate := now.AddDate(0, 3, 0) // 3 tháng sau
	firstDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	func(roomID string) {

		record := kafka.InsertAvailabilityKafka{RoomID: roomID, Days: []time.Time{}}

		// Duyệt qua từng ngày từ hôm nay đến cuối 3 tháng sau
		for date := firstDay; date.Before(endDate); date = date.AddDate(0, 0, 1) {
			record.Days = append(record.Days, date)
		}

		// Gửi dữ liệu đến Kafka
		kafka.SendInsertAvailibility(ctx, &record, func(err error) {
			if err != nil {
				global.Logger.Error(constant.KafkaFailure, zap.Error(err))
			} else {
				global.Logger.Info(constant.KafkaSuccess)
			}
		})

	}(roomID)

	global.Logger.Info("Availability inserted successfully for new rooms", zap.Any("ID:", roomID))
	return nil
}
