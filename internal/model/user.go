package model

import (
	"time"
)

type User struct {
	ID        uint64 `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Password  string
	Email     string `gorm:"unique"`
	CreatedAt time.Time
}
