package masterColumnController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/masterColumnService"
	"github.com/syahlan1/golos/utils"
)

func CreateMasterColumn(c *fiber.Ctx) error {
	tableIdStr := c.Params("id")
	tableId, err := strconv.Atoi(tableIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var data models.MasterColumn

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
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err = masterColumnService.CreateMasterColumn(claims, tableId, data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Return success response
	return c.JSON(models.Response{
		Code:    fiber.StatusCreated,
		Message: "Master Column Created!",
	})
}

func ShowMasterColumn(c *fiber.Ctx) error {
	result := masterColumnService.ShowMasterColumn()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowMasterColumnDetail(c *fiber.Ctx) error {
	masterColumnId := c.Params("id")

	result, err := masterColumnService.ShowMasterColumnDetail(masterColumnId)
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

func ShowMasterColumnByTable(c *fiber.Ctx) error {
	masterTableId := c.Params("id")

	result, err := masterColumnService.ShowMasterColumnByTable(masterTableId)
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

func UpdateColumnTable(c *fiber.Ctx) error {
	masterTableId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var data models.MasterColumn
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid Master Column Data",
		})
	}

	result, err := masterColumnService.UpdateMasterColumn(claims, masterTableId, data)
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

func DeleteMasterColumn(c *fiber.Ctx) error {
	masterTableId := c.Params("id")
	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err = masterColumnService.DeleteMasterColumn(claims, masterTableId)
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

func ShowFormColumn(c *fiber.Ctx) error {
	masterTableId := c.Params("id")

	result, err := masterColumnService.GetFormColumn(masterTableId)
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

func CheckQuery(c *fiber.Ctx) error {

	var data models.CheckQuery

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err := masterColumnService.CheckQuery(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Return success response
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Query Valid",
	})
}

func GetUiType(c *fiber.Ctx) error {
	result := masterColumnService.GetUiType()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func GetFieldType(c *fiber.Ctx) error {
	result := masterColumnService.GetFieldType()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}