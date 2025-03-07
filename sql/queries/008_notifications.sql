-- name: CreateNotification :execresult
INSERT INTO notifications (id, user_id, message, is_read)
VALUES (?, ?, ?, ?);

-- name: GetNotificationByID :one
SELECT 
  notifications.id AS notification_id,
  notifications.user_id,
  notifications.message,
  notifications.is_read,
  notifications.created_at
FROM notifications
WHERE notifications.id = ? AND notifications.deleted_at IS NULL;

-- name: ListNotificationsByUser :many
SELECT 
  notifications.id AS notification_id,
  notifications.user_id,
  notifications.message,
  notifications.is_read,
  notifications.created_at
FROM notifications
WHERE notifications.user_id = ?
  AND notifications.deleted_at IS NULL
  AND ( ? IS NULL OR notifications.is_read = ? )
ORDER BY notifications.created_at DESC
LIMIT ? OFFSET ?;

-- name: MarkNotificationAsRead :exec
UPDATE notifications
SET is_read = TRUE
WHERE id = ? AND deleted_at IS NULL;
