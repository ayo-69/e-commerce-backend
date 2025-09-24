package auth

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) GetProfile(c *gin.Context) {

}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	
}