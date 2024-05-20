package models

import "time"

type MasterParameter struct {
	Id                 int       `json:"id"`
	CreatedDate        time.Time `json:"created_date"`
	CreatedBy          string    `json:"created_by"`
	UpdatedDate        time.Time `json:"updated_date"`
	UpdatedBy          string    `json:"updated_by"`
	Status             string    `json:"status"`
	IsEncrypted        int       `json:"is_encrypted"`
	ModuleId           int       `json:"module_id"`
	Description        string    `json:"description"`
	EnglishDescription string    `json:"english_description"`
	ParamKey           string    `json:"param_key"`
	ParamValue         string    `json:"param_value"`
}
