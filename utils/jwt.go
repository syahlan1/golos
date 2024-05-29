package utils

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func ExtractJWT(c *fiber.Ctx) (string, error) {
	cookie := c.Cookies("jwt")
	if cookie == "" {
		log.Println("JWT cookie not found")
		return "", errors.New("JWT cookie not found")
	}

	log.Println("JWT cookie:", cookie)
	SecretKey := []byte(os.Getenv("JWT_PRIVATE_KEY"))

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		log.Println("Error parsing JWT:", err)
		return "", err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims.Issuer, nil
	} else {
		log.Println("Invalid JWT token")
		return "", errors.New("invalid JWT token")
	}
}

func GenerateJWT(id uint, tokenTTL int) (string, error) {
	SecretKey := []byte(os.Getenv("JWT_PRIVATE_KEY"))
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(id)),
		ExpiresAt: time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}

	return token, nil
}
