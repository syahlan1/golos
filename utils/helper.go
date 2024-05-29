package utils

import (
	"fmt"
	"os"
	"strings"
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
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func ToCamelCase(input string) string {
	words := strings.Split(input, "_")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
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
	default:
		return "string" // Default to string if unknown type
	}
}
