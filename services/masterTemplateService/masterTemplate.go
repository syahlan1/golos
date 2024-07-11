package masterTemplateService

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
)

func ShowMasterTemplate(schema, tableName, username, tableGroup, approval, approvalId, id string) (result []map[string]interface{}, err error) {

	tableId := tableName

	var columnSource []models.ColumnSource
	err = connection.DB.
		Select("*").
		Model(&models.MasterColumn{}).
		Where("table_id = ? AND deleted_at is null", tableName).
		Where("ui_source_type = ? OR ui_source_type = ?", "Q", "C").
		Find(&columnSource).Error

	err = connection.DB.
		Select("database_name, table_name").
		Joins("JOIN master_modules mm ON mm.id = master_tables.module_id").
		Model(&models.MasterTable{}).
		Where("master_tables.id = ? AND mm.id = ? AND mm.deleted_at is null", tableName, schema).Row().Scan(&schema, &tableName)
	if err != nil {
		return nil, err
	}

	column := FindColumn(tableId, tableName, true)

	checkTableGroup, err := CheckTableGroup(schema, tableName)
	if err != nil {
		return nil, err
	}

	db := connection.DB.Debug().
		Select(column).
		Table(schema + "." + tableName).
		Where(tableName + ".deleted_date is null").
		Order(tableName + ".id")

	if approval != "" {
		db = db.Select(column,"tgis.id AS id_status").
		Joins("JOIN table_group_item_statuses tgis on tgis.id = "+tableName+".item_status_id AND tgis.status = ?", approval)
	}

	if approvalId != "" {
		db = db.Where("item_status_id = ?", approvalId)
	}

	if id != "" {
		db = db.Where("id = ?", id)
	}

	if username != "" {
		db = db.Where("created_by = ?", username)
	}


	if checkTableGroup {
		if tableGroup == "" {
			db = db.Where(schema + "_group_id is null")
		} else {
			db = db.Where(schema+"_group_id = ?", tableGroup)
		}
	}

	var rows *sql.Rows
	rows, err = db.Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var data map[string]interface{}
		err = connection.DB.ScanRows(rows, &data)
		if err != nil {
			return nil, err
		}

		// log.Println(data[columnSource[0].FieldName])
		datas := make(map[string]interface{})
		for _, v := range columnSource {
			if _, ok := data[v.FieldName]; ok {
				datas = nil
				// var datas map[string]interface{}
				if v.UiSourceType == "C" {
					err = connection.DB.
						Select("code, description, english_description").
						Model(&models.MasterCode{}).
						Where("code_group_id = ? AND code = ?", v.CodeGroupId, utils.InterfaceToString(data[v.FieldName])).
						Order("sequence").
						Scan(&datas).Error
					if err != nil {
						return nil, err
					}

					data[v.FieldName+"_data"] = datas
				} else if v.UiSourceType == "Q" {
					err = connection.DB.
						Raw(fmt.Sprintf(`select * from(%s) t where "code" = ?`, *v.UiSourceQuery), utils.InterfaceToString(data[v.FieldName])).
						Scan(&datas).Error
					if err != nil {
						return nil, err
					}

					// if len(datas) == 0 {
					// 	datas["code"] = ""
					// 	datas["description"] = ""
					// 	datas["english_description"] = ""
					// }

					data[v.FieldName+"_data"] = datas
				}
			}
			// if data["id"] == v.Id {
			// 	fmt.Println(v.CodeGroupId)
			// }
		}

		result = append(result, data)
	}

	return result, nil
}

