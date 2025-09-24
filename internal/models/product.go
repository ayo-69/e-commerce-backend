package models

import "time"

type Product struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Description string `json:"description"`
	Price uint `json:"price"`
	Stock uint `json:"stock"`
	Category string `json"category"`
	CreatedAt time.Time `json:"created_at"`
}