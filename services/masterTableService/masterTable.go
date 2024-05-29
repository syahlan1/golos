package masterTableService

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
	"github.com/syahlan1/golos/utils/templates"
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

	var masterTable models.MasterTable
	if err := db.First(&masterTable, tableID).Error; err != nil {
		return errors.New("Data not found")
	}

	var columns []models.MasterColumn
	db.Where("table_id = ?", masterTable.Id).Find(&columns)

	var tableExists bool
	err = db.Raw("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = ?)", masterTable.TableName).Scan(&tableExists).Error
	if err != nil {
		return err
	}

	if !tableExists {
		createTableSQL := fmt.Sprintf("CREATE TABLE %s (\n", masterTable.TableName)
		createTableSQL += "\tID SERIAL PRIMARY KEY,\n"

		for _, column := range columns {
			fieldSQL := fmt.Sprintf("\t%s %s", column.FieldName, utils.MapFieldType(column.FieldType, column.FieldLength))
			if column.IsMandatory {
				fieldSQL += " NOT NULL"
			}
			createTableSQL += fieldSQL + ",\n"
		}

		createTableSQL = createTableSQL[:len(createTableSQL)-2] + "\n);"
		if err := db.Exec(createTableSQL).Error; err != nil {
			return err
		}
	} else {
		var existingColumns []struct {
			ColumnName             string
			DataType               string
			CharacterMaximumLength sql.NullInt32
		}
		rows, err := db.Raw(fmt.Sprintf("SELECT column_name, data_type, character_maximum_length FROM information_schema.columns WHERE table_name = '%s'", masterTable.TableName)).Rows()
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var column struct {
				ColumnName             string
				DataType               string
				CharacterMaximumLength sql.NullInt32
			}
			rows.Scan(&column.ColumnName, &column.DataType, &column.CharacterMaximumLength)
			existingColumns = append(existingColumns, column)
		}

		columnMap := make(map[string]models.MasterColumn)
		for _, column := range columns {
			columnMap[column.FieldName] = column
		}

		var alterTableQueries []string

		for _, column := range columns {
			existingColumn, exists := findExistingColumn(existingColumns, column.FieldName)
			if !exists {
				fieldSQL := fmt.Sprintf("%s %s", column.FieldName, utils.MapFieldType(column.FieldType, column.FieldLength))
				if column.IsMandatory {
					fieldSQL += " NOT NULL"
				}
				alterTableQueries = append(alterTableQueries, fmt.Sprintf("ADD COLUMN %s", fieldSQL))
			} else {
				modifySQL := ""
				if existingColumn.DataType != utils.MapFieldType(column.FieldType, column.FieldLength) {
					modifySQL = fmt.Sprintf("%s TYPE %s", column.FieldName, utils.MapFieldType(column.FieldType, column.FieldLength))
				} else if column.FieldType == "A" && existingColumn.CharacterMaximumLength.Valid && existingColumn.CharacterMaximumLength.Int32 != int32(column.FieldLength) {
					modifySQL = fmt.Sprintf("%s TYPE VARCHAR(%d)", column.FieldName, column.FieldLength)
				}
				if modifySQL != "" {
					alterTableQueries = append(alterTableQueries, fmt.Sprintf("ALTER COLUMN %s", modifySQL))
				}
			}
		}

		for _, existingColumn := range existingColumns {
			if existingColumn.ColumnName != "id" && !containsMasterColumn(columns, existingColumn.ColumnName) {
				alterTableQueries = append(alterTableQueries, fmt.Sprintf("DROP COLUMN %s", existingColumn.ColumnName))
			}
		}

		if len(alterTableQueries) > 0 {
			alterTableSQL := fmt.Sprintf("ALTER TABLE %s\n%s;", masterTable.TableName, strings.Join(alterTableQueries, ",\n"))
			if err := db.Exec(alterTableSQL).Error; err != nil {
				return err
			}
		}
	}

	// Generate Model and CRUD handlers
	GenerateModel(masterTable, columns)
	GenerateCRUDHandlers(masterTable, columns)

	return nil
}

func findExistingColumn(existingColumns []struct {
	ColumnName             string
	DataType               string
	CharacterMaximumLength sql.NullInt32
}, columnName string) (struct {
	ColumnName             string
	DataType               string
	CharacterMaximumLength sql.NullInt32
}, bool) {
	for _, column := range existingColumns {
		if column.ColumnName == columnName {
			return column, true
		}
	}
	return struct {
		ColumnName             string
		DataType               string
		CharacterMaximumLength sql.NullInt32
	}{}, false
}

func containsMasterColumn(columns []models.MasterColumn, fieldName string) bool {
	for _, column := range columns {
		if column.FieldName == fieldName {
			return true
		}
	}
	return false
}

// generateModel generates the Go model file
func GenerateModel(masterTable models.MasterTable, columns []models.MasterColumn) {
	modelTemplate := templates.ModelTemplate()

	funcMap := template.FuncMap{
		"ToCamel":           utils.ToCamelCase,
		"mapFieldTypeModel": utils.MapFieldTypeModel,
	}

	tmpl, err := template.New("model").Funcs(funcMap).Parse(modelTemplate)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(filepath.Join("models", masterTable.TableName+".go"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data := struct {
		TableName string
		Columns   []models.MasterColumn
	}{
		TableName: masterTable.TableName,
		Columns:   columns,
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		panic(err)
	}
}

// generateCRUDHandlers generates the CRUD handler files
func GenerateCRUDHandlers(masterTable models.MasterTable, columns []models.MasterColumn) {
	handlerTemplate := templates.HandlerCRUDTemplate()

	funcMap := template.FuncMap{
		"ToCamel":           utils.ToCamelCase,
		"mapFieldTypeModel": utils.MapFieldTypeModel,
	}

	tmpl, err := template.New("controller").Funcs(funcMap).Parse(handlerTemplate)
	if err != nil {
		panic(err)
	}

	handlersDir := "controllers"
	utils.CreateDirIfNotExist(handlersDir)

	f, err := os.Create(filepath.Join(handlersDir, utils.ToCamelCase(masterTable.TableName+"Controller.go")))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data := struct {
		TableName   string
		PackagePath string
		Columns     []models.MasterColumn
	}{
		TableName:   masterTable.TableName,
		PackagePath: "github.com/syahlan1/golos",
		Columns:     columns,
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		panic(err)
	}
}
