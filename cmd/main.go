package main

import (
	"log"
	"github.com/ayo-69/e-commerce-backend/internal/routes"
	"github.com/ayo-69/e-commerce-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main () {
	DSN := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Cart{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.OrderItem{})
	
	
	r := routes.SetupRouter(db)
	r.Run(":8080")
}