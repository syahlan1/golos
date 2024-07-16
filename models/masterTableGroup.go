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
	Order              int    `json:"order"`
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
	IsMandatory     bool   `json:"is_mandatory"`
	ModelMasterForm `json:"-"`
	Group           MasterTableGroup `json:"-" gorm:"foreignKey:GroupId"`
	MasterTable     MasterTable      `json:"-" gorm:"foreignKey:TableId"`
}

type TableGroupItemStatus struct {
	Id              int     `json:"id"`
	GroupId         int     `json:"group_id"`
	Status          string  `json:"status"`
	Reason          *string `json:"reason"`
	Username        string  `json:"username"`
	ModelMasterForm `json:"-"`
	Group           MasterTableGroup `json:"-" gorm:"foreignKey:GroupId"`
}

type ShowApprovalTableGroup struct {
	Submitted []map[string]interface{} `json:"submitted"`
	Approved  []map[string]interface{} `json:"approved"`
	Rejected  []map[string]interface{} `json:"rejected"`
}

type ShowDetailApprovalTableGroup struct {
	Id      int                            `json:"id"`
	Status  string                         `json:"status"`
	Reason  *string                        `json:"reason"`
	GroupId int                            `json:"group_id"`
	Data    []DataDetailApprovalTableGroup `json:"data" gorm:"-"`
}

type DataDetailApprovalTableGroup struct {
	TableName string                   `json:"table_name"`
	Type      int                      `json:"type"`
	SchemaId  string                   `json:"-"`
	TableId   string                   `json:"-"`
	Data      []map[string]interface{} `json:"data" gorm:"-"`
}

type ShowDetailApprovalTableGroupParent struct {
	Id                 int     `json:"id"`
	Status             string  `json:"status,omitempty"`
	Description        string  `json:"description"`
	EnglishDescription string  `json:"english_description"`
	Reason             *string `json:"reason,omitempty"`
	GroupId            int     `json:"group_id,omitempty"`
	ParentType         string  `json:"-"`
	// ParentId           *int                                 `json:"-"`
	Data  []DataDetailApprovalTableGroup       `json:"data,omitempty" gorm:"-"`
	Child []ShowDetailApprovalTableGroupParent `json:"child,omitempty" gorm:"-"`
}

type FormMasterTableGroup struct {
	Id                 int                   `json:"id"`
	Type               string                `json:"type"`
	Description        string                `json:"description"`
	EnglishDescription string                `json:"english_description"`
	CanSubmit          bool                  `json:"can_submit"`
	Form               []FormMasterTableItem `json:"form" gorm:"-"`
}

type FormMasterTableItem struct {
	Id          int        `json:"id"`
	IsMandatory bool       `json:"-"`
	TableId     int        `json:"table_id"`
	TableName   string     `json:"table_name"`
	DataId      []int      `json:"data_id" gorm:"-"`
	Type        int        `json:"type"`
	Sequence    int        `json:"sequence"`
	FormList    []FormList `json:"form_list"  gorm:"-"`
}

type FormMasterTableGroupParent struct {
	Id                 int                          `json:"id"`
	Type               string                       `json:"type"`
	ParentType         string                       `json:"-"`
	Description        string                       `json:"description"`
	EnglishDescription string                       `json:"english_description"`
	CanSubmit          bool                         `json:"can_submit"`
	Form               []FormMasterTableItem        `json:"form,omitempty" gorm:"-"`
	ParentId           *int                         `json:"parent_id,omitempty"`
	Child              []FormMasterTableGroupParent `json:"child,omitempty" gorm:"-"`
}
