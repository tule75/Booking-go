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
