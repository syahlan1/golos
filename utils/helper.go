package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/syahlan1/golos/models"
	"gorm.io/gorm"
)

// Helper function to get string value from map
func GetStringValue(data map[string]interface{}, key string) string {
	value, ok := data[key].(string)
	if !ok {
		return "" // return empty string if not found or not a string
	}
	return value
}

// Helper function to get int value from map
func GetIntValue(data map[string]interface{}, key string) int {
	value, ok := data[key].(float64)
	if !ok {
		return 0 // return 0 if not found or not a number
	}
	return int(value)
}

// Helper function to get bool value from map
func GetBoolValue(data map[string]interface{}, key string) bool {
	value, ok := data[key].(bool)
	return ok && value
}

// createDirIfNotExist creates a directory if it does not exist
func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func ToCamelCase(input string) string {
	words := strings.Split(input, "_")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
}

func ToLowerCamelCase(input string) string {
	words := strings.Split(input, "_")
	for i, word := range words {
		if i == 0 {
			words[i] = strings.ToLower(word)
		} else {
			words[i] = strings.Title(word)
		}
	}
	return strings.Join(words, "")
}

func ToKebabCase(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '-')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

// mapFieldType berfungsi untuk memetakan tipe data dari MasterColumn ke tipe data SQL
func MapFieldType(fieldType string, fieldLength int) string {
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
	case "TIMESTAMP":
		return "TIMESTAMP"
	default:
		return "VARCHAR"
	}
}

func MapFieldTypeModel(fieldType string) string {
	switch fieldType {
	case "A":
		return "string"
	case "N":
		return "int"
	case "D":
		return "time.Time"
	case "TIMESTAMP":
		return "time.Time"
	default:
		return "string" // Default to string if unknown type
	}
}

func GetColumnIds(columns []models.MasterColumn) []int {
	var columnIds []int
	for _, column := range columns {
		columnIds = append(columnIds, column.Id)
	}
	return columnIds
}

func ApplyValidations(db *gorm.DB, data map[string]interface{}, columns []models.MasterColumn, validations []models.MasterValidation) (map[int][]string, error) {
	errorMessages := make(map[int][]string)
	columnIds := GetColumnIds(columns)

	for _, validation := range validations {
		if !contains(columnIds, validation.ColumnId) {
			continue
		}

		var columnName string
		for _, column := range columns {
			if column.Id == validation.ColumnId {
				columnName = column.FieldName
				break
			}
		}

		fieldValue, exists := data[columnName]
		if !exists {
			continue
		}

		fieldValueStr, ok := fieldValue.(string)
		if !ok {
			errorMessages[validation.ColumnId] = append(errorMessages[validation.ColumnId], "Field value is not a string")
			continue
		}

		isValid, err := executeValidationFunction(db, fieldValueStr, validation.ValidationFunction)
		if err != nil {
			errorMessages[validation.ColumnId] = append(errorMessages[validation.ColumnId], fmt.Sprintf("Error validating field %s: %s", columnName, err.Error()))
			continue
		}
		if !isValid {
			errorMessages[validation.ColumnId] = append(errorMessages[validation.ColumnId], validation.EnglishDescription)
		}
	}

	if len(errorMessages) > 0 {
		return errorMessages, errors.New("validation errors")
	}
	return nil, nil
}

// contains adalah fungsi bantuan untuk memeriksa apakah suatu elemen ada dalam slice
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ExecuteValidationFunction executes the validation function as SQL query
func executeValidationFunction(db *gorm.DB, value string, validationFunction string) (bool, error) {
	valueEscaped := strings.ReplaceAll(value, "'", "''")
	validationQuery := strings.Replace(validationFunction, "#", valueEscaped, -1)
	validationQuery = "SELECT " + validationQuery

	// Tambahkan log untuk melihat query yang dihasilkan
	fmt.Println("Generated SQL query:", validationQuery)

	var result string
	err := db.Raw(validationQuery).Scan(&result).Error
	if err != nil {
		return false, err
	}

	return result == "valid", nil
}