func CreateMasterTemplate(schema, tableName, username, tableGroup string, data map[string]interface{}) (err error, errValidation []models.Validate) {

	var columnId []int

	connection.DB.Model(&models.MasterColumn{}).Select("id").Where("table_id = ?", tableName).Find(&columnId)

	var validations []models.Validate
	connection.DB.
		Select("master_validations.*, mc.field_name AS column_name").
		Joins("JOIN master_columns mc ON mc.id = master_validations.column_id").
		Model(&models.MasterValidation{}).
		Where("is_active = ? AND column_id IN ? AND mc.deleted_at is null", 1, columnId).Find(&validations)

	if errorMessages, err := utils.ApplyValidations2(connection.DB, data, validations); err != nil {
		return errors.New("Validation errors: " + fmt.Sprintf("%v", errorMessages)), errorMessages
	}

	err = connection.DB.
		Select("database_name, table_name").
		Joins("JOIN master_modules mm ON mm.id = master_tables.module_id").
		Model(&models.MasterTable{}).
		Where("master_tables.id = ? AND mm.id = ? AND mm.deleted_at is null", tableName, schema).Row().Scan(&schema, &tableName)
	if err != nil {
		return err, nil
	}

	var placeholders, columns []string
	var values []interface{}

	data["created_date"] = utils.GetDateTimeNow()
	data["created_by"] = username

	for key, value := range data {
		columns = append(columns, key)
		values = append(values, value)
		placeholders = append(placeholders, "?")
	}

	if tableGroup != "" {
		columns = append(columns, schema+"_group_id")
		values = append(values, tableGroup)
		placeholders = append(placeholders, "?")
	}

	query := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s) RETURNING id",
		schema, tableName,
		strings.Join(columns, ", "), strings.Join(placeholders, ", "))

	var id int
	err = connection.DB.Debug().Raw(query, values...).Row().Scan(&id)
	if err != nil {
		return err, nil
	}

	return
}

func UpdateMasterTemplate(schema, tableName, id, username string, data map[string]interface{}) (err error, errValidation []models.Validate) {

	var columnId []int

	connection.DB.Model(&models.MasterColumn{}).Select("id").Where("table_id = ?", tableName).Find(&columnId)

	var validations []models.Validate
	connection.DB.
		Select("master_validations.*, mc.field_name AS column_name").
		Joins("JOIN master_columns mc ON mc.id = master_validations.column_id").
		Model(&models.MasterValidation{}).
		Where("is_active = ? AND column_id IN ? AND mc.deleted_at is null", 1, columnId).Find(&validations)

	if errorMessages, err := utils.ApplyValidations2(connection.DB, data, validations); err != nil {
		return errors.New("Validation errors: " + fmt.Sprintf("%v", errorMessages)), errorMessages
	}

	err = connection.DB.
		Select("database_name, table_name").
		Joins("JOIN master_modules mm ON mm.id = master_tables.module_id").
		Model(&models.MasterTable{}).
		Where("master_tables.id = ? AND mm.id = ? AND mm.deleted_at is null", tableName, schema).Row().Scan(&schema, &tableName)
	if err != nil {
		return err, nil
	}

	data["updated_date"] = utils.GetDateTimeNow()
	data["updated_by"] = username

	err = connection.DB.Table(schema+"."+tableName).Where("id = ?", id).Updates(data).Error

	if err != nil {
		return err, nil
	}

	return
}

func DeleteMasterTemplate(schema, tableName, id, username string) (err error) {

	data := make(map[string]interface{})

	err = connection.DB.
		Select("database_name, table_name").
		Joins("JOIN master_modules mm ON mm.id = master_tables.module_id").
		Model(&models.MasterTable{}).
		Where("master_tables.id = ? AND mm.id = ? AND mm.deleted_at is null", tableName, schema).Row().Scan(&schema, &tableName)
	if err != nil {
		return err
	}

	data["deleted_date"] = utils.GetDateTimeNow()
	data["updated_by"] = username

	err = connection.DB.Table(schema+"."+tableName).Where("id = ?", id).Updates(data).Error

	if err != nil {
		return err
	}

	return
}

func FindColumn(tableId, tableName string, withId bool) (result string) {
	var column string

	rows, err := connection.DB.
		Select("field_name").
		Model(&models.MasterColumn{}).
		Where("table_id = ?", tableId).
		Rows()

	if err != nil {
		return ""
	}
	defer rows.Close()

	if withId {
		result = tableName + ".id, "
	}

	for rows.Next() {

		if err := rows.Scan(&column); err != nil {
			return ""
		}

		result += tableName + "." + column + ", "
	}

	return result[:len(result)-2]
}

func CheckTableGroup(schema, tableName string) (check bool, err error) {

	err = connection.DB.Raw("SELECT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_schema = ? AND table_name = ? AND column_name = ?)", schema, tableName, schema+"_group_id").Scan(&check).Error
	if err != nil {
		return
	}

	return check, nil
}
