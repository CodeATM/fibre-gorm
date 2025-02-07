package models

import "time"

type User struct {
	ID        uint `json:i"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}
