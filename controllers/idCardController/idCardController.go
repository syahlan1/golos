package idCardController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/idCardService"
)

func ShowAddressType(c *fiber.Ctx) error {
	result := idCardService.ShowAddressType()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}
