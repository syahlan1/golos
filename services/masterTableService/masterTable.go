package masterTableService

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func CreateMasterTable(claims string, data models.CreateMasterTable) (err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return err
	}

	timeNow := time.Now()

	newMasterCode := models.MasterTable{
		CreatedBy:          user.Username,
		CreatedDate:        timeNow,
		ModuleName:         "LOS",
		TableName:          data.TableName,
		Status:             "L",
		Description:        data.Description,
		EnglishDescription: data.EnglishDescription,
		OrderField:         data.OrderField,
		FormType:           data.FormType,
		PeriodType:         data.PeriodType,
		UsePeriod:          data.UsePeriod,
		UseWorkflow:        data.UseWorkflow,
		UseBranch:          data.UseBranch,
		UseDataLoader:      data.UseDataLoader,
	}

	if err := connection.DB.Create(&newMasterCode).Error; err != nil {
		return errors.New("failed to create Master Code")
	}

	// Return success response
	return nil
}

func ShowMasterTable() (result []models.MasterTable) {
	var masterCode []models.MasterTable

	connection.DB.Where("status = ?", "L").Find(&masterCode)

	return masterCode
}

func ShowMasterTableDetail(masterTableId string) (result models.MasterTable, err error) {
	var masterTable models.MasterTable

	// Mencari detail MasterTable berdasarkan id
	if err := connection.DB.Where("id = ?", masterTableId).First(&masterTable).Error; err != nil {
		return result, errors.New("masterTable not found")
	}

	return masterTable, nil
}

func UpdateMasterTable(claims, masterTableId string, updatedMasterTable models.MasterTable) (result models.MasterTable, err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return result, err
	}

	var masterTable models.MasterTable
	if err := connection.DB.First(&masterTable, masterTableId).Error; err != nil {
		return result, errors.New("data Not Found")
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
		return result, errors.New("failed to update Master Table")
	}

	return masterTable, nil
}

func DeleteMasterTable(claims, masterTableId string) (result models.MasterTable, err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return result, err
	}

	var masterTable models.MasterTable
	if err := connection.DB.First(&masterTable, masterTableId).Error; err != nil {
		return result, errors.New("Data Not Found")
	}

	masterTable.UpdatedBy = user.Username
	masterTable.UpdatedDate = time.Now()
	masterTable.Status = "D"

	if err := connection.DB.Save(&masterTable).Error; err != nil {
		return result, errors.New("Failed to delete Master Table")
	}

	return masterTable, nil
}

func GenerateTable(tableID string) (err error) {
	db := connection.DB

	// Ambil data tabel berdasarkan ID dari database
	var masterTable models.MasterTable
	if err := db.First(&masterTable, tableID).Error; err != nil {
		return errors.New("data Not Found")
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
		return err
	}

	return nil
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
