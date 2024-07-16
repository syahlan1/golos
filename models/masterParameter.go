package models

type MasterParameter struct {
	Id                 int    `json:"id"`
	IsEncrypted        int    `json:"is_encrypted"`
	ModuleId           int    `json:"module_id"`
	ModuleName         string `json:"module_name,omitempty" gorm:"-:migration"`
	Description        string `json:"description"`
	EnglishDescription string `json:"english_description"`
	ParamKey           string `json:"param_key"`
	ParamValue         string `json:"param_value"`
	ModelMasterForm    `json:"-"`
}

type CreateMasterParameter struct {
	Description        string `json:"description"`
	EnglishDescription string `json:"english_description"`
	ParamKey           string `json:"param_key"`
	ParamValue         string `json:"param_value"`
}

type ShowAllMasterParameter struct {
	ModuleId   int               `json:"module_id"`
	ModuleName string            `json:"module_name"`
	Parameter  []MasterParameter `json:"parameter" gorm:"-"`
}
