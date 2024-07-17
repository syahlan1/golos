package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
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

func ToDashCase(s string) string {
	replacer := strings.NewReplacer("_", "-")
	result := replacer.Replace(s)

	return strings.ToLower(result)
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

func ApplyValidations2(db *gorm.DB, data map[string]interface{}, validations []models.Validate) (errorMessages []models.Validate, err error) {
	for _, validation := range validations {
		isValid, err := executeValidationFunction(db, InterfaceToString(data[validation.ColumnName]), validation.ValidationFunction)
		if err != nil {
			errorMessages = append(errorMessages, models.Validate{
				ColumnId:           validation.ColumnId,
				ColumnName:         validation.ColumnName,
				Description:        err.Error(),
				EnglishDescription: err.Error(),
			})
			continue
		}
		if !isValid {
			errorMessages = append(errorMessages, models.Validate{
				ColumnId:           validation.ColumnId,
				ColumnName:         validation.ColumnName,
				Description:        validation.Description,
				EnglishDescription: validation.EnglishDescription,
			})
		}
	}

	if len(errorMessages) > 0 {
		return errorMessages, errors.New("validation errors")
	}
	return nil, nil
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

		// fieldValueStr, _ := fieldValue.(string)
		// log.Println(fieldValueStr)
		// if !ok {
		// 	errorMessages[validation.ColumnId] = append(errorMessages[validation.ColumnId], "Field value is not a string")
		// 	continue
		// }

		isValid, err := executeValidationFunction(db, InterfaceToString(fieldValue), validation.ValidationFunction)
		if err != nil {
			errorMessages[validation.ColumnId] = append(errorMessages[validation.ColumnId], fmt.Sprintf("Error validating field %s: %s", columnName, err.Error()))
			continue
		}
		if !isValid {
			errorMessages[validation.ColumnId] = append(errorMessages[validation.ColumnId], []string{"aaa", validation.EnglishDescription}...)
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
	validationQuery := strings.Replace(validationFunction, "#", "'"+valueEscaped+"'", -1)
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

func InterfaceToString(i interface{}) string {
	switch v := i.(type) {
	case int:
		return strconv.Itoa(v)
	case int8, int16, int32, int64:
		return fmt.Sprintf("%d", v)
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%f", v)
	case bool:
		return strconv.FormatBool(v)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}

func Encrypt(plainText []byte) (string, error) {
	key, err := hex.DecodeString(os.Getenv("PARAM_KEY"))
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nil, nonce, plainText, nil)
	return hex.EncodeToString(nonce) + hex.EncodeToString(cipherText), nil
}

func Decrypt(encryptedText string) (string, error) {
	key, err := hex.DecodeString(os.Getenv("PARAM_KEY"))
	if err != nil {
		return "", err
	}
	data, err := hex.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	if len(data) < 12 {
		return "", fmt.Errorf("invalid data size")
	}

	nonce, cipherText := data[:12], data[12:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func IsValidSqlName(name string) error {
	match, _ := regexp.MatchString(`^[a-zA-Z_][a-zA-Z0-9_]*$`, name)
	if !match {
		return errors.New("invalid field_name")
	}

	if len(name) > 63 {
		return errors.New("name too long")
	}
	return nil
}

func IsValidSQL(query string) error {
	match, _ := regexp.MatchString("(?i)\\b(INSERT|UPDATE|DELETE|DROP|CREATE|ALTER|TRUNCATE|GRANT|REVOKE|EXEC|INTO)\\b", query)
	if match {
		return errors.New("invalid characters in query")
	}
	return nil
}