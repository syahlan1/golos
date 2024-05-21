package controllers

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
)

func CreateMasterColumn(c *fiber.Ctx) error {
	tableIdStr := c.Params("id")
	tableId, err := strconv.Atoi(tableIdStr)
	if err != nil {
		return err
	}

	var data map[string]interface{}

	timeNow := time.Now()

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	newMasterCode := models.MasterColumn{
		CreatedBy:          user.Username,
		CreatedDate:        timeNow,
		FieldName:          getStringValue(data, "field_name"),
		Status:             "L",
		Description:        getStringValue(data, "description"),
		EnglishDescription: getStringValue(data, "english_description"),
		FieldType:          getStringValue(data, "field_type"),
		FieldLength:        getIntValue(data, "field_length"),
		Sequence:           getIntValue(data, "sequence"),
		IsMandatory:        getBoolValue(data, "is_mandatory"),
		IsExport:           getBoolValue(data, "is_export"),
		OnblurScript:       getStringValue(data, "onblur_script"),
		SqlFunction:        getStringValue(data, "sql_function"),
		TableId:            tableId,
	}

	if err := connection.DB.Create(&newMasterCode).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create Master Column"})
	}

	// Return success response
	return c.JSON(fiber.Map{"message": "Master Column Created!"})
}

// Helper function to get bool value from map
func getBoolValue(data map[string]interface{}, key string) bool {
	value, ok := data[key].(bool)
	return ok && value
}

func ShowMasterColumn(c *fiber.Ctx) error {
	var masterCode []models.MasterTable

	connection.DB.Where("status = ?", "L").Find(&masterCode)

	return c.JSON(masterCode)
}

func ShowMasterColumnDetail(c *fiber.Ctx) error {
	masterColumnId := c.Params("id")
	var masterColumn models.MasterColumn

	// Mencari detail MasterTable berdasarkan id
	if err := connection.DB.Where("id = ?", masterColumnId).First(&masterColumn).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "MasterTable not found"})
	}

	return c.JSON(masterColumn)
}

func ShowMasterColumnByTable(c *fiber.Ctx) error {
	masterTableId := c.Params("id")
	var masterColumn []models.MasterColumn

	// Mencari detail MasterTable berdasarkan id
	if err := connection.DB.Where("table_id = ?", masterTableId).First(&masterColumn).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "MasterTable not found"})
	}

	return c.JSON(masterColumn)
}

func UpdateColumnTable(c *fiber.Ctx) error {
	masterTableId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	//

	var masterTable models.MasterTable
	if err := connection.DB.First(&masterTable, masterTableId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Data Not Found",
		})
	}

	var updatedMasterTable models.MasterTable
	if err := c.BodyParser(&updatedMasterTable); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid Master Table Data",
		})
	}

	masterTable.UpdatedBy = user.Username
	masterTable.UpdatedDate = time.Now()
	masterTable.ModuleName = updatedMasterTable.ModuleName
	masterTable.Description = updatedMasterTable.Description
	masterTable.EnglishDescription = updatedMasterTable.EnglishDescription
	masterTable.OrderField = updatedMasterTable.OrderField
	masterTable.FormType = updatedMasterTable.FormType

	if err := connection.DB.Save(&masterTable).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to update Master Table",
		})
	}

	return c.JSON(fiber.Map{
		"status": "Updated!",
		"data":   masterTable,
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

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	//

	var masterTable models.MasterTable
	if err := connection.DB.First(&masterTable, masterTableId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Data Not Found",
		})
	}

	masterTable.UpdatedBy = user.Username
	masterTable.UpdatedDate = time.Now()
	masterTable.Status = "D"

	if err := connection.DB.Save(&masterTable).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to delete Master Table",
		})
	}

	return c.JSON(fiber.Map{
		"status": "Deleted!",
		"data":   masterTable,
	})
}
