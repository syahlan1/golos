package models

type MasterCode struct {
	Id                 int    `json:"id"`
	Code               string `json:"code"`
	CodeGroupId        int    `json:"code_group_id"`
	Description        string `json:"description"`
	EnglishDescription string `json:"english_description"`
	Sequence           int    `json:"sequence"`
	Parent             *int   `json:"parent"`
	CodeGroup          string `json:"code_group"`
	ModelMasterForm    `json:"-"`
	MasterCodeGroup    MasterCodeGroup `json:"-" gorm:"foreignKey:CodeGroupId"`
}

type CreateMasterCode struct {
	Code               string `json:"code"`
	CodeGroupId        int    `json:"code_group_id"`
	Description        string `json:"description"`
	EnglishDescription string `json:"english_description"`
	Sequence           int    `json:"sequence"`
	CodeGroup          string `json:"code_group"`
}

type MasterCodeGroup struct {
	Id                 int    `json:"id"`
	CodeGroup          string `json:"code_group"`
	Description        string `json:"description"`
	EnglishDescription string `json:"english_description"`
	Parent             *int   `json:"parent"`
	Sequence           int    `json:"sequence"`
	ModelMasterForm    `json:"-"`
}
