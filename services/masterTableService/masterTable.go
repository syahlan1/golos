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

func CreateMasterTable(claims string, data models.MasterTable) (err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return err
	}

	data.CreatedBy = user.Username

	if err := connection.DB.Create(&data).Error; err != nil {
		return errors.New("failed to create Master Code")
	}

	// Return success response
	return nil
}

func ShowMasterTable(moduleId string) (result []models.MasterTable) {
	var masterCode []models.MasterTable

	connection.DB.Select("master_tables.*, md.module_name").
		Joins("JOIN master_modules md ON md.id = master_tables.module_id").
		Find(&masterCode, "module_id = ?", moduleId)

	return masterCode
}

func ShowMasterTableDetail(masterTableId string) (result models.MasterTable, err error) {
	var masterTable models.MasterTable

	// Mencari detail MasterTable berdasarkan id
	if err := connection.DB.Select("master_tables.*, md.module_name").
		Joins("JOIN master_modules md ON md.id = master_tables.module_id").
		Where("master_tables.id = ?", masterTableId).
		First(&masterTable).Error; err != nil {
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
	masterTable.UpdatedAt = time.Now()
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

func DeleteMasterTable(claims, masterTableId string) (err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return err
	}

	var masterTable models.MasterTable

	masterTable.ModelMasterForm = utils.SoftDelete(user.Username)
	return connection.DB.Model(&masterTable).Where("id = ?", masterTableId).Updates(&masterTable).Error
}

func GenerateTable(tableID string) (err error) {
	db := connection.DB

	var masterTable models.MasterTable
	if err := db.Select("master_tables.*, md.database_name as module_name").
		Joins("JOIN master_modules md ON md.id = master_tables.module_id").
		First(&masterTable, tableID).Error; err != nil {
		return errors.New("Data not found")
	}

	// log.Println("Generating table", masterTable)
	// return

	var columns []models.MasterColumn
	db.Where("table_id = ?", masterTable.Id).Find(&columns)

	// Define the mandatory columns
	mandatoryColumns := []models.MasterColumn{
		{FieldName: "created_by", FieldType: "VARCHAR", FieldLength: 255},
		{FieldName: "created_date", FieldType: "TIMESTAMP"},
		{FieldName: "updated_by", FieldType: "VARCHAR", FieldLength: 255},
		{FieldName: "updated_date", FieldType: "TIMESTAMP"},
		{FieldName: "is_delete", FieldType: "VARCHAR", FieldLength: 255},
		{FieldName: "deleted_by", FieldType: "VARCHAR", FieldLength: 255},
		{FieldName: "deleted_date", FieldType: "TIMESTAMP"},
	}

	var SchemaExists bool
	err = db.Raw("SELECT EXISTS (SELECT FROM information_schema.schemata WHERE schema_name = ?)", masterTable.ModuleName).Scan(&SchemaExists).Error
	if err != nil {
		return err
	}

	if !SchemaExists {
		// Create schema
		err = db.Exec(fmt.Sprintf(`CREATE SCHEMA "%s"`, masterTable.ModuleName)).Error
		if err != nil {
			return err
		}
	}

	var tableExists bool
	err = db.Raw("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_schema = ? AND table_name = ?)", masterTable.ModuleName, masterTable.TableName).Scan(&tableExists).Error
	if err != nil {
		return err
	}

	if !tableExists {
		createTableSQL := fmt.Sprintf(`CREATE TABLE "%s"."%s" (`+"\n", masterTable.ModuleName, masterTable.TableName)
		createTableSQL += "\tID SERIAL PRIMARY KEY,\n"

		// Add mandatory columns to createTableSQL
		for _, column := range mandatoryColumns {
			fieldSQL := fmt.Sprintf("\t%s %s", column.FieldName, utils.MapFieldType(column.FieldType, column.FieldLength))
			createTableSQL += fieldSQL + ",\n"
		}

		for _, column := range columns {
			fieldSQL := fmt.Sprintf("\t%s %s", column.FieldName, utils.MapFieldType(column.FieldType, column.FieldLength))
			if column.IsMandatory {
				fieldSQL += " NOT NULL"
			}
			createTableSQL += fieldSQL + ",\n"
		}

		createTableSQL = createTableSQL[:len(createTableSQL)-2] + "\n);"

		// log.Println("createTableSQL", createTableSQL)
		// return
		if err := db.Exec(createTableSQL).Error; err != nil {
			return err
		}
	} else {
		var existingColumns []struct {
			ColumnName             string
			DataType               string
			CharacterMaximumLength sql.NullInt32
		}
		rows, err := db.Raw(fmt.Sprintf("SELECT column_name, data_type, character_maximum_length FROM information_schema.columns WHERE table_schema = '%s' AND table_name = '%s'",masterTable.ModuleName, masterTable.TableName)).Rows()
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

		// Check and add mandatory columns if not exist
		for _, column := range mandatoryColumns {
			if !containsColumn(existingColumns, column.FieldName) {
				fieldSQL := fmt.Sprintf("%s %s", column.FieldName, utils.MapFieldType(column.FieldType, column.FieldLength))
				alterTableQueries = append(alterTableQueries, fmt.Sprintf("ADD COLUMN %s", fieldSQL))
			}
		}

		for _, column := range columns {
			existingColumn, exists := findExistingColumn(existingColumns, column.FieldName)
			if !exists {
				fieldSQL := fmt.Sprintf("%s %s", column.FieldName, utils.MapFieldType(column.FieldType, column.FieldLength))
				if column.IsMandatory {
					fieldSQL += " NOT NULL"
				}
				alterTableQueries = append(alterTableQueries, fmt.Sprintf("ADD COLUMN %s", fieldSQL))
			} else {
				if needsModification(existingColumn, column) {
					modifySQL := fmt.Sprintf("%s TYPE %s", column.FieldName, utils.MapFieldType(column.FieldType, column.FieldLength))
					alterTableQueries = append(alterTableQueries, fmt.Sprintf("ALTER COLUMN %s", modifySQL))
				}
			}
		}

		if len(alterTableQueries) > 0 {
			alterTableSQL := fmt.Sprintf(`ALTER TABLE "%s"."%s"`+"\n"+`%s;`, masterTable.ModuleName, masterTable.TableName, strings.Join(alterTableQueries, ",\n"))
			if err := db.Exec(alterTableSQL).Error; err != nil {
				return err
			}
		}
	}

	// Generate Model and CRUD handlers
	// GenerateModel(masterTable, columns)
	// GenerateServiceHandler(masterTable, columns)
	// GenerateControllerHandler(masterTable, columns)
	// GenerateRouteHandler(masterTable)

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

func containsColumn(existingColumns []struct {
	ColumnName             string
	DataType               string
	CharacterMaximumLength sql.NullInt32
}, columnName string) bool {
	for _, column := range existingColumns {
		if column.ColumnName == columnName {
			return true
		}
	}
	return false
}

func needsModification(existingColumn struct {
	ColumnName             string
	DataType               string
	CharacterMaximumLength sql.NullInt32
}, newColumn models.MasterColumn) bool {
	if existingColumn.DataType != utils.MapFieldType(newColumn.FieldType, newColumn.FieldLength) {
		return true
	}
	if newColumn.FieldType == "A" && existingColumn.CharacterMaximumLength.Valid && existingColumn.CharacterMaximumLength.Int32 != int32(newColumn.FieldLength) {
		return true
	}
	return false
}

func containsMasterColumn(columns []models.MasterColumn, fieldName string) bool {
	for _, column := range columns {
		if column.FieldName == fieldName {
			return true
		}
	}
	return false
}

func GenerateRouteHandler(masterTable models.MasterTable) {
	log.Println("Starting to generate route handlers")

	filePath := "routes/routes.go"
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer f.Close()

	log.Println("Adding route and import")
	err = templates.AddRouteAndImport(masterTable.TableName, utils.ToKebabCase(masterTable.TableName), "github.com/syahlan1/golos")
	if err != nil {
		log.Fatalf("Failed to add route and import: %v", err)
	}

	log.Printf("Route for %s added successfully", masterTable.TableName)
}

// generateModel generates the Go model file
func GenerateModel(masterTable models.MasterTable, columns []models.MasterColumn) {
	modelTemplate := templates.ModelTemplate()

	funcMap := template.FuncMap{
		"ToCamel":           utils.ToCamelCase,
		"ToLowerCamel":      utils.ToLowerCamelCase,
		"mapFieldTypeModel": utils.MapFieldTypeModel,
	}

	tmpl, err := template.New("model").Funcs(funcMap).Parse(modelTemplate)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(filepath.Join("models", utils.ToLowerCamelCase(masterTable.TableName)+".go"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data := struct {
		ModuleName  string
		TableName   string
		PackagePath string
		Columns     []models.MasterColumn
	}{
		ModuleName:  masterTable.ModuleName,
		TableName:   masterTable.TableName,
		PackagePath: "github.com/syahlan1/golos",
		Columns:     columns,
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		panic(err)
	}
}

// generateCRUDHandlers generates the CRUD handler files
func GenerateServiceHandler(masterTable models.MasterTable, columns []models.MasterColumn) {
	log.Println("Starting to generate service handlers")
	handlerTemplate := templates.ServiceTemplate()

	funcMap := template.FuncMap{
		"ToCamel":           utils.ToCamelCase,
		"ToLowerCamel":      utils.ToLowerCamelCase,
		"mapFieldTypeModel": utils.MapFieldTypeModel,
	}

	tmpl, err := template.New("service").Funcs(funcMap).Parse(handlerTemplate)
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	handlersDir := filepath.Join("services", utils.ToLowerCamelCase(masterTable.TableName)+"Service")
	err = utils.CreateDirIfNotExist(handlersDir)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	filePath := filepath.Join(handlersDir, utils.ToLowerCamelCase(masterTable.TableName)+".go")
	log.Printf("Creating file at %s", filePath)
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
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

	log.Println("Executing template")
	err = tmpl.Execute(f, data)
	if err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("Service for %s generated successfully at %s", masterTable.TableName, filePath)
}

func GenerateControllerHandler(masterTable models.MasterTable, columns []models.MasterColumn) {
	log.Println("Starting to generate controller handlers")
	handlerTemplate := templates.ControllerTemplate()

	funcMap := template.FuncMap{
		"ToCamel":           utils.ToCamelCase,
		"ToLowerCamel":      utils.ToLowerCamelCase,
		"mapFieldTypeModel": utils.MapFieldTypeModel,
	}

	tmpl, err := template.New("controller").Funcs(funcMap).Parse(handlerTemplate)
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	handlersDir := filepath.Join("controllers", utils.ToLowerCamelCase(masterTable.TableName)+"Controller")
	err = utils.CreateDirIfNotExist(handlersDir)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	filePath := filepath.Join(handlersDir, utils.ToLowerCamelCase(masterTable.TableName)+"Controller"+".go")
	log.Printf("Creating file at %s", filePath)
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
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

	log.Println("Executing template")
	err = tmpl.Execute(f, data)
	if err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("Controller for %s generated successfully at %s", masterTable.TableName, filePath)
}
