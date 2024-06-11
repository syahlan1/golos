package templates

func ServiceTemplate() string {
	return `package {{ .TableName | ToLowerCamel}}Service

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"{{.PackagePath}}/connection"
	"{{.PackagePath}}/models"
	"{{.PackagePath}}/utils"
)

func ShowDetail{{ .TableName | ToCamel }}(Id string) (result models.{{ .TableName | ToCamel }}) {
	connection.DB.Where("id = ? AND is_delete = ?", Id, "L").First(&result)

	return result
}

func Show{{ .TableName | ToCamel }}() (result []models.{{ .TableName | ToCamel }}, err error) {
	if err := connection.DB.Where("is_delete = ?", "L").Order("GREATEST(created_date, updated_date) DESC").Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func Create{{ .TableName | ToCamel }}(username string, data models.{{ .TableName | ToCamel }}) (name string, err error) {
	dataMap := map[string]interface{}{
		{{range $column := .Columns}}"{{ $column.FieldName }}" : data.{{$column.FieldName | ToCamel}},
		{{end}}
	}

	var columns []models.MasterColumn
	if err := connection.DB.Find(&columns).Error; err != nil {
		return name, errors.New("No master column found")
	}

	var validations []models.MasterValidation
	connection.DB.Where("column_id IN ?", utils.GetColumnIds(columns)).Find(&validations)

	// Apply validations
	if errorMessages, err := utils.ApplyValidations(connection.DB, dataMap, columns, validations); err != nil {
		return name, errors.New("Validation errors: " + fmt.Sprintf("%v", errorMessages))
	}

	{{ .TableName | ToLowerCamel }} := models.{{ .TableName | ToCamel }}{
		{{range $column := .Columns}}{{ $column.FieldName | ToCamel}} : data.{{$column.FieldName | ToCamel}},
		{{end}}	
		CreatedBy:          username,
		CreatedDate:        time.Now(),
		IsDelete:           "L",
	}

	if err := connection.DB.Create(&{{ .TableName | ToLowerCamel }}).Error; err != nil {
		return name, err
	}

	return name, nil
}

func Update{{ .TableName | ToCamel }}(Id string, username string, updatedData models.{{ .TableName | ToCamel }}) (result models.{{ .TableName | ToCamel }}, err error) {
	idInt, err := strconv.Atoi(Id)
	if err != nil {
		return result, errors.New("invalid ID format")
	}

	var data models.{{ .TableName | ToCamel }}
	if err := connection.DB.First(&data, idInt).Error; err != nil {
		return result, errors.New("data not found")
	}

	dataMap := map[string]interface{}{
		{{range $column := .Columns}}"{{ $column.FieldName }}" : updatedData.{{$column.FieldName | ToCamel}},
		{{end}}
	}

	var columns []models.MasterColumn
	if err := connection.DB.Find(&columns).Error; err != nil {
		return result, errors.New("No master column found")
	}

	var validations []models.MasterValidation
	connection.DB.Where("column_id IN ?", utils.GetColumnIds(columns)).Find(&validations)

	// Apply validations
	if errorMessages, err := utils.ApplyValidations(connection.DB, dataMap, columns, validations); err != nil {
		return result, errors.New("Validation errors: " + fmt.Sprintf("%v", errorMessages))
	}

	{{range $column := .Columns}}data.{{ $column.FieldName | ToCamel }} = updatedData.{{$column.FieldName | ToCamel}}
	{{end}}
	data.UpdatedBy = username
	data.UpdatedDate = time.Now()

	if err := connection.DB.Save(&data).Error; err != nil {
		return result, errors.New("failed to update data")
	}

	return data, nil
}

func Delete{{ .TableName | ToCamel }}(username string, Id string) error {
	idInt, err := strconv.Atoi(Id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	var {{ .TableName | ToLowerCamel }} models.{{ .TableName | ToCamel }}
	if err := connection.DB.First(&{{ .TableName | ToLowerCamel }}, idInt).Error; err != nil {
		return errors.New("data not found")
	}

	{{ .TableName | ToLowerCamel }}.IsDelete = "D"
	{{ .TableName | ToLowerCamel }}.DeletedBy = username
	{{ .TableName | ToLowerCamel }}.DeletedDate = time.Now()

	if err := connection.DB.Save(&{{ .TableName | ToLowerCamel }}).Error; err != nil {
		return errors.New("failed to delete data")
	}

	return nil
}
	`
}
