// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: 003_rooms.sql

package database

import (
	"context"
	"database/sql"
)

const createRoom = `-- name: CreateRoom :exec
INSERT INTO rooms (id, property_id, name, price, max_guests, is_available)
VALUES (?, ?, ?, ?, ?, ?)
`

type CreateRoomParams struct {
	ID          string
	PropertyID  string
	Name        sql.NullString
	Price       string
	MaxGuests   int32
	IsAvailable sql.NullBool
}

func (q *Queries) CreateRoom(ctx context.Context, arg CreateRoomParams) error {
	_, err := q.db.ExecContext(ctx, createRoom,
		arg.ID,
		arg.PropertyID,
		arg.Name,
		arg.Price,
		arg.MaxGuests,
		arg.IsAvailable,
	)
	return err
}

const getRoomByID = `-- name: GetRoomByID :one
SELECT 
  rooms.id AS room_id,
  rooms.property_id,
  rooms.name,
  rooms.price,
  rooms.max_guests,
  rooms.is_available,
  rooms.created_at
FROM rooms
WHERE rooms.id = ? AND rooms.deleted_at IS NULL
`

type GetRoomByIDRow struct {
	RoomID      string
	PropertyID  string
	Name        sql.NullString
	Price       string
	MaxGuests   int32
	IsAvailable sql.NullBool
	CreatedAt   sql.NullTime
}

func (q *Queries) GetRoomByID(ctx context.Context, id string) (GetRoomByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getRoomByID, id)
	var i GetRoomByIDRow
	err := row.Scan(
		&i.RoomID,
		&i.PropertyID,
		&i.Name,
		&i.Price,
		&i.MaxGuests,
		&i.IsAvailable,
		&i.CreatedAt,
	)
	return i, err
}

const listRoomsByProperty = `-- name: ListRoomsByProperty :many
SELECT 
  rooms.id AS room_id,
  rooms.property_id,
  rooms.name,
  rooms.price,
  rooms.max_guests,
  rooms.is_available,
  rooms.created_at
FROM rooms
WHERE rooms.property_id = ? AND rooms.deleted_at IS NULL
ORDER BY rooms.created_at DESC
LIMIT ? OFFSET ?
`

type ListRoomsByPropertyParams struct {
	PropertyID string
	Limit      int32
	Offset     int32
}

type ListRoomsByPropertyRow struct {
	RoomID      string
	PropertyID  string
	Name        sql.NullString
	Price       string
	MaxGuests   int32
	IsAvailable sql.NullBool
	CreatedAt   sql.NullTime
}

func (q *Queries) ListRoomsByProperty(ctx context.Context, arg ListRoomsByPropertyParams) ([]ListRoomsByPropertyRow, error) {
	rows, err := q.db.QueryContext(ctx, listRoomsByProperty, arg.PropertyID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListRoomsByPropertyRow
	for rows.Next() {
		var i ListRoomsByPropertyRow
		if err := rows.Scan(
			&i.RoomID,
			&i.PropertyID,
			&i.Name,
			&i.Price,
			&i.MaxGuests,
			&i.IsAvailable,
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

const softDeleteRoom = `-- name: SoftDeleteRoom :exec
UPDATE rooms SET deleted_at = NOW() WHERE id = ?
`

func (q *Queries) SoftDeleteRoom(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, softDeleteRoom, id)
	return err
}

const updateRoom = `-- name: UpdateRoom :exec
UPDATE rooms
SET name = ?, price = ?, max_guests = ?, is_available = ?
WHERE id = ? AND deleted_at IS NULL
`

type UpdateRoomParams struct {
	Name        sql.NullString
	Price       string
	MaxGuests   int32
	IsAvailable sql.NullBool
	ID          string
}

func (q *Queries) UpdateRoom(ctx context.Context, arg UpdateRoomParams) error {
	_, err := q.db.ExecContext(ctx, updateRoom,
		arg.Name,
		arg.Price,
		arg.MaxGuests,
		arg.IsAvailable,
		arg.ID,
	)
	return err
}
