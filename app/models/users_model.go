package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct to describe User object.
type User struct {
	gorm.Model
	UUID  uuid.UUID `json:"uuid" gorm:"not null"`
	Name  string    `json:"name" gorm:"not null" validate:"required"`
	Email string    `json:"email" gorm:"not null" validate:"required"`
}
