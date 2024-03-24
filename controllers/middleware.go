package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func Authorize(permissionName string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Mendapatkan token dari cookie
		cookie := c.Cookies("jwt")

		// Memverifikasi token dan mendapatkan klaim
		token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{"message": "Unauthorized"})
		}

		claims := token.Claims.(*jwt.StandardClaims)

		// Mendapatkan data pengguna (user) dari database
		var user models.Users
		if err := connection.DB.Where("id = ?", claims.Issuer).Preload("Role").First(&user).Error; err != nil {
			return err
		}

		// Preload izin-izin (permissions) dari peran (role) pengguna
		connection.DB.Model(&user.Role).Association("Permissions").Find(&user.Role.Permissions)

		// Memeriksa apakah pengguna memiliki izin yang diperlukan
		hasPermission := false
		for _, permission := range user.Role.Permissions {
			if permission.Name == permissionName {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.Status(fiber.StatusForbidden)
			return c.JSON(fiber.Map{"message": "Access Denied!"})
		}

		// Izinkan akses ke rute
		return c.Next()
	}
}
