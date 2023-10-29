package models

import "time"

type Users struct {
	ID        int `gorm:"primaryKey"`
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}