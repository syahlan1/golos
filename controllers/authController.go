package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//check existing userame
	var existingUser models.Users
	if err := connection.DB.Where("username = ?", data["username"]).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Username already exist",
		})
	}

	// Retrieve the password from the request data
	password := []byte(data["password"])

	// Generate the hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.Users{
		Username: data["username"],
		Roles:    "user",
		Status:   0,
		Password: hashedPassword,
	}

	connection.DB.Create(&user)

	return c.JSON(fiber.Map{
		"message": "registrasi sukses",
	})
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

	if user.Status == 1 {
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

	if err := connection.DB.Model(&user).Update("status", 1).Error; err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Failed to update user status"})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
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
		Expires:  time.Now().Add(time.Hour * 45),
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

	var user models.Users

	connection.DB.Where("id = ?", claims.Issuer).First(&user)

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

	connection.DB.Model(&user).Where("id = ?", userId).Update("status", 0)

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
