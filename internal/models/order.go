package models

import (
	"time"
)

type Order struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserId uint `json:"user_id`
	TotalAmount uint `json:"total_amount`
	Status string `json:"status" gorm:"default:pending"`
	PaymentId string `json:"payment_id"`
	CreatedAt time.Time `json:"created_at"`
}