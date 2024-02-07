package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `json:"id"`
	Username string `gorm:"uniqueIndex; not null" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Name     string `json:"name"`
}

type LoginInput struct {
	Identity string `json:"identity" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterInput struct {
	Name                 string `json:"name" validate:"required"`
	Username             string `json:"username" validate:"required"`
	Email                string `json:"email" validate:"required"`
	Password             string `gorm:"not null" json:"password" validate:"required"`
	PasswordConfirmation string `gorm:"not null" json:"passwordConfirmation" validate:"required"`
}

type UserData struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `gorm:"not null" json:"password"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func TransformUserRecord(user *User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}
}
