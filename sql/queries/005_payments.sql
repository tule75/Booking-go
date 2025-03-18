-- name: CreatePayment :execresult
INSERT INTO payments (id, booking_id, user_id, amount, payment_method, status, stripe_payment_intent_id)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetPaymentByID :one
SELECT 
  payments.id AS payment_id,
  payments.booking_id,
  payments.user_id,
  payments.amount,
  payments.payment_method,
  payments.status,
  payments.stripe_payment_intent_id,
  payments.created_at
FROM payments
WHERE payments.id = ? AND payments.deleted_at IS NULL;

-- name: ListPaymentsByUser :many
SELECT 
  payments.id AS payment_id,
  payments.booking_id,
  payments.user_id,
  payments.amount,
  payments.payment_method,
  payments.status,
  payments.stripe_payment_intent_id,
  payments.created_at
FROM payments
WHERE payments.user_id = ? 
  AND payments.deleted_at IS NULL
  AND ( sqlc.arg(payment) IS NULL OR payments.payment_method = sqlc.arg(payment) )
ORDER BY payments.created_at DESC
LIMIT ? OFFSET ?;

-- name: UpdatePaymentStatus :exec
UPDATE payments 
SET status = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: SoftDeletePayment :exec
UPDATE payments SET deleted_at = NOW() WHERE id = ?;
