package masterColumnService

import (
	"errors"
	"fmt"
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

	if err := utils.IsValidSqlName(data.FieldName); err != nil {
		return err
	}

	if data.Disable && data.IsMandatory {
		return errors.New("is_mandatory cannot be true when disable is true")
	}

	// prevent SQL injection
	if data.AutoFill {
		if err := utils.IsValidSQL(*data.FillQuery); err != nil {
			return err
		}
	}

	if data.UiSourceType == "Q" {
		if err := utils.IsValidSQL(*data.UiSourceQuery); err != nil {
			return err
		}
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

	if updatedMasterColumn.Disable && updatedMasterColumn.IsMandatory {
		return result, errors.New("is_mandatory cannot be true when disable is true")
	}

	// prevent SQL injection
	if updatedMasterColumn.AutoFill {
		if err := utils.IsValidSQL(*updatedMasterColumn.FillQuery); err != nil {
			return result, err
		}
	}

	if updatedMasterColumn.UiSourceType == "Q" {
		if err := utils.IsValidSQL(*updatedMasterColumn.UiSourceQuery); err != nil {
			return result, err
		}
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
	masterColumn.AutoFill = updatedMasterColumn.AutoFill
	masterColumn.FillQuery = updatedMasterColumn.FillQuery
	masterColumn.Disable = updatedMasterColumn.Disable

	if err := connection.DB.Save(&masterColumn).Error; err != nil {
		return result, errors.New("failed to update Master Column")
	}

	return masterColumn, nil
}

func DeleteMasterColumn(claims, masterColumnId string) (err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return err
	}

	var masterColumn models.MasterColumn

	masterColumn.ModelMasterForm = utils.SoftDelete(user.Username)
	return connection.DB.Model(&masterColumn).Where("id = ?", masterColumnId).Updates(&masterColumn).Error
}

func GetFormColumn(masterTableId string) (result models.TableForm, err error) {

	if err := connection.DB.
		Select("description").
		Model(&models.MasterTable{}).
		Where("id = ?", masterTableId).Scan(&result.FormName).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	rows, err := connection.DB.
		Select("master_columns.id AS field_id, description AS name, english_description AS name_en ,field_name, is_mandatory, need_first_empty, ut.name as ui_type, ui_source_type, ui_source_query, code_group_id, auto_fill, fill_query, disable",
			"(CASE WHEN field_type = 'N' OR field_type = 'F' THEN 'number' ELSE ut.name_ui END) AS ui_name",
			"(CASE WHEN field_type = 'N' THEN 1 WHEN field_type = 'F' THEN 0.01 ELSE NULL END) AS ui_step").
		Joins("JOIN ui_types ut ON ut.id = master_columns.ui_type_id").
		Model(&models.MasterColumn{}).
		Where("table_id = ?", masterTableId).
		Order("sequence").
		Rows()

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

		if data.AutoFill {
			err = connection.DB.
				Raw(fmt.Sprintf(`select * from(select %s) t`, data.FillQuery)).
				Scan(&data.Fill).Error
			if err != nil {
				return result, err
			}
		}

		var UiSource []models.DropdownEn

		if data.UiSourceType == "C" {
			err = connection.DB.
				Select("code, description, english_description").
				Model(&models.MasterCode{}).
				Where("code_group_id = ?", data.CodeGroupId).
				Order("sequence").
				Scan(&UiSource).Error
			if err != nil {
				return result, err
			}
			// if data.NeedFirstEmpty {
			// 	UiSource = utils.Prepend(UiSource, models.DropdownEn{Code: "", Description: "", EnglishDescription: ""})
			// }

		} else if data.UiSourceType == "Q" {
			err = connection.DB.
				Raw(fmt.Sprintf(`select * from(%s) t order by "code"`, data.UiSourceQuery)).
				Scan(&UiSource).Error
			if err != nil {
				return result, err
			}
			// if data.NeedFirstEmpty {
			// 	UiSource = utils.Prepend(UiSource, models.DropdownEn{Code: "", Description: "", EnglishDescription: ""})
			// }
		}

		data.UiSource = UiSource

		result.Form = append(result.Form, data)
	}

	return
}

func GetFormColumnCh(masterTableId string, dataCh chan<- models.TableForm) {

	var result models.TableForm

	if err := connection.DB.
		Select("description").
		Model(&models.MasterTable{}).
		Where("id = ?", masterTableId).Scan(&result.FormName).Error; err != nil {
		return 
	}

	rows, err := connection.DB.Debug().
		Select("master_columns.id AS field_id, description AS name, english_description AS name_en ,field_name, is_mandatory, need_first_empty, ut.name as ui_type, ui_source_type, ui_source_query, code_group_id, auto_fill, fill_query, disable",
			"(CASE WHEN field_type = 'N' OR field_type = 'F' THEN 'number' ELSE ut.name_ui END) AS ui_name",
			"(CASE WHEN field_type = 'N' THEN 1 WHEN field_type = 'F' THEN 0.01 ELSE NULL END) AS ui_step").
		Joins("JOIN ui_types ut ON ut.id = master_columns.ui_type_id").
		Model(&models.MasterColumn{}).
		Where("table_id = ?", masterTableId).
		Order("sequence").
		Rows()

	var data models.FormList

	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = connection.DB.ScanRows(rows, &data)
		if err != nil {
			return
		}

		if data.AutoFill {
			err = connection.DB.
				Raw(fmt.Sprintf(`select * from(select %s) t`, data.FillQuery)).
				Scan(&data.Fill).Error
			if err != nil {
				return
			}
		}

		var UiSource []models.DropdownEn

		if data.UiSourceType == "C" {
			err = connection.DB.
				Select("code, description, english_description").
				Model(&models.MasterCode{}).
				Where("code_group_id = ?", data.CodeGroupId).
				Order("sequence").
				Scan(&UiSource).Error
			if err != nil {
				return 
			}
			// if data.NeedFirstEmpty {
			// 	UiSource = utils.Prepend(UiSource, models.DropdownEn{Code: "", Description: "", EnglishDescription: ""})
			// }

		} else if data.UiSourceType == "Q" {
			err = connection.DB.
				Raw(fmt.Sprintf(`select * from(%s) t order by "code"`, data.UiSourceQuery)).
				Scan(&UiSource).Error
			if err != nil {
				return
			}
			// if data.NeedFirstEmpty {
			// 	UiSource = utils.Prepend(UiSource, models.DropdownEn{Code: "", Description: "", EnglishDescription: ""})
			// }
		}

		data.UiSource = UiSource

		result.Form = append(result.Form, data)
	}

	dataCh <- result
	close(dataCh)
}

func CheckQuery(data models.CheckQuery) (err error) {
	return connection.DB.Raw(`?`, data.Query).Error
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
