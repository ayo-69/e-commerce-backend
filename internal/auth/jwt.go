package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey []byte

func SetJwtKey(key string) {
	jwtKey = []byte(key)
}

func GenerateJWT(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
