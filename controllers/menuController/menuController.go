package menuController

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/menuService"
	"github.com/syahlan1/golos/utils"
)

func CreateMenu(c *fiber.Ctx) error {
	var data models.Menu
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// Get user ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return errors.New("status unauthorized")
	}

	err = menuService.CreateMenu(claims, data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success Create Menu",
	})
}

func ShowMenu(c *fiber.Ctx) error {
	result, err := menuService.ShowMenu()
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

