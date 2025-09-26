package product

import (
	"e-commerce/internal/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func (p *Product) ListProducts(c *gin.Context) {
	var products []models.Product
	if err := p.DB.Find(&products).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve products"})
		return
	}
	c.JSON(200, products)
}

func (p *Product) GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := p.DB.First(&product, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(200, product)
}

func (p *Product) CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		Category:    input.Category,
		CreatedAt:   time.Now(),
		// Add other fields as needed
	}

	if err := p.DB.Create(&product).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create product"})
		return
	}
	c.JSON(201, product)
}

func (p *Product) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var product models.Product
	if err := p.DB.First(&product, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Stock = input.Stock
	// Add other fields as needed

	if err := p.DB.Save(&product).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update product"})
		return
	}
	c.JSON(200, product)
}

func (p *Product) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := p.DB.First(&product, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	if err := p.DB.Delete(&product).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete product"})
		return
	}
	c.JSON(200, gin.H{"message": "Product deleted successfully"})
}
