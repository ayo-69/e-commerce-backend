package models

type OrderItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OrderId uint `json:"order_id"`
	ProductId uint `json:"product_id"`
	Quantity uint `json:"quantity"`
	Price uint `json:"price"`
}