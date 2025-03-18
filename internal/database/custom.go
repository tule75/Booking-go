package database

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
)

type BulkInsertAvailabilityStruct struct {
	ID          string
	RoomID      string
	Date        time.Time
	IsAvailable bool
}

type BulkUpdateAvailabilityStruct struct {
	RoomID      string
	CheckIn     time.Time
	CheckOut    time.Time
	IsAvailable bool
}

// BulkInsertAvailability dynamically inserts multiple records at once.
func (r *Queries) BulkInsertAvailability(ctx context.Context, records []BulkInsertAvailabilityStruct) error {
	if len(records) == 0 {
		return fmt.Errorf("no records to insert")
	}

	placeholders := []string{}
	args := []interface{}{}

	for _, r := range records {
		placeholders = append(placeholders, "(?, ?, ?, ?)")
		args = append(args, r.ID, r.RoomID, r.Date, r.IsAvailable)
	}

	query := fmt.Sprintf(
		"INSERT INTO availability (id, room_id, date, is_available) VALUES %s",
		strings.Join(placeholders, ","),
	)

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		log.Println("Bulk insert failed:", err)
		return err
	}

	log.Println("Bulk insert successful!")
	return nil
}

func (r *Queries) BulkUpdateAvailable(ctx context.Context, record BulkUpdateAvailabilityStruct) error {
	query := `
		UPDATE availability a
		JOIN (SELECT id FROM availability WHERE room_id = ? AND date BETWEEN ? AND ?) sub
		ON a.id = sub.id
		SET a.is_available = ?;
	`

	_, err := r.db.ExecContext(ctx, query, record.RoomID, record.CheckIn, record.CheckOut, record.IsAvailable)
	if err != nil {
		log.Println("Bulk update available failed:", err)
		return err
	}

	log.Println("Bulk update available successful!")
	return nil
}
