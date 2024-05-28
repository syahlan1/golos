package models

import "time"

type MasterTable struct {
	Id                 int       `json:"id"`
	CreatedDate        time.Time `json:"created_date"`
	CreatedBy          string    `json:"created_by"`
	Status             string    `json:"status"`
	UpdatedDate        time.Time `json:"updated_date"`
	UpdatedBy          string    `json:"updated_by"`
	Description        string    `json:"description"`
	EnglishDescription string    `json:"english_description"`
	OrderField         string    `json:"order_field"`
	TableName          string    `json:"table_name"`
	// ModuleId           int       `json:"module_id"`
	ModuleName    string `json:"module_name"`
	UseWorkflow   int    `json:"use_workflow"`
	FormType      string `json:"form_type"`
	UseBranch     int    `json:"use_branch"`
	UsePeriod     int    `json:"use_period"`
	UseMessage    int    `json:"use_message"`
	UseDataLoader int    `json:"use_data_loader"`
	PeriodType    string `json:"period_type"`
	TableId       int    `json:"table_id"`
}

type CreateMasterTable struct {
	TableName          string `json:"table_name"`
	Description        string `json:"description"`
	EnglishDescription string `json:"english_description"`
	OrderField         string `json:"order_field"`
	FormType           string `json:"form_type"`
	PeriodType         string `json:"period_type"`
	UsePeriod          int    `json:"use_period"`
	UseWorkflow        int    `json:"use_workflow"`
	UseBranch          int    `json:"use_branch"`
	UseDataLoader      int    `json:"use_data_loader"`
}

type MasterColumn struct {
	Id                 int                  `json:"id"`
	CreatedDate        time.Time            `json:"created_date"`
	CreatedBy          string               `json:"created_by"`
	Status             string               `json:"status"`
	UpdatedDate        time.Time            `json:"updated_date"`
	UpdatedBy          string               `json:"updated_by"`
	CodeGroupId        int                  `json:"code_group_id"`
	Description        string               `json:"description"`
	EnglishDescription string               `json:"english_description"`
	FieldLength        int                  `json:"field_length"`
	FieldName          string               `json:"field_name"`
	FieldType          string               `json:"field_type"`
	IsMandatory        bool                 `json:"is_mandatory"`
	Sequence           int                  `json:"sequence"`
	TableId            int                  `json:"table_id"`
	UiSourceQuery      string               `json:"ui_source_query"`
	UiSourceType       string               `json:"ui_source_type"`
	UiType             string               `json:"ui_type"`
	IsExport           bool                 `json:"es_export"`
	CodeGroup          string               `json:"code_group"`
	IsNegative         int                  `json:"is_negative"`
	SqlFunction        string               `json:"sql_function"`
	OnblurScript       string               `json:"onblur_script"`
	MasterMapperColumn []MasterMapperColumn `gorm:"foreignKey:ColumnId"`
	MasterSourceColumn []MasterSourceColumn `gorm:"foreignKey:ColumnId"`
}

type CreateMasterColumn struct {
	FieldName          string `json:"field_name"`
	FieldType          string `json:"field_type"`
	FieldLength        int    `json:"field_length"`
	Description        string `json:"description"`
	EnglishDescription string `json:"english_description"`
	IsMandatory        bool   `json:"is_mandatory"`
	Sequence           int    `json:"sequence"`
	TableId            int    `json:"table_id"`
	UiSourceQuery      string `json:"ui_source_query"`
	UiSourceType       string `json:"ui_source_type"`
	UiType             string `json:"ui_type"`
	IsExport           bool   `json:"es_export"`
	SqlFunction        string `json:"sql_function"`
	OnblurScript       string `json:"onblur_script"`
}
type MasterSourceColumn struct {
	Id           int       `json:"id"`
	CreatedDate  time.Time `json:"created_date"`
	CreatedBy    string    `json:"created_by"`
	Status       string    `json:"status"`
	UpdatedDate  time.Time `json:"updated_date"`
	UpdatedBy    string    `json:"updated_by"`
	ColumnId     int       `json:"column_id"`
	SourceField  string    `json:"source_field"`
	SqlFunction  string    `json:"sql_function"`
	TableId      int       `json:"table_id"`
	SourceNumber int       `json:"source_number"`
	Sequence     int       `json:"sequence"`
}

type MasterMapperColumn struct {
	Id            int    `json:"id"`
	AttributeName string `json:"attribute_name"`
	SystemValue   string `json:"system_value"`
	MapperTableId int    `json:"mapper_table_id"`
	ClassPath     string `json:"class_path"`
	DeclareField  string `json:"declare_field"`
	Sequence      int    `json:"sequence"`
	ColumnId      int    `json:"column_id"`
}

type MasterMapperTable struct {
	Id                 int                  `json:"id"`
	CreatedDate        time.Time            `json:"created_date"`
	CreatedBy          string               `json:"created_by"`
	Status             string               `json:"status"`
	UpdatedDate        time.Time            `json:"updated_date"`
	UpdatedBy          string               `json:"updated_by"`
	MethodType         string               `json:"method_type"`
	TableId            int                  `json:"table_id"`
	ProcessId          int                  `json:"process_id"`
	MasterMapperColumn []MasterMapperColumn `gorm:"foreignKey:MapperTableId"`
}
