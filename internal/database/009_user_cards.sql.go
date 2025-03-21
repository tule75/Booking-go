// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: 009_user_cards.sql

package database

import (
	"context"
	"database/sql"
)

const createUserCard = `-- name: CreateUserCard :execresult
INSERT INTO user_cards (id, user_id, stripe_customer_id, stripe_card_id, last4, brand, exp_month, exp_year)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateUserCardParams struct {
	ID               string `json:"id"`
	UserID           string `json:"user_id"`
	StripeCustomerID string `json:"stripe_customer_id"`
	StripeCardID     string `json:"stripe_card_id"`
	Last4            string `json:"last4"`
	Brand            string `json:"brand"`
	ExpMonth         int32  `json:"exp_month"`
	ExpYear          int32  `json:"exp_year"`
}

func (q *Queries) CreateUserCard(ctx context.Context, arg CreateUserCardParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUserCard,
		arg.ID,
		arg.UserID,
		arg.StripeCustomerID,
		arg.StripeCardID,
		arg.Last4,
		arg.Brand,
		arg.ExpMonth,
		arg.ExpYear,
	)
}

const getUserCards = `-- name: GetUserCards :many
SELECT 
  user_cards.id AS card_id,
  user_cards.user_id,
  user_cards.stripe_customer_id,
  user_cards.stripe_card_id,
  user_cards.last4,
  user_cards.brand,
  user_cards.exp_month,
  user_cards.exp_year,
  user_cards.created_at
FROM user_cards
WHERE user_cards.user_id = ? AND user_cards.deleted_at IS NULL
`

type GetUserCardsRow struct {
	CardID           string       `json:"card_id"`
	UserID           string       `json:"user_id"`
	StripeCustomerID string       `json:"stripe_customer_id"`
	StripeCardID     string       `json:"stripe_card_id"`
	Last4            string       `json:"last4"`
	Brand            string       `json:"brand"`
	ExpMonth         int32        `json:"exp_month"`
	ExpYear          int32        `json:"exp_year"`
	CreatedAt        sql.NullTime `json:"created_at"`
}

func (q *Queries) GetUserCards(ctx context.Context, userID string) ([]GetUserCardsRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserCards, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserCardsRow
	for rows.Next() {
		var i GetUserCardsRow
		if err := rows.Scan(
			&i.CardID,
			&i.UserID,
			&i.StripeCustomerID,
			&i.StripeCardID,
			&i.Last4,
			&i.Brand,
			&i.ExpMonth,
			&i.ExpYear,
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

const softDeleteUserCard = `-- name: SoftDeleteUserCard :exec
UPDATE user_cards SET deleted_at = NOW() WHERE id = ? AND user_id = ?
`

type SoftDeleteUserCardParams struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
}

func (q *Queries) SoftDeleteUserCard(ctx context.Context, arg SoftDeleteUserCardParams) error {
	_, err := q.db.ExecContext(ctx, softDeleteUserCard, arg.ID, arg.UserID)
	return err
}

const updateUserCard = `-- name: UpdateUserCard :exec
UPDATE user_cards
SET stripe_customer_id = COALESCE(?, stripe_customer_id), stripe_card_id = COALESCE(?, stripe_card_id), 
  last4 = COALESCE(?, last4), brand = COALESCE(?, brand), 
  exp_month = COALESCE(?, exp_month), exp_year = COALESCE(?, exp_year)
WHERE id = ? AND user_id = ? AND deleted_at IS NULL
`

type UpdateUserCardParams struct {
	StripeCustomerID string `json:"stripe_customer_id"`
	StripeCardID     string `json:"stripe_card_id"`
	Last4            string `json:"last4"`
	Brand            string `json:"brand"`
	ExpMonth         int32  `json:"exp_month"`
	ExpYear          int32  `json:"exp_year"`
	ID               string `json:"id"`
	UserID           string `json:"user_id"`
}

func (q *Queries) UpdateUserCard(ctx context.Context, arg UpdateUserCardParams) error {
	_, err := q.db.ExecContext(ctx, updateUserCard,
		arg.StripeCustomerID,
		arg.StripeCardID,
		arg.Last4,
		arg.Brand,
		arg.ExpMonth,
		arg.ExpYear,
		arg.ID,
		arg.UserID,
	)
	return err
}
