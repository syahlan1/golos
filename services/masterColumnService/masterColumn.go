package masterColumnService

import (
	"errors"
	"log"
	"time"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
)

func CreateMasterColumn(claims string, tableId int, data models.MasterColumn) (err error) {

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	data.CreatedBy = user.Username
	data.TableId = tableId

	if err := connection.DB.Create(&data).Error; err != nil {
		return errors.New("failed to create Master Column")
	}

	// Return success response
	return nil
}

func ShowMasterColumn() (result []models.MasterTable) {
	connection.DB.Where("status = ?", "L").Find(&result)

	return result
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
	if err := connection.DB.Where("table_id = ?", masterTableId).Find(&masterColumn).Error; err != nil {
		return result, errors.New("MasterTable not found")
	}

	return masterColumn, nil
}

func UpdateMasterColumn(claims, masterColumnId string, updatedMasterColumn models.MasterColumn) (result models.MasterColumn, err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return result, err
	}

	var masterColumn models.MasterColumn
	if err := connection.DB.First(&masterColumn, masterColumnId).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	masterColumn.UpdatedBy = user.Username
	masterColumn.UpdatedAt = time.Now()

	masterColumn.FieldName = updatedMasterColumn.FieldName
	masterColumn.Description = updatedMasterColumn.Description
	masterColumn.EnglishDescription = updatedMasterColumn.EnglishDescription
	masterColumn.FieldType = updatedMasterColumn.FieldType
	masterColumn.FieldLength = updatedMasterColumn.FieldLength
	masterColumn.CanFieldNegative = updatedMasterColumn.CanFieldNegative
	masterColumn.IsMandatory = updatedMasterColumn.IsMandatory
	masterColumn.Sequence = updatedMasterColumn.Sequence
	masterColumn.UiTypeId = updatedMasterColumn.UiTypeId
	masterColumn.UiSourceType = updatedMasterColumn.UiSourceType
	masterColumn.UiSourceQuery = updatedMasterColumn.UiSourceQuery
	masterColumn.CodeGroupId = updatedMasterColumn.CodeGroupId
	masterColumn.IsExport = updatedMasterColumn.IsExport
	masterColumn.IsNegative = updatedMasterColumn.IsNegative
	masterColumn.SqlFunction = updatedMasterColumn.SqlFunction
	masterColumn.OnblurScript = updatedMasterColumn.OnblurScript

	if err := connection.DB.Save(&masterColumn).Error; err != nil {
		return result, errors.New("failed to update Master Column")
	}

	return masterColumn, nil
}

func DeleteMasterColumn(claims, masterColumnId string) (err error) {
	// var user models.Users
	// if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
	// 	log.Println("Error retrieving user:", err)
	// 	return result, err
	// }

	// var masterColumn models.MasterColumn
	// if err := connection.DB.First(&masterColumn, masterColumnId).Error; err != nil {
	// 	return result, errors.New("Data Not Found")
	// }

	// masterColumn.UpdatedBy = user.Username
	// // masterColumn.UpdatedDate = time.Now()
	// // masterColumn.Status = "D"

	// if err := connection.DB.Save(&masterColumn).Error; err != nil {
	// 	return result, errors.New("Failed to delete Master Column")
	// }

	return connection.DB.Delete(&models.MasterColumn{}, masterColumnId).Error
}

func GetFormColumn(masterTableId string) (result models.TableForm, err error) {

	if err := connection.DB.
		Select("description").
		Model(&models.MasterTable{}).
		Where("id = ?", masterTableId).Scan(&result.FormName).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	rows, err := connection.DB.
		Select("field_name, is_mandatory, need_first_empty, ut.name as ui_type, ui_source_type, ui_source_query, code_group_id").
		Joins("JOIN ui_types ut ON ut.id = master_columns.ui_type_id").
		Model(&models.MasterColumn{}).
		Where("table_id = ?", masterTableId).Rows()

	var data models.FormList

	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		err = connection.DB.ScanRows(rows, &data)
		if err != nil {
			return result, err
		}

		if data.UiSourceType == "C" {
			err = connection.DB.
				Select("id, description as name, english_description as name_en").
				Model(&models.MasterCode{}).
				Where("code_group_id = ?", data.CodeGroupId).
				Order("sequence").
				Scan(&data.UiSource).Error
			if err != nil {
				return result, err
			}

		} else if data.UiSourceType == "Q" {
			err = connection.DB.
				Raw(data.UiSourceQuery).
				Scan(&data.UiSource).Error
			if err != nil {
				return result, err
			}
		}

		if data.NeedFirstEmpty {
			data.UiSource = utils.Prepend(data.UiSource, models.DropdownEn{Id: 0, Name: "", NameEn: ""})
		}

		result.Form = append(result.Form, data)
	}

	return
}

func CheckQuery(data models.CheckQuery) (err error) {
	return connection.DB.Raw(`?`,data.Query).Error
}

func GetUiType() (result []models.UiType) {
	connection.DB.Order("id").Find(&result)

	result = utils.Prepend(result, models.UiType{Id: 0, Name: ""})
	return result
}

func GetFieldType() (result []models.FieldType) {

	results := []models.FieldType{
		{
			Code: "",
			Name: "",
		},
		{
			Code: "A",
			Name: "Alphanumeric",
		},
		{
			Code: "B",
			Name: "Sign",
		},
		{
			Code: "D",
			Name: "Date",
		},
		{
			Code: "F",
			Name: "Float",
		},
		{
			Code: "N",
			Name: "Numeric",
		},
	}

	return results

}
