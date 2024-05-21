package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
)

func CreateMasterTable(c *fiber.Ctx) error {
	var data map[string]interface{}

	timeNow := time.Now()

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

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	newMasterCode := models.MasterTable{
		CreatedBy:          user.Username,
		CreatedDate:        timeNow,
		ModuleName:         "LOS",
		TableName:          data["table_name"].(string),
		Status:             "L",
		Description:        data["description"].(string),
		EnglishDescription: data["english_description"].(string),
		OrderField:         data["order_field"].(string),
		FormType:           data["form_type"].(string),
		PeriodType:         data["period_type"].(string),
		UsePeriod:          int(data["use_period"].(float64)),
		UseWorkflow:        int(data["use_workflow"].(float64)),
		UseBranch:          int(data["use_branch"].(float64)),
		UseDataLoader:      int(data["use_data_loader"].(float64)),
	}

	if err := connection.DB.Create(&newMasterCode).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to create Master Code"})
	}

	// Return success response
	return c.JSON(models.Response{
		Code:    fiber.StatusCreated,
		Message: "Master Code Created!",
	})
}

func ShowMasterTable(c *fiber.Ctx) error {
	var masterCode []models.MasterTable

	connection.DB.Where("status = ?", "L").Find(&masterCode)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    masterCode,
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

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	//

	var masterTable models.MasterTable
	if err := connection.DB.First(&masterTable, masterTableId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Code:    fiber.StatusNotFound,
			Message: "Data Not Found",
		})
	}

	var updatedMasterTable models.MasterTable
	if err := c.BodyParser(&updatedMasterTable); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid Master Table Data",
		})
	}

	masterTable.UpdatedBy = user.Username
	masterTable.UpdatedDate = time.Now()
	masterTable.ModuleName = updatedMasterTable.ModuleName
	masterTable.Description = updatedMasterTable.Description
	masterTable.EnglishDescription = updatedMasterTable.EnglishDescription
	masterTable.OrderField = updatedMasterTable.OrderField
	masterTable.FormType = updatedMasterTable.FormType
	masterTable.PeriodType = updatedMasterTable.PeriodType
	masterTable.UsePeriod = updatedMasterTable.UsePeriod
	masterTable.UseWorkflow = updatedMasterTable.UseWorkflow
	masterTable.UseBranch = updatedMasterTable.UseBranch
	masterTable.UseDataLoader = updatedMasterTable.UseDataLoader

	if err := connection.DB.Save(&masterTable).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to update Master Table",
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Updated!",
		Data:   masterTable,
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

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	//

	var masterTable models.MasterTable
	if err := connection.DB.First(&masterTable, masterTableId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Code:    fiber.StatusNotFound,
			Message: "Data Not Found",
		})
	}

	masterTable.UpdatedBy = user.Username
	masterTable.UpdatedDate = time.Now()
	masterTable.Status = "D"

	if err := connection.DB.Save(&masterTable).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to delete Master Table",
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Deleted!",
		Data:   masterTable,
	})
}

func GenerateTable(c *fiber.Ctx) error {
	// Koneksi ke database
	connection.Connect()
	db := connection.DB

	// Ambil ID tabel dari parameter rute
	tableID := c.Params("id")

	// Ambil data tabel berdasarkan ID dari database
	var masterTable models.MasterTable
	if err := db.First(&masterTable, tableID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Code:    fiber.StatusNotFound,
			Message: "Data Not Found",
		})
	}

	// Ambil semua kolom dari tabel
	var columns []models.MasterColumn
	db.Where("table_id = ?", masterTable.Id).Find(&columns)

	// Buat definisi SQL untuk membuat tabel baru
	createTableSQL := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", masterTable.TableName)
	createTableSQL += "\tID SERIAL PRIMARY KEY,\n" // Tambahkan kolom ID secara manual dengan auto-increment

	// Tambahkan kolom lain dari MasterColumn
	for _, column := range columns {
		fieldSQL := fmt.Sprintf("\t%s %s", column.FieldName, mapFieldType(column.FieldType, column.FieldLength))
		if column.IsMandatory {
			fieldSQL += " NOT NULL"
		}
		createTableSQL += fieldSQL + ",\n"
	}

	// Hapus koma terakhir dan tambahkan penutup query
	createTableSQL = createTableSQL[:len(createTableSQL)-2] + "\n);"

	// Eksekusi query untuk membuat tabel
	if err := db.Exec(createTableSQL).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Table generated successfully!",
	})
}

// mapFieldType berfungsi untuk memetakan tipe data dari MasterColumn ke tipe data SQL
func mapFieldType(fieldType string, fieldLength int) string {
	switch fieldType {
	case "A":
		if fieldLength > 0 {
			return fmt.Sprintf("VARCHAR(%d)", fieldLength)
		}
		return "VARCHAR"
	case "N":
		return "INTEGER"
	case "B":
		return "BOOLEAN"
	case "F":
		return "FLOAT"
	case "D":
		return "TIMESTAMP"
	default:
		return "VARCHAR"
	}
}
