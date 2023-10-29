package models

import (
	"time"
)

type Login struct {
  Login_id  int `gorm:"primaryKey"`
  Login_user_id int
	UserRefer Users `gorm:"foreignKey:Login_user_id;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
  Login_token string
  Login_created_at time.Time
}