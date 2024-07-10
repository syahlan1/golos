package models

type MasterTableGroup struct {
	Id                 int    `json:"id"`
	Description        string `json:"description"`
	EnglishDescription string `json:"english_description"`
	GroupName          string `json:"group_name"`
	ModuleId           int    `json:"module_id"`
	TableName          string `json:"table_name"`
	Type               string `json:"type"`
	ParentType         string `json:"parent_type"`
	ParentId           *int   `json:"parent_id"`
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

type TableGroupItemStatus struct {
	Id       int              `json:"id"`
	GroupId  int              `json:"group_id"`
	Status   string           `json:"status"`
	Reason   *string          `json:"reason"`
	Username string           `json:"username"`
	Group    MasterTableGroup `json:"-" gorm:"foreignKey:GroupId"`
}

type FormMasterTableGroup struct {
	Id   int                   `json:"id"`
	Type string                `json:"type"`
	Form []FormMasterTableItem `json:"form"`
}

type FormMasterTableItem struct {
	Id        int        `json:"id"`
	TableId   int        `json:"table_id"`
	TableName string     `json:"table_name"`
	Type      int        `json:"type"`
	Sequence  int        `json:"sequence"`
	FormList  []FormList `json:"form_list"  gorm:"-"`
}
