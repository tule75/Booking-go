package cronjob

import (
	"context"
	"ecommerce_go/global"
	"ecommerce_go/internal/database"
	"ecommerce_go/internal/utils/kafka"
	"fmt"
	"log"
	"sync"
	"time"

	"go.uber.org/zap"
)

func StartMonthlyTask(sqlc *database.Queries) {
	go func() {
		for {
			now := time.Now()
			firstDayNextMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.Local)

			sleepDuration := time.Until(firstDayNextMonth)
			log.Printf("Scheduled availability insertion in %v", sleepDuration)

			time.Sleep(sleepDuration)

			GenerateAvailabilityForNewMonth(context.Background(), sqlc)
		}
	}()
}

func GenerateAvailabilityForNewMonth(ctx context.Context, sqlc *database.Queries) error {
	now := time.Now()
	nextMonth := now.AddDate(0, 3, 0)
	firstDayNextMonth := time.Date(nextMonth.Year(), nextMonth.Month(), 1, 0, 0, 0, 0, time.UTC)

	roomIDs, err := sqlc.GetAllRoomIds(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch room IDs: %v", err)
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 20)
	for _, roomID := range roomIDs {
		wg.Add(1)

		go func(roomID string) {
			defer wg.Done()
			semaphore <- struct{}{}

			record := kafka.InsertAvailabilityKafka{RoomID: roomID, Days: []time.Time{}}
			for i := 0; i < daysInMonth(firstDayNextMonth.Year(), firstDayNextMonth.Month()); i++ {
				record.Days = append(record.Days, firstDayNextMonth.AddDate(0, 0, i))
			}

			kafka.SendInsertAvailibility(ctx, &record, func(err error) {
				if err != nil {
					global.Logger.Error("Kafka send failed:", zap.Error(err))
				} else {
					global.Logger.Info("Kafka successfully received the availability data")
				}
			})

			<-semaphore
		}(roomID)
	}

	wg.Wait()

	global.Logger.Info("Availability inserted successfully for the new month for all rooms", zap.Any("IDs:", roomIDs))
	return nil
}

func daysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
