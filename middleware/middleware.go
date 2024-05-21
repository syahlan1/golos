package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
)

func Authorize(permissionName string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Mendapatkan token dari cookie
		claims, err := utils.ExtractJWT(c)
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{"message": "Unauthorized"})
		}

		// Mendapatkan data pengguna (user) dari database
		var user models.Users
		if err := connection.DB.Where("id = ?", claims).Preload("Role").First(&user).Error; err != nil {
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
