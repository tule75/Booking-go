-- name: CreateUser :exec
INSERT INTO `users` (id, full_name, email, password, phone, role)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetUserByID :one
SELECT 
  `users`.id AS user_id,
  `users`.full_name,
  `users`.email,
  `users`.password,
  `users`.phone,
  `users`.role,
  `users`.created_at
FROM `users`
WHERE `users`.id = ? AND `users`.deleted_at IS NULL;

-- name: GetUserByEmail :one
SELECT 
  `users`.id AS user_id,
  `users`.full_name,
  `users`.email,
  `users`.password,
  `users`.phone,
  `users`.role,
  `users`.created_at
FROM `users`
WHERE `users`.email = ? AND `users`.deleted_at IS NULL;

-- name: UpdateUser :exec
UPDATE `users` 
SET full_name = ?, email = ?, password = ?, phone = ?, role = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: SoftDeleteUser :exec
UPDATE `users` SET deleted_at = NOW() WHERE id = ?;
