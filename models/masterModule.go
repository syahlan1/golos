package models

type MasterModule struct {
	Id                 int    `json:"id" gorm:"primaryKey"`
	ModuleName         string `json:"module_name"`
	DatabaseName       string `json:"database_name"`
	Description        string `json:"description"`
	EnglishDescription string `json:"english_description"`
	IsActive           bool   `json:"is_active"`
	UseBranch          bool   `json:"use_branch"`
	UsePeriod          bool   `json:"use_period"`
	ModelMasterForm    `json:"-"`
}
