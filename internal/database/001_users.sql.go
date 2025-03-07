// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: 001_users.sql

package database

import (
	"context"
	"database/sql"
)

const checkUserExists = `-- name: CheckUserExists :one
SELECT COUNT(*) FROM ` + "`" + `users` + "`" + ` where email = ?
`

func (q *Queries) CheckUserExists(ctx context.Context, email string) (int64, error) {
	row := q.db.QueryRowContext(ctx, checkUserExists, email)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUser = `-- name: CreateUser :execresult
INSERT INTO ` + "`" + `users` + "`" + ` (id, full_name, email, password, phone, role, salt)
VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreateUserParams struct {
	ID       string
	FullName string
	Email    string
	Password string
	Phone    sql.NullString
	Role     NullUsersRole
	Salt     string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.FullName,
		arg.Email,
		arg.Password,
		arg.Phone,
		arg.Role,
		arg.Salt,
	)
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT 
  ` + "`" + `users` + "`" + `.id AS user_id,
  ` + "`" + `users` + "`" + `.full_name,
  ` + "`" + `users` + "`" + `.email,
  ` + "`" + `users` + "`" + `.password,
  ` + "`" + `users` + "`" + `.phone,
  ` + "`" + `users` + "`" + `.role,
  ` + "`" + `users` + "`" + `.created_at,
  ` + "`" + `users` + "`" + `.salt
FROM ` + "`" + `users` + "`" + `
WHERE ` + "`" + `users` + "`" + `.email = ? AND ` + "`" + `users` + "`" + `.deleted_at IS NULL
`

type GetUserByEmailRow struct {
	UserID    string
	FullName  string
	Email     string
	Password  string
	Phone     sql.NullString
	Role      NullUsersRole
	CreatedAt sql.NullTime
	Salt      string
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.UserID,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.Phone,
		&i.Role,
		&i.CreatedAt,
		&i.Salt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT 
  ` + "`" + `users` + "`" + `.id AS user_id,
  ` + "`" + `users` + "`" + `.full_name,
  ` + "`" + `users` + "`" + `.email,
  ` + "`" + `users` + "`" + `.password,
  ` + "`" + `users` + "`" + `.phone,
  ` + "`" + `users` + "`" + `.role,
  ` + "`" + `users` + "`" + `.created_at,
  ` + "`" + `users` + "`" + `.salt
FROM ` + "`" + `users` + "`" + `
WHERE ` + "`" + `users` + "`" + `.id = ? AND ` + "`" + `users` + "`" + `.deleted_at IS NULL
`

type GetUserByIDRow struct {
	UserID    string
	FullName  string
	Email     string
	Password  string
	Phone     sql.NullString
	Role      NullUsersRole
	CreatedAt sql.NullTime
	Salt      string
}

func (q *Queries) GetUserByID(ctx context.Context, id string) (GetUserByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i GetUserByIDRow
	err := row.Scan(
		&i.UserID,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.Phone,
		&i.Role,
		&i.CreatedAt,
		&i.Salt,
	)
	return i, err
}

const softDeleteUser = `-- name: SoftDeleteUser :exec
UPDATE ` + "`" + `users` + "`" + ` SET deleted_at = NOW() WHERE id = ?
`

func (q *Queries) SoftDeleteUser(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, softDeleteUser, id)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE ` + "`" + `users` + "`" + ` 
SET full_name = ?, email = ?, password = ?, phone = ?, role = ?
WHERE id = ? AND deleted_at IS NULL
`

type UpdateUserParams struct {
	FullName string
	Email    string
	Password string
	Phone    sql.NullString
	Role     NullUsersRole
	ID       string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.FullName,
		arg.Email,
		arg.Password,
		arg.Phone,
		arg.Role,
		arg.ID,
	)
	return err
}
