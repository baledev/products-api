package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string     `json:"name"`
	Alias       string     `gorm:"not null" json:"alias"`
	Description string     `gorm:"not null" json:"description"`
	CategoryId  string     `gorm:"not null" json:"category_id"`
	CreatedAt   *time.Time `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"not null;default:now()" json:"updated_at"`
}

type ProductData struct {
	Name        string     `json:"name"`
	Alias       string     `json:"alias"`
	Description string     `json:"description"`
	CategoryId  string     `json:"category_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
