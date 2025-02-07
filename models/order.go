package models

import "time"

type Order struct {
	ID        uint `json:i"id" gorm:"primaryKey"`
	createdAt time.Time
	firstname string `json:"first_name"`
	lastname  string `json:"last_name"`
}
