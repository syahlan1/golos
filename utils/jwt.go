package utils

import (
	"errors"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func ExtractJWT(c *fiber.Ctx) (string, error) {

	cookie := c.Cookies("jwt")
	SecretKey := []byte(os.Getenv("JWT_PRIVATE_KEY"))
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		log.Println("Error parsing JWT:", err)
		c.Status(fiber.StatusUnauthorized)
		return "", errors.New("status unauthorized")

	}
	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Issuer, nil
}
