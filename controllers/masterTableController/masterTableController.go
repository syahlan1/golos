package masterTableController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/masterTableService"
	"github.com/syahlan1/golos/utils"
)

func CreateMasterTable(c *fiber.Ctx) error {
	var data models.CreateMasterTable

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

	err = masterTableService.CreateMasterTable(claims, data)
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

func ShowMasterTable(c *fiber.Ctx) error {
	result := masterTableService.ShowMasterTable()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowMasterTableDetail(c *fiber.Ctx) error {
	masterTableId := c.Params("id")
	var masterTable models.MasterTable

	// Mencari detail MasterTable berdasarkan id
	if err := connection.DB.Where("id = ?", masterTableId).First(&masterTable).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Code:    fiber.StatusNotFound,
			Message: "MasterTable not found"})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    masterTable,
	})
}

func UpdateMasterTable(c *fiber.Ctx) error {
	masterTableId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var data models.MasterTable
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid Master Table Data",
		})
	}

	result, err := masterTableService.UpdateMasterTable(claims, masterTableId, data)
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

func DeleteMasterTable(c *fiber.Ctx) error {
	masterTableId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	result, err := masterTableService.DeleteMasterTable(claims, masterTableId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Deleted!",
		Data:    result,
	})
}

func GenerateTable(c *fiber.Ctx) error {
	// Ambil ID tabel dari parameter rute
	tableID := c.Params("id")

	err := masterTableService.GenerateTable(tableID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Table generated successfully!",
	})
}