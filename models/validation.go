package models

import "time"

type MasterValidation struct {
	Id                 int          `json:"id" gorm:"autoIncrement"`
	CreatedDate        time.Time    `json:"created_date"`
	CreatedBy          string       `json:"created_by"`
	UpdatedDate        time.Time    `json:"updated_date"`
	UpdatedBy          string       `json:"updated_by"`
	Status             string       `json:"status"`
	ColumnId           int          `json:"column_id"`
	Description        string       `json:"description"`
	EnglishDescription string       `json:"english_description"`
	MasterCodeId       int          `json:"master_code_id"`
	MessageType        string       `json:"message_type"`
	ValidationFunction string       `json:"validation_function"`
	IsActive           int          `json:"is_active"`
	MasterColumn       MasterColumn `json:"-" gorm:"foreignKey:ColumnId"`
	MasterCode         MasterCode   `json:"-" gorm:"foreignKey:MasterCodeId"`
}

type CreateValidation struct {
	Description        string `json:"description"`
	EnglishDescription string `json:"english_description"`
	MessageType        string `json:"message_type"`
	ValidationFunction string `json:"validation_function"`
	IsActive           int    `json:"is_active"`
	MasterCodeId       int    `json:"master_code_id"`
}

type MasterValidationRelation struct {
	Id              int       `json:"id"`
	CreatedDate     time.Time `json:"created_Date"`
	CreatedBy       string    `json:"created_by"`
	UpdatedDate     time.Time `json:"updated_date"`
	UpdatedBy       string    `json:"updated_by"`
	Status          string    `json:"status"`
	FieldName       string    `json:"field_name"`
	IsActive        int       `json:"is_aactive"`
	TableId         int       `json:"table_id"`
	ValidationQuery string    `json:"validation_query"`
}
