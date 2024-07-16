package masterParameterController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/masterParameterService"
	"github.com/syahlan1/golos/utils"
)

func CreateParameter(c *fiber.Ctx) error {
	var data models.MasterParameter

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to parse request",
		})
	}

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err = masterParameterService.CreateParameter(claims, data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Return success response
	return c.JSON(models.Response{
		Code:    fiber.StatusCreated,
		Message: "Master Parameter Created!",
	})
}

func ShowAllParameter(c *fiber.Ctx) error {

	result, err := masterParameterService.ShowAllParameter()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowParameterDetail(c *fiber.Ctx) error {
	parameterId := c.Params("id")

	result, err := masterParameterService.ShowParameterDetail(parameterId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func UpdateMasterParameter(c *fiber.Ctx) error {
	parameterId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var data models.MasterParameter
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid Data Parameter",
		})
	}

	result, err := masterParameterService.UpdateMasterParameter(claims, parameterId, data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Updated!",
		Data:    result,
	})
}

func DeleteMasterParameter(c *fiber.Ctx) error {
	parameterId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err = masterParameterService.DeleteMasterParameter(claims, parameterId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Deleted!",
	})
}
