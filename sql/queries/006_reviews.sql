-- name: CreateReview :execresult
INSERT INTO reviews (id, user_id, property_id, rating, comment)
VALUES (?, ?, ?, ?, ?);

-- name: GetReviewByID :one
SELECT 
  reviews.id AS review_id,
  reviews.user_id,
  reviews.property_id,
  reviews.rating,
  reviews.comment,
  reviews.created_at
FROM reviews
WHERE reviews.id = ? AND reviews.deleted_at IS NULL;

-- name: ListReviewsByProperty :many
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
LIMIT ? OFFSET ?;

-- name: UpdateReview :exec
UPDATE reviews 
SET rating = COALESCE(?, rating), comment = COALESCE(?, comment)
WHERE id = ? AND deleted_at IS NULL;

-- name: SoftDeleteReview :exec
UPDATE reviews SET deleted_at = NOW() WHERE id = ?;
