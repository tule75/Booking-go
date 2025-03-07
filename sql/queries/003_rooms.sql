-- name: CreateRoom :exec
INSERT INTO rooms (id, property_id, name, price, max_guests, is_available)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetRoomByID :one
SELECT 
  rooms.id AS room_id,
  rooms.property_id,
  rooms.name,
  rooms.price,
  rooms.max_guests,
  rooms.is_available,
  rooms.created_at
FROM rooms
WHERE rooms.id = ? AND rooms.deleted_at IS NULL;

-- name: ListRoomsByProperty :many
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
LIMIT ? OFFSET ?;

-- name: UpdateRoom :exec
UPDATE rooms
SET name = ?, price = ?, max_guests = ?, is_available = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: SoftDeleteRoom :exec
UPDATE rooms SET deleted_at = NOW() WHERE id = ?;
