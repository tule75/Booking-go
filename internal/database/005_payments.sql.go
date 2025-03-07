// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: 005_payments.sql

package database

import (
	"context"
	"database/sql"
)

const createPayment = `-- name: CreatePayment :execresult
INSERT INTO payments (id, booking_id, user_id, amount, payment_method, status, stripe_payment_intent_id)
VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreatePaymentParams struct {
	ID                    string
	BookingID             string
	UserID                string
	Amount                string
	PaymentMethod         NullPaymentsPaymentMethod
	Status                NullPaymentsStatus
	StripePaymentIntentID sql.NullString
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createPayment,
		arg.ID,
		arg.BookingID,
		arg.UserID,
		arg.Amount,
		arg.PaymentMethod,
		arg.Status,
		arg.StripePaymentIntentID,
	)
}

const getPaymentByID = `-- name: GetPaymentByID :one
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
WHERE payments.id = ? AND payments.deleted_at IS NULL
`

type GetPaymentByIDRow struct {
	PaymentID             string
	BookingID             string
	UserID                string
	Amount                string
	PaymentMethod         NullPaymentsPaymentMethod
	Status                NullPaymentsStatus
	StripePaymentIntentID sql.NullString
	CreatedAt             sql.NullTime
}

func (q *Queries) GetPaymentByID(ctx context.Context, id string) (GetPaymentByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getPaymentByID, id)
	var i GetPaymentByIDRow
	err := row.Scan(
		&i.PaymentID,
		&i.BookingID,
		&i.UserID,
		&i.Amount,
		&i.PaymentMethod,
		&i.Status,
		&i.StripePaymentIntentID,
		&i.CreatedAt,
	)
	return i, err
}

const listPaymentsByUser = `-- name: ListPaymentsByUser :many
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
  AND ( ? IS NULL OR payments.payment_method = ? )
ORDER BY payments.created_at DESC
LIMIT ? OFFSET ?
`

type ListPaymentsByUserParams struct {
	UserID        string
	Column2       interface{}
	PaymentMethod NullPaymentsPaymentMethod
	Limit         int32
	Offset        int32
}

type ListPaymentsByUserRow struct {
	PaymentID             string
	BookingID             string
	UserID                string
	Amount                string
	PaymentMethod         NullPaymentsPaymentMethod
	Status                NullPaymentsStatus
	StripePaymentIntentID sql.NullString
	CreatedAt             sql.NullTime
}

func (q *Queries) ListPaymentsByUser(ctx context.Context, arg ListPaymentsByUserParams) ([]ListPaymentsByUserRow, error) {
	rows, err := q.db.QueryContext(ctx, listPaymentsByUser,
		arg.UserID,
		arg.Column2,
		arg.PaymentMethod,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPaymentsByUserRow
	for rows.Next() {
		var i ListPaymentsByUserRow
		if err := rows.Scan(
			&i.PaymentID,
			&i.BookingID,
			&i.UserID,
			&i.Amount,
			&i.PaymentMethod,
			&i.Status,
			&i.StripePaymentIntentID,
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

const softDeletePayment = `-- name: SoftDeletePayment :exec
UPDATE payments SET deleted_at = NOW() WHERE id = ?
`

func (q *Queries) SoftDeletePayment(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, softDeletePayment, id)
	return err
}

const updatePaymentStatus = `-- name: UpdatePaymentStatus :exec
UPDATE payments 
SET status = ?
WHERE id = ? AND deleted_at IS NULL
`

type UpdatePaymentStatusParams struct {
	Status NullPaymentsStatus
	ID     string
}

func (q *Queries) UpdatePaymentStatus(ctx context.Context, arg UpdatePaymentStatusParams) error {
	_, err := q.db.ExecContext(ctx, updatePaymentStatus, arg.Status, arg.ID)
	return err
}
