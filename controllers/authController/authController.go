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
	var data models.Login

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

func UpdateUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	var data models.Users

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	result, err := authService.UpdateUser(userID, data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Updated!",
		Data:    result,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	err := authService.DeleteUser(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Deleted!",
	})
}

func ChangePassword(c *fiber.Ctx) error {
	userID := c.Params("id")

	var data models.Register

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	result, err := authService.ChangePassword(userID, data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Change Password Success!",
		Data:    result,
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

func ShowAllUser(c *fiber.Ctx) error {
	result := authService.ShowAllUser()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}
func UserDetailShow(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := authService.UserDetailShow(id)
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

func LogoutFromAdmin(c *fiber.Ctx) error {
	userId := c.Params("id")

	result, err := authService.LogoutFromAdmin(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Logged out",
		Data:    result,
	})
}

func UserPermission(c *fiber.Ctx) error {
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return errors.New("status unauthorized")
	}

	result, err := authService.UserPermission(claims)
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
