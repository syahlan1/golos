package models

type Menu struct {
	Id              int     `json:"id" gorm:"primaryKey"`
	Command         *string `json:"command"`
	Icon            string  `json:"icon"`
	Label           string  `json:"label"`
	Order           int     `json:"order"`
	ParentId        *int    `json:"parent_id"`
	Type            string  `json:"type"`
	MenuCode        *string `json:"menu_code"`
	ModelMasterForm `json:"-"`
}

// type ShowMenu struct {
// 	Id    int    `json:"id"`
// 	Icon  string `json:"icon"`
// 	Label string `json:"label"`
// 	Order int    `json:"order"`
// 	Type  string `json:"type"`
// 	Child []Menu `json:"child" gorm:"-"`
// }

type ShowMenu struct {
	Id       int        `json:"id"`
	ParentId int        `json:"parent_id,omitempty"`
	Type     string     `json:"-"`
	Icon     string     `json:"icon"`
	Title    string     `json:"title"`
	Path     string     `json:"path,omitempty"`
	Subnav   []ShowMenu `json:"subnav,omitempty" gorm:"-"`
}

type ShowParentMenuPermission struct {
	Id    int                  `json:"id"`
	Icon  string               `json:"icon"`
	Label string               `json:"label"`
	Order int                  `json:"order"`
	Type  string               `json:"type"`
	Child []ShowMenuPermission `json:"child" gorm:"-"`
}

type ShowMenuPermission struct {
	Id       int     `json:"id" gorm:"primaryKey"`
	Command  *string `json:"command"`
	Icon     string  `json:"icon"`
	Label    string  `json:"label"`
	Type     string  `json:"type"`
	MenuCode *string `json:"menu_code"`
	Read     bool    `json:"read"`
	Delete   bool    `json:"delete"`
	Update   bool    `json:"update"`
	Download bool    `json:"download"`
	Write    bool    `json:"write"`
}

type RoleMenu struct {
	Id              int  `json:"id" gorm:"primaryKey"`
	RoleId          int  `json:"role_id"`
	Read            bool `json:"read"`
	Delete          bool `json:"delete"`
	Update          bool `json:"update"`
	Download        bool `json:"download"`
	Write           bool `json:"write"`
	MenuId          int  `json:"menu_id"`
	ModelMasterForm `json:"-"`
	Role            Roles `json:"-" gorm:"foreignKey:RoleId"`
	Menu            Menu  `json:"-" gorm:"foreignKey:MenuId"`
}

// type CreateRoleMenu struct {
// 	RoleId   int        `json:"role_id"`
// 	RoleMenu []RoleMenu `json:"role_menu"`
// }

type ShowRoleMenu struct {
	Id       int    `json:"id"`
	MenuId   int    `json:"menu_id"`
	Menu     string `json:"menu"`
	Read     bool   `json:"read"`
	Delete   bool   `json:"delete"`
	Update   bool   `json:"update"`
	Download bool   `json:"download"`
	Write    bool   `json:"write"`
}
