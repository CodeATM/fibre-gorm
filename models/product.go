package models

import "time"

type Product struct {
	ID           uint `json:i"id" gorm:"primaryKey"`
	createdAt    time.Time
	Name         string `json:"product_name"`
	SerialNumber string `json:"serial_name"`
}
