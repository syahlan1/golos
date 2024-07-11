package masterModuleController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/masterModuleService"
	"github.com/syahlan1/golos/utils"
)

func CreateMasterModule(c *fiber.Ctx) error {
	var data models.MasterModule

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
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err = masterModuleService.CreateMasterModule(claims, data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Return success response
	return c.JSON(models.Response{
		Code:    fiber.StatusCreated,
		Message: "Master Code Created!",
	})
}

func ShowMasterModule(c *fiber.Ctx) error {
	result := masterModuleService.ShowAllMasterModule(false)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowActiveMasterModule(c *fiber.Ctx) error {
	result := masterModuleService.ShowAllMasterModule(true)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowMasterModuleDetail(c *fiber.Ctx) error {
	masterModuleId := c.Params("id")
	result, err := masterModuleService.ShowMasterModuleDetail(masterModuleId)
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

func UpdateMasterModule(c *fiber.Ctx) error {
	masterModuleId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var data models.MasterModule
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid Master Module Data",
		})
	}

	result, err := masterModuleService.UpdateMasterModule(claims, masterModuleId, data)
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

func DeleteMasterModule(c *fiber.Ctx) error {
	masterModuleId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err = masterModuleService.DeleteMasterModule(claims, masterModuleId)
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
