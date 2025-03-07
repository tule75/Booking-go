-- name: CreateUserCard :execresult
INSERT INTO user_cards (id, user_id, stripe_customer_id, stripe_card_id, last4, brand, exp_month, exp_year)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetUserCards :many
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
WHERE user_cards.user_id = ? AND user_cards.deleted_at IS NULL;

-- name: UpdateUserCard :exec
UPDATE user_cards
SET stripe_customer_id = ?, stripe_card_id = ?, last4 = ?, brand = ?, exp_month = ?, exp_year = ?
WHERE id = ? AND user_id = ? AND deleted_at IS NULL;

-- name: SoftDeleteUserCard :exec
UPDATE user_cards SET deleted_at = NOW() WHERE id = ? AND user_id = ?;
