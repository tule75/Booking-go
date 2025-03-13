// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: 006_reviews.sql

package database

import (
	"context"
	"database/sql"
)

const createReview = `-- name: CreateReview :execresult
INSERT INTO reviews (id, user_id, property_id, rating, comment)
VALUES (?, ?, ?, ?, ?)
`

type CreateReviewParams struct {
	ID         string         `json:"id"`
	UserID     string         `json:"user_id"`
	PropertyID string         `json:"property_id"`
	Rating     sql.NullInt32  `json:"rating"`
	Comment    sql.NullString `json:"comment"`
}

func (q *Queries) CreateReview(ctx context.Context, arg CreateReviewParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createReview,
		arg.ID,
		arg.UserID,
		arg.PropertyID,
		arg.Rating,
		arg.Comment,
	)
}

const getReviewByID = `-- name: GetReviewByID :one
SELECT 
  reviews.id AS review_id,
  reviews.user_id,
  reviews.property_id,
  reviews.rating,
  reviews.comment,
  reviews.created_at
FROM reviews
WHERE reviews.id = ? AND reviews.deleted_at IS NULL
`

type GetReviewByIDRow struct {
	ReviewID   string         `json:"review_id"`
	UserID     string         `json:"user_id"`
	PropertyID string         `json:"property_id"`
	Rating     sql.NullInt32  `json:"rating"`
	Comment    sql.NullString `json:"comment"`
	CreatedAt  sql.NullTime   `json:"created_at"`
}

func (q *Queries) GetReviewByID(ctx context.Context, id string) (GetReviewByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getReviewByID, id)
	var i GetReviewByIDRow
	err := row.Scan(
		&i.ReviewID,
		&i.UserID,
		&i.PropertyID,
		&i.Rating,
		&i.Comment,
		&i.CreatedAt,
	)
	return i, err
}

const listReviewsByProperty = `-- name: ListReviewsByProperty :many
SELECT 
  r.id AS review_id,
  r.user_id,
  r.property_id,
  r.rating,
  r.comment,
  r.created_at,
  (SELECT AVG(r2.rating) FROM reviews r2 WHERE r2.property_id = r.property_id AND r2.deleted_at IS NULL) AS avg_rating
FROM reviews r
WHERE r.property_id = ? 
  AND r.deleted_at IS NULL
ORDER BY r.created_at DESC
LIMIT ? OFFSET ?
`

type ListReviewsByPropertyParams struct {
	PropertyID string `json:"property_id"`
	Limit      int32  `json:"limit"`
	Offset     int32  `json:"offset"`
}

type ListReviewsByPropertyRow struct {
	ReviewID   string         `json:"review_id"`
	UserID     string         `json:"user_id"`
	PropertyID string         `json:"property_id"`
	Rating     sql.NullInt32  `json:"rating"`
	Comment    sql.NullString `json:"comment"`
	CreatedAt  sql.NullTime   `json:"created_at"`
	AvgRating  interface{}    `json:"avg_rating"`
}

func (q *Queries) ListReviewsByProperty(ctx context.Context, arg ListReviewsByPropertyParams) ([]ListReviewsByPropertyRow, error) {
	rows, err := q.db.QueryContext(ctx, listReviewsByProperty, arg.PropertyID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListReviewsByPropertyRow
	for rows.Next() {
		var i ListReviewsByPropertyRow
		if err := rows.Scan(
			&i.ReviewID,
			&i.UserID,
			&i.PropertyID,
			&i.Rating,
			&i.Comment,
			&i.CreatedAt,
			&i.AvgRating,
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

const softDeleteReview = `-- name: SoftDeleteReview :exec
UPDATE reviews SET deleted_at = NOW() WHERE id = ?
`

func (q *Queries) SoftDeleteReview(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, softDeleteReview, id)
	return err
}

const updateReview = `-- name: UpdateReview :exec
UPDATE reviews 
SET rating = COALESCE(?, rating), comment = COALESCE(?, comment)
WHERE id = ? AND deleted_at IS NULL
`

type UpdateReviewParams struct {
	Rating  sql.NullInt32  `json:"rating"`
	Comment sql.NullString `json:"comment"`
	ID      string         `json:"id"`
}

func (q *Queries) UpdateReview(ctx context.Context, arg UpdateReviewParams) error {
	_, err := q.db.ExecContext(ctx, updateReview, arg.Rating, arg.Comment, arg.ID)
	return err
}
