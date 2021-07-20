package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Project struct to describe project object.
type Project struct {
	gorm.Model
	UUID  uuid.UUID `json:"uuid" gorm:"not null"`
	Title string    `json:"title" validate:"required,lte=255"`
}

type ProjectSerializer struct {
	Title string    `json:"title"`
	UUID  uuid.UUID `json:"uuid"`
}
