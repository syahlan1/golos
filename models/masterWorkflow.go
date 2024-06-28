package models

type MasterWorkflow struct {
	Id                       int    `json:"id" gorm:"primaryKey"`
	StatusDescription        string `json:"status_description"`
	StatusEnglishDescription string `json:"status_english_description"`
	StatusName               string `json:"status_name"`
	ModelMasterForm          `json:"-"`
}
