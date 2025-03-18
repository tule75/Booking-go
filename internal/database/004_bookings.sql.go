// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: 004_bookings.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const cancelBooking = `-- name: CancelBooking :exec
UPDATE bookings SET status = 'Cancelled' WHERE id = ? and status <> 'completed'
`

func (q *Queries) CancelBooking(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, cancelBooking, id)
	return err
}

const createBooking = `-- name: CreateBooking :execresult
INSERT INTO bookings (id, user_id, property_id, room_id, check_in, check_out, guests, total_price, status)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateBookingParams struct {
	ID         string             `json:"id"`
	UserID     string             `json:"user_id"`
	PropertyID string             `json:"property_id"`
	RoomID     sql.NullString     `json:"room_id"`
	CheckIn    time.Time          `json:"check_in"`
	CheckOut   time.Time          `json:"check_out"`
	Guests     int32              `json:"guests"`
	TotalPrice string             `json:"total_price"`
	Status     NullBookingsStatus `json:"status"`
}

func (q *Queries) CreateBooking(ctx context.Context, arg CreateBookingParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createBooking,
		arg.ID,
		arg.UserID,
		arg.PropertyID,
		arg.RoomID,
		arg.CheckIn,
		arg.CheckOut,
		arg.Guests,
		arg.TotalPrice,
		arg.Status,
	)
}

const getBookingByID = `-- name: GetBookingByID :one
SELECT 
  bookings.id AS booking_id,
  bookings.user_id,
  bookings.property_id,
  bookings.room_id,
  bookings.check_in,
  bookings.check_out,
  bookings.guests,
  bookings.total_price,
  bookings.status,
  bookings.created_at
FROM bookings
WHERE bookings.id = ? AND bookings.deleted_at IS NULL
`

type GetBookingByIDRow struct {
	BookingID  string             `json:"booking_id"`
	UserID     string             `json:"user_id"`
	PropertyID string             `json:"property_id"`
	RoomID     sql.NullString     `json:"room_id"`
	CheckIn    time.Time          `json:"check_in"`
	CheckOut   time.Time          `json:"check_out"`
	Guests     int32              `json:"guests"`
	TotalPrice string             `json:"total_price"`
	Status     NullBookingsStatus `json:"status"`
	CreatedAt  sql.NullTime       `json:"created_at"`
}

func (q *Queries) GetBookingByID(ctx context.Context, id string) (GetBookingByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getBookingByID, id)
	var i GetBookingByIDRow
	err := row.Scan(
		&i.BookingID,
		&i.UserID,
		&i.PropertyID,
		&i.RoomID,
		&i.CheckIn,
		&i.CheckOut,
		&i.Guests,
		&i.TotalPrice,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const listBookingsByUser = `-- name: ListBookingsByUser :many
SELECT 
  bookings.id AS booking_id,
  bookings.user_id,
  bookings.property_id,
  bookings.room_id,
  bookings.check_in,
  bookings.check_out,
  bookings.guests,
  bookings.total_price,
  bookings.status,
  bookings.created_at
FROM bookings
WHERE bookings.user_id = ?
  AND bookings.deleted_at IS NULL
  AND (? IS NULL OR bookings.status = ?)
  AND (? IS NULL OR ? = FALSE OR bookings.check_in >= ?)
  AND (? IS NULL OR ? = FALSE OR bookings.check_out <= ?)
ORDER BY bookings.check_in DESC
LIMIT ? OFFSET ?
`

type ListBookingsByUserParams struct {
	UserID         string             `json:"user_id"`
	Status         NullBookingsStatus `json:"status"`
	FilterCheckIn  interface{}        `json:"filter_check_in"`
	CheckIn        time.Time          `json:"check_in"`
	FilterCheckOut interface{}        `json:"filter_check_out"`
	CheckOut       time.Time          `json:"check_out"`
	Limit          int32              `json:"limit"`
	Offset         int32              `json:"offset"`
}

type ListBookingsByUserRow struct {
	BookingID  string             `json:"booking_id"`
	UserID     string             `json:"user_id"`
	PropertyID string             `json:"property_id"`
	RoomID     sql.NullString     `json:"room_id"`
	CheckIn    time.Time          `json:"check_in"`
	CheckOut   time.Time          `json:"check_out"`
	Guests     int32              `json:"guests"`
	TotalPrice string             `json:"total_price"`
	Status     NullBookingsStatus `json:"status"`
	CreatedAt  sql.NullTime       `json:"created_at"`
}

func (q *Queries) ListBookingsByUser(ctx context.Context, arg ListBookingsByUserParams) ([]ListBookingsByUserRow, error) {
	rows, err := q.db.QueryContext(ctx, listBookingsByUser,
		arg.UserID,
		arg.Status,
		arg.Status,
		arg.FilterCheckIn,
		arg.FilterCheckIn,
		arg.CheckIn,
		arg.FilterCheckOut,
		arg.FilterCheckOut,
		arg.CheckOut,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListBookingsByUserRow
	for rows.Next() {
		var i ListBookingsByUserRow
		if err := rows.Scan(
			&i.BookingID,
			&i.UserID,
			&i.PropertyID,
			&i.RoomID,
			&i.CheckIn,
			&i.CheckOut,
			&i.Guests,
			&i.TotalPrice,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const softDeleteBooking = `-- name: SoftDeleteBooking :exec
UPDATE bookings SET deleted_at = NOW() WHERE id = ?
`

func (q *Queries) SoftDeleteBooking(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, softDeleteBooking, id)
	return err
}

const updateBooking = `-- name: UpdateBooking :exec
UPDATE bookings
SET check_in = COALESCE(?, check_in), check_out = COALESCE(?, check_out), guests = COALESCE(?, guests), total_price = COALESCE(?, total_price), status = COALESCE(?, status)
WHERE id = ? AND deleted_at IS NULL
`

type UpdateBookingParams struct {
	CheckIn    time.Time          `json:"check_in"`
	CheckOut   time.Time          `json:"check_out"`
	Guests     int32              `json:"guests"`
	TotalPrice string             `json:"total_price"`
	Status     NullBookingsStatus `json:"status"`
	ID         string             `json:"id"`
}

func (q *Queries) UpdateBooking(ctx context.Context, arg UpdateBookingParams) error {
	_, err := q.db.ExecContext(ctx, updateBooking,
		arg.CheckIn,
		arg.CheckOut,
		arg.Guests,
		arg.TotalPrice,
		arg.Status,
		arg.ID,
	)
	return err
}
