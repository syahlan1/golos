package templates

func ModelTemplate() string {
	return `package models

	import (
		"time"
	)

	type {{.TableName | ToCamel}} struct {
		Id int ` + "`gorm:\"primaryKey;autoIncrement\" json:\"id\"`" + `
	{{range .Columns}}    {{.FieldName | ToCamel}} {{mapFieldTypeModel .FieldType}} ` + "`gorm:\"column:{{.FieldName}}\" json:\"{{.FieldName}}\"`" + `
	{{end}}}
	`
}
