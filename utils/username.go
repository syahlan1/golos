package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func TakeUsername(c *fiber.Ctx) (string, error) {
	claims, err := ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return "", errors.New("status unauthorized")
	}

	// Mendapatkan data pengguna (user) dari database
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		return "", err
	}
	return user.Username, nil
}
