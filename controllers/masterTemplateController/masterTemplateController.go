package masterTemplateController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/masterTemplateService"
	"github.com/syahlan1/golos/utils"
)

func ShowMasterTemplate(c *fiber.Ctx) error {
	module := c.Params("module")
	table := c.Params("table")

	result, err := masterTemplateService.ShowMasterTemplate(module, table, "", "", "", "", "")
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

func ShowMasterTemplateById(c *fiber.Ctx) error {
	module := c.Params("module")
	table := c.Params("table")
	id := c.Params("id")

	result, err := masterTemplateService.ShowMasterTemplate(module, table, "", "", "", "", id)
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

func CreateMasterTemplate(c *fiber.Ctx) error {
	module := c.Params("module")
	table := c.Params("table")

	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	claims, err := utils.TakeUsername(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err, errValidation := masterTemplateService.CreateMasterTemplate(module, table, claims, "", data)
	if errValidation != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Error Validation",
			Data:    errValidation,
		})
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response{
		Code:    fiber.StatusCreated,
		Message: "Success Create",
	})
}

func UpdateMasterTemplate(c *fiber.Ctx) error {
	module := c.Params("module")
	table := c.Params("table")
	id := c.Params("id")

	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	claims, err := utils.TakeUsername(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err, errValidation := masterTemplateService.UpdateMasterTemplate(module, table, id, claims, data)
	if errValidation != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Error Validation",
			Data:    errValidation,
		})
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success Update",
	})
}

func DeleteMasterTemplate(c *fiber.Ctx) error {
	module := c.Params("module")
	table := c.Params("table")
	id := c.Params("id")

	claims, err := utils.TakeUsername(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err = masterTemplateService.DeleteMasterTemplate(module, table, id, claims)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success Delete",
	})
}
