package templates

func ModelTemplate() string {
	return `package models

import (
	"time"
)

type {{.TableName | ToCamel}} struct {
	Id int ` + "`gorm:\"primaryKey;autoIncrement\" json:\"id\"`" + `
	CreatedBy    string     ` + "`gorm:\"column:created_by\" json:\"created_by\"`" + `
	CreatedDate  time.Time  ` + "`gorm:\"column:created_date\" json:\"created_date\"`" + `
	UpdatedBy    string     ` + "`gorm:\"column:updated_by\" json:\"updated_by\"`" + `
	UpdatedDate  time.Time  ` + "`gorm:\"column:updated_date\" json:\"updated_date\"`" + `
	IsDelete     string       ` + "`gorm:\"column:is_delete\" json:\"is_delete\"`" + `
	DeletedBy    string     ` + "`gorm:\"column:deleted_by\" json:\"deleted_by\"`" + `
	DeletedDate  time.Time ` + "`gorm:\"column:deleted_date\" json:\"deleted_date\"`" + `
	{{range .Columns}}    {{.FieldName | ToCamel}} {{mapFieldTypeModel .FieldType}} ` + "`gorm:\"column:{{.FieldName}}\" json:\"{{.FieldName}}\"`" + `
{{end}}}

func ({{.TableName | ToCamel}}) TableName() string {
	return "{{.TableName}}"
}
	`
}
