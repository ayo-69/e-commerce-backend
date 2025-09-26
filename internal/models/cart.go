package models

type Cart struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserId uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Quantity uint `json:"quantity"`
}