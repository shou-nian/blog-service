package models

import "gorm.io/gorm"

type role string

const (
	Admin role = "admin"
	User  role = "user"
)

type Users struct {
	gorm.Model

	Name         string `json:"name" gorm:"not null;size:50"`
	Role         role   `json:"role" gorm:"default:user"`
	Email        string `json:"email" gorm:"uniqueIndex;not null;size:100"`
	PasswordHash string `json:"-" gorm:"not null;size:255"`
}

func (Users) TableName() string { return "users" }
