package authController

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/authService"
	"github.com/syahlan1/golos/utils"
)

func Register(c *fiber.Ctx) error {
	var data models.Register

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error()})
	}

	err := authService.Register(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "User registered successfully",
	})
}

func Login(c *fiber.Ctx) error {
	var data models.Register

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	token, tokenTTL, err := authService.Login(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
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
		"token":   token,
	})
}

// show username
func User(c *fiber.Ctx) error {
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return errors.New("status unauthorized")
	}

	result, err := authService.User(claims)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowRole(c *fiber.Ctx) error {

	result, err := authService.ShowRole()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowAllPermissions(c *fiber.Ctx) error {
	result, err := authService.ShowAllPermissions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowPermissions(c *fiber.Ctx) error {
	roleID := c.Params("id")
	result, err := authService.ShowPermissions(roleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

// logout
func Logout(c *fiber.Ctx) error {
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return errors.New("status unauthorized")
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	authService.Logout(claims)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Logged out Successfully!",
	})
}

// Fungsi untuk mendapatkan ID pengguna dari token JWT
// func getUserIdFromToken(c *fiber.Ctx) uint {
// 	claims, err := utils.ExtractJWT(c)
// 	if err != nil {
// 		c.Status(fiber.StatusUnauthorized)
// 		return 0
// 	}

// 	// Mengonversi issuer token menjadi tipe data uint
// 	userId, _ := strconv.ParseUint(claims, 10, 64)

// 	return uint(userId)
// }

func CreateRole(c *fiber.Ctx) error {
	// Parse request body
	var data models.CreateRole
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	err = authService.CreateRole(claims, data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Role created successfully",
	})
}

func DeleteRole(c *fiber.Ctx) error {
	roleID := c.Params("id")

	err := authService.DeleteRole(roleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Role deleted successfully",
	})
}

func UpdateRole(c *fiber.Ctx) error {
	// Parse request body
	var data models.CreateRole
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return errors.New("status unauthorized")
	}

	roleID := c.Params("id")

	err = authService.UpdateRole(claims, roleID, data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Role updated successfully",
	})
}
