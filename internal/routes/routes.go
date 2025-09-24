package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/ayo-69/e-commerce-backend/internal/auth"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	
	authHandler := &auth.AuthHandler{DB:db}
	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)
	
	return r
}