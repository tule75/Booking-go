-- name: CreateBooking :execresult
INSERT INTO bookings (id, user_id, property_id, room_id, check_in, check_out, guests, total_price, status)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetBookingByID :one
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
WHERE bookings.id = ? AND bookings.deleted_at IS NULL;

-- name: ListBookingsByUser :many
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
  AND ( ? IS NULL OR bookings.status = ? )
  AND ( ? IS NULL OR bookings.check_in >= ? )
  AND ( ? IS NULL OR bookings.check_out <= ? )
ORDER BY bookings.check_in DESC
LIMIT ? OFFSET ?;

-- name: UpdateBooking :exec
UPDATE bookings
SET check_in = COALESCE(?, check_in), check_out = COALESCE(?, check_out), guests = COALESCE(?, guests), total_price = COALESCE(?, total_price), status = COALESCE(?, status)
WHERE id = ? AND deleted_at IS NULL;

-- name: SoftDeleteBooking :exec
UPDATE bookings SET deleted_at = NOW() WHERE id = ?;
