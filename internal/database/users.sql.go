// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const countUsers = `-- name: CountUsers :one
SELECT COUNT(*) FROM users
`

func (q *Queries) CountUsers(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countUsers)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    id, first_name, last_name, email, age, phone, role, password_hash, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    NOW(),
    NOW()
)
RETURNING id, first_name, last_name, age, role, email, verified, phone, password_hash, created_at, updated_at
`

type CreateUserParams struct {
	FirstName    string
	LastName     string
	Email        string
	Age          int32
	Phone        string
	Role         string
	PasswordHash string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Age,
		arg.Phone,
		arg.Role,
		arg.PasswordHash,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Age,
		&i.Role,
		&i.Email,
		&i.Verified,
		&i.Phone,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getPass = `-- name: GetPass :one
SELECT password_hash, id FROM users WHERE email = $1
`

type GetPassRow struct {
	PasswordHash string
	ID           uuid.UUID
}

func (q *Queries) GetPass(ctx context.Context, email string) (GetPassRow, error) {
	row := q.db.QueryRowContext(ctx, getPass, email)
	var i GetPassRow
	err := row.Scan(&i.PasswordHash, &i.ID)
	return i, err
}

const getRole = `-- name: GetRole :one
SELECT role FROM users WHERE id = $1
`

func (q *Queries) GetRole(ctx context.Context, id uuid.UUID) (string, error) {
	row := q.db.QueryRowContext(ctx, getRole, id)
	var role string
	err := row.Scan(&role)
	return role, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, first_name, last_name, age, role, email, verified, phone, password_hash, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Age,
		&i.Role,
		&i.Email,
		&i.Verified,
		&i.Phone,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, first_name, last_name, age, role, email, verified, phone, password_hash, created_at, updated_at FROM users WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Age,
		&i.Role,
		&i.Email,
		&i.Verified,
		&i.Phone,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByPhone = `-- name: GetUserByPhone :one
SELECT id, first_name, last_name, age, role, email, verified, phone, password_hash, created_at, updated_at FROM users WHERE phone = $1
`

func (q *Queries) GetUserByPhone(ctx context.Context, phone string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByPhone, phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Age,
		&i.Role,
		&i.Email,
		&i.Verified,
		&i.Phone,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const resetUsers = `-- name: ResetUsers :exec
DELETE FROM users
`

func (q *Queries) ResetUsers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, resetUsers)
	return err
}

const setVerified = `-- name: SetVerified :exec
UPDATE users
SET
    verified = TRUE,
    updated_at = NOW()
WHERE id = $1
`

func (q *Queries) SetVerified(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, setVerified, id)
	return err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
    first_name = $1,
    last_name = $2,
    email = $3,
    age = $4,
    phone = $5,
    updated_at = NOW()
WHERE id = $6
RETURNING id, first_name, last_name, age, role, email, verified, phone, password_hash, created_at, updated_at
`

type UpdateUserParams struct {
	FirstName string
	LastName  string
	Email     string
	Age       int32
	Phone     string
	ID        uuid.UUID
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Age,
		arg.Phone,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Age,
		&i.Role,
		&i.Email,
		&i.Verified,
		&i.Phone,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
