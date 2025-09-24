package product

import (
	"net/http"
	"github.com/ayo-69/e-commerce-backend/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler struct {
	DB *gorm.DB
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	// var products []models.Product
	
	// Fetch all products 
	// Return result
}

func (h *ProductHandler) GetProductById(c *gin.Context) {
	var product models.Product
	
	// Fetch product by id
	// productId := c.Params("id")

	// if err := h.DB.Where("id = ?", id).First(&product).Error; err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"error": "Product not found",
	// 	})
	// 	return
	// }
	
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
		Description string `json:"email"`
		Price uint `json:"password"`
		Stock uint `json:"stock"` 
		Category string `json"category"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}
	
	product := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
	}
	if err := h.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error creating product",
		})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created",
		"product": product,
	})
}


func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
		Description string `json:"email"`
		Price string `json:"password"`
		Stock uint `json:"stock"` 
		Category string `json"category"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}
	
	// Todo: Update and save product 
	
	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created",
		// "product": product,
	})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	
	// Todo: Find and Delete product 
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted", 
	})
}
