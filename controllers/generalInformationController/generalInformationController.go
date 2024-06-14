package generalInformationController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/generalInformationService"
)

func ShowCabangPencairan(c *fiber.Ctx) error {
	result := generalInformationService.ShowCabangPencairan()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowCabangAdmin(c *fiber.Ctx) error {
	result := generalInformationService.ShowCabangAdmin()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowSegment(c *fiber.Ctx) error {
	result := generalInformationService.ShowSegment()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}
func ShowProgram(c *fiber.Ctx) error {
	result := generalInformationService.ShowProgram()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func GenerateApplicationNumber(c *fiber.Ctx) error {
	cabangAdmin := c.Params("id")
	result, err := generalInformationService.GenerateApplicationNumber(cabangAdmin)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}
