package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID        uint       `json:"id"`
	Name      string     `gorm:"not null" json:"name"`
	Slug      string     `gorm:"not null" json:"slug"`
	CreatedAt *time.Time `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt *time.Time `gorm:"not null;default:now()" json:"updated_at"`
}
