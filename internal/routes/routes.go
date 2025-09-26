package routes

import (
	auth "e-commerce/internal/auth"
	user "e-commerce/internal/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	authHandler := &auth.AuthHandler{DB: db}
	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)

	userHandler := &user.UserHandler{DB: db}

	authenticated := r.Group("/")
	authenticated.Use(auth.AuthMiddleware())
	{
		authenticated.GET("/user/me", userHandler.GetProfile)
		authenticated.PUT("/user/me", userHandler.UpdateProfile)
	}

	return r
}
