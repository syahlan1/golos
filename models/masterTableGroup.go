package models

type MasterTableGroup struct {
	Id                 int    `json:"id"`
	Description        string `json:"description"`
	EnglishDescription string `json:"english_description"`
	GroupName          string `json:"group_name"`
	ModuleId           int    `json:"module_id"`
	TableName          string `json:"table_name"`
	Type               string `json:"type"`
	MenuIcon           string `json:"menu_icon"`
	ModelMasterForm    `json:"-"`
	MasterModule       MasterModule `json:"-" gorm:"foreignKey:ModuleId"`
}

type MasterTableItem struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	FieldKey        string `json:"field_key"`
	Sequence        int    `json:"sequence"`
	GroupId         int    `json:"group_id"`
	TableId         int    `json:"table_id"`
	Type            int    `json:"type"`
	IsMaster        bool   `json:"is_master"`
	ModelMasterForm `json:"-"`
	Group           MasterTableGroup `json:"-" gorm:"foreignKey:GroupId"`
	MasterTable     MasterTable      `json:"-" gorm:"foreignKey:TableId"`
}
