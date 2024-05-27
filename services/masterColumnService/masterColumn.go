package masterColumnService

import (
	"errors"
	"log"
	"time"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func CreateMasterColumn(claims string, tableId int, data models.CreateMasterColumn) (err error) {
	timeNow := time.Now()
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	newMasterCode := models.MasterColumn{
		CreatedBy:          user.Username,
		CreatedDate:        timeNow,
		FieldName:          data.FieldName,
		Status:             "L",
		Description:        data.Description,
		EnglishDescription: data.EnglishDescription,
		FieldType:          data.FieldType,
		FieldLength:        data.FieldLength,
		Sequence:           data.Sequence,
		IsMandatory:        data.IsMandatory,
		IsExport:           data.IsExport,
		OnblurScript:       data.OnblurScript,
		SqlFunction:        data.SqlFunction,
		TableId:            tableId,
	}

	if err := connection.DB.Create(&newMasterCode).Error; err != nil {
		return errors.New("failed to create Master Column")
	}

	// Return success response
	return nil
}

func ShowMasterColumn() (result []models.MasterTable) {
	var masterCode []models.MasterTable

	connection.DB.Where("status = ?", "L").Find(&masterCode)

	return masterCode
}

func ShowMasterColumnDetail(masterColumnId string) (result models.MasterColumn, err error) {
	var masterColumn models.MasterColumn

	// Mencari detail MasterTable berdasarkan id
	if err := connection.DB.Where("id = ?", masterColumnId).First(&masterColumn).Error; err != nil {
		return result, errors.New("MasterTable not found")
	}

	return masterColumn, nil
}

func ShowMasterColumnByTable(masterTableId string) (result []models.MasterColumn, err error) {
	var masterColumn []models.MasterColumn

	// Mencari detail MasterTable berdasarkan id
	if err := connection.DB.Where("table_id = ?", masterTableId).First(&masterColumn).Error; err != nil {
		return result, errors.New("MasterTable not found")
	}

	return masterColumn, nil
}

func UpdateColumnTable(claims, masterTableId string, updatedMasterTable models.MasterTable) (result models.MasterTable, err error) {
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

	if err := connection.DB.Save(&masterTable).Error; err != nil {
		return result, errors.New("failed to update Master Table")
	}

	return masterTable, nil
}

func DeleteMasterColumn(claims, masterTableId string) (result models.MasterTable, err error) {
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
