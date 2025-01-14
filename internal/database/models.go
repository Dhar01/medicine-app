// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type AdminPermission struct {
	AdminID      uuid.UUID
	PermissionID uuid.UUID
}

type AdminRole struct {
	UserID    uuid.UUID
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Medicine struct {
	ID           uuid.UUID
	Name         string
	Dosage       string
	Description  string
	Manufacturer string
	Price        int32
	Stock        int32
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Permission struct {
	ID   uuid.UUID
	Name string
}

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Age       int32
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserAddress struct {
	UserID        uuid.UUID
	Country       string
	City          string
	StreetAddress string
	PostalCode    sql.NullString
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
