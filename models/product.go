package models

import "time"

type Product struct {
	ID           uint `json:i"id" gorm:"primaryKey"`
	createdAt    time.Time
	name         string `json:"product_name"`
	serialNumber string `json:"serial_name"`
}
