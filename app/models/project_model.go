package models

import (
	"time"

	"github.com/google/uuid"
)

// Project struct to describe project object.
type Project struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UserID    uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	Title     string    `db:"title" json:"title" validate:"required,lte=255"`
}
