package routes

import (
	auth "e-commerce/internal/auth"
	"e-commerce/internal/product"
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
	productHandler := &product.Product{DB: db}

	authenticated := r.Group("/")
	authenticated.Use(auth.AuthMiddleware())
	{
		authenticated.GET("/user/me", userHandler.GetProfile)
		authenticated.PUT("/user/me", userHandler.UpdateProfile)
	}

	// Does not need to be authenticated
	r.GET("/products", productHandler.ListProducts)
	r.GET("/products/:id", productHandler.GetProduct)

	onlyAdmin := authenticated.Group("/")
	onlyAdmin.Use(auth.AdminOnly())
	{
		onlyAdmin.POST("/products", productHandler.CreateProduct)
		onlyAdmin.PUT("/products/:id", productHandler.UpdateProduct)
		onlyAdmin.DELETE("/products/:id", productHandler.DeleteProduct)
	}

	return r
}
