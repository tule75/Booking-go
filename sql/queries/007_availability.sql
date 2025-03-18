-- name: CreateAvailability :execresult
INSERT INTO `availability` (id, room_id, date, is_available)
VALUES (?, ?, ?, ?);


-- name: GetAvailabilityByID :one
SELECT 
  availability.id AS availability_id,
  availability.room_id,
  availability.date,
  availability.is_available
FROM `availability`
WHERE availability.id = ? AND availability.deleted_at IS NULL;

-- name: ListAvailabilityByRoomAndDateRange :many
SELECT 
  availability.id AS availability_id,
  availability.room_id,
  availability.date,
  availability.is_available
FROM availability
WHERE availability.room_id = ?
  AND availability.deleted_at IS NULL
  AND availability.date BETWEEN ? AND ?
ORDER BY availability.date ASC;

-- name: UpdateAvailability :exec
UPDATE availability
SET is_available = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: CheckRoomAvailability :one
SELECT COUNT(*) = 0 AS is_available
FROM availability
WHERE room_id = sqlc.arg(room_id)
  AND date BETWEEN sqlc.arg(check_in) AND sqlc.arg(check_out)
  AND is_available = FALSE;