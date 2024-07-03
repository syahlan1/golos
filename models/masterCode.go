package models

type MasterCode struct {
	Id int `json:"id"`
	// Authoriser         string          `json:"authoriser"`
	// AuthorizeDate      time.Time       `json:"authorize_time"`
	// CreatedBy          string          `json:"created_by"`
	// CreatedDate        time.Time       `json:"created_date"`
	// Status             string          `json:"status"`
	// UpdatedBy          string          `json:"updated_by"`
	// UpdatedDate        time.Time       `json:"updated_date"`
	Code               string `json:"code"`
	CodeGroupId        int    `json:"code_group_id"`
	Description        string `json:"description"`
	EnglishDescription string `json:"english_description"`
	Sequence           int    `json:"sequence"`
	CodeGroup          string `json:"code_group"`
	ModelMasterForm
	MasterCodeGroup MasterCodeGroup `json:"-" gorm:"foreignKey:CodeGroupId"`
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
	Id int `json:"id"`
	// Authoriser         string    `json:"authoriser"`
	// AuthorizeDate      time.Time `json:"authorize_date"`
	// CreatedBy          string    `json:"created_by"`
	// CreatedDate        time.Time `json:"created_date"`
	// Status             string    `json:"status"`
	// UpdatedBy          string    `json:"updated_by"`
	// UpdatedDate        time.Time `json:"updated_date"`
	CodeGroup          string `json:"code_group"`
	Description        string `json:"description"`
	EnglishDescription string `json:"english_description"`
	ModelMasterForm
}
