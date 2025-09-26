package main

import (
	"log"

	"e-commerce/internal/auth"
	"e-commerce/internal/config"
	"e-commerce/internal/models"
	"e-commerce/internal/routes"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()
	auth.SetJwtKey(cfg.JWTSecret)
	dsn := "host=localhost user=user password=mysecretpassword dbname=ecommerce port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.OrderItem{}, &models.Cart{})

	r := routes.SetupRouter(db)
	r.Run(":" + cfg.Port)
}