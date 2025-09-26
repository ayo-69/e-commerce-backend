package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"e-commerce/internal/auth"
	"e-commerce/internal/models"
	"e-commerce/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestRouter() (*gin.Engine, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.OrderItem{}, &models.Cart{})

	auth.SetJwtKey("testsecret")

	r := routes.SetupRouter(db)
	return r, db
}

func TestUserRoutes(t *testing.T) {
	r, db := setupTestRouter()

	// Clean up the database after the test
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Register a new user
	registerBody := `{"name": "testuser2", "password": "password", "email": "test2@example.com"}`
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer([]byte(registerBody)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// Login to get the token
	loginBody := `{"email": "test2@example.com", "password": "password"}`
	req, _ = http.NewRequest("POST", "/auth/login", bytes.NewBuffer([]byte(loginBody)))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)

	token := response["token"]

	// Test GetProfile
	req, _ = http.NewRequest("GET", "/user/me", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var profile models.User
	json.Unmarshal(w.Body.Bytes(), &profile)

	assert.Equal(t, "testuser2", profile.Name)
}
