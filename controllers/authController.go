package controllers

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Check apakah username sudah ada
	var existingUser models.Users
	if err := connection.DB.Where("username = ?", data["username"]).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Username already exists",
		})
	}

	// Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create new user entry
	newUser := models.Users{
		Username: data["username"],
		Email:    data["email"],
		Password: hashedPassword,
		IsLogin:  0, // IsLogin default nya 0
		RoleId:   0,
	}
	if err := connection.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	// Return success response
	return c.JSON(fiber.Map{"message": "User registered successfully"})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.Users

	connection.DB.Where("username = ?", data["username"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if user.IsLogin == 1 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "User is already active",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	if err := connection.DB.Model(&user).Update("is_login", 1).Error; err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Failed to update user status"})
	}

	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could't create token",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Second * time.Duration(tokenTTL)),
		HTTPOnly: true,
		Secure:   true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Login Succcessfully",
	})
}

// show username
func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	// Memverifikasi token dan mendapatkan klaim
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "status unauthorized",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	// Mendapatkan data pengguna (user) dari database
	var user models.Users
	if err := connection.DB.Where("id = ?", claims.Issuer).Preload("Role").First(&user).Error; err != nil {
		return err
	}

	// Preload izin-izin (permissions) dari peran (role) pengguna
	if err := connection.DB.Model(&user.Role).Association("Permissions").Find(&user.Role.Permissions); err != nil {
		return err
	}

	return c.JSON(user)
}

// logout
func Logout(c *fiber.Ctx) error {
	userId := getUserIdFromToken(c)

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	var user models.Users

	connection.DB.Model(&user).Where("id = ?", userId).Update("is_login", 0)

	return c.JSON(fiber.Map{
		"message": "Logged out Successfully!",
	})
}

// Fungsi untuk mendapatkan ID pengguna dari token JWT
func getUserIdFromToken(c *fiber.Ctx) uint {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		// Jika terjadi kesalahan saat parsing token, kembalikan ID 0
		return 0
	}

	claims := token.Claims.(*jwt.StandardClaims)

	// Mengonversi issuer token menjadi tipe data uint
	userId, _ := strconv.ParseUint(claims.Issuer, 10, 64)

	return uint(userId)
}

func CreateRole(c *fiber.Ctx) error {
	// Parse request body
	var req struct {
		Name        string   `json:"name"`
		Permissions []string `json:"permissions"`
	}
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	// Cek apakah role dengan nama yang sama sudah ada
	var existingRole models.Roles
	if err := connection.DB.Where("name = ?", req.Name).First(&existingRole).Error; err == nil {
		// Role sudah ada, kirim respons konflik
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Role already exists"})
	}

	// Buat role baru
	newRole := models.Roles{Name: req.Name}

	// Simpan role baru ke dalam database
	if err := connection.DB.Create(&newRole).Error; err != nil {
		return err
	}

	// Dapatkan ID permissions berdasarkan nama permissions yang diberikan
	var permissions []models.Permission
	if err := connection.DB.Where("name IN ?", req.Permissions).Find(&permissions).Error; err != nil {
		return err
	}

	// Buat entri RolePermission untuk setiap permission yang terkait dengan role baru
	for _, permission := range permissions {
		// Periksa apakah RolePermission sudah ada
		var existingRolePermission models.RolePermission
		if err := connection.DB.Where("roles_id = ? AND permission_id = ?", newRole.Id, permission.Id).First(&existingRolePermission).Error; err != nil {
			// RolePermission belum ada, tambahkan entri baru
			rolePermission := models.RolePermission{RolesId: newRole.Id, PermissionId: permission.Id}
			if err := connection.DB.Create(&rolePermission).Error; err != nil {
				return err
			}
		}
	}

	return c.JSON(fiber.Map{"message": "Role created successfully"})
}
