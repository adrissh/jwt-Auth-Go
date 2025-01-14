package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint   `gorm:"primaryKey; auto increment"`
	Uuid          string `gorm:"not null"`
	Username      string `gorm:"unique;not null"`
	Email         string `gorm:"unique;not null"`
	Password_hash string `gorm:"not nul"`
	Role          string `gorm:"not nul"`
	Last_login    *time.Time
	gorm.Model
}
