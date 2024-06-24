package models

type Roles struct {
	Id              uint         `gorm:"primaryKey" json:"id"`
	Name            string       `gorm:"unique"`
	Description     string       `json:"description"`
	Permissions     []Permission `gorm:"many2many:role_permissions;"`
	Users           []Users      `json:"-" gorm:"foreignKey:RoleId"`
	ModelMasterForm `json:"-"`
}

type ShowRoles struct {
	Id          uint                       `json:"id"`
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Module      []MasterModule             `json:"module" gorm:"-"`
	Menu        []ShowParentMenuPermission `json:"menu" gorm:"-"`
}

type CreateRole struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

type Permission struct {
	Id              uint   `gorm:"primaryKey" json:"id"`
	Name            string `gorm:"unique"`
	ModelMasterForm `json:"-"`
}

type RolePermission struct {
	RolesId      uint `json:"roles_id"`
	PermissionId uint `json:"permission_id"`
}

type RoleModules struct {
	Id              uint `gorm:"primaryKey" json:"id"`
	RolesId         uint `json:"roles_id"`
	ModuleId        int  `json:"module_id"`
	ModelMasterForm `json:"-"`
	Roles           Roles        `json:"-" gorm:"foreignKey:RolesId"`
	Module          MasterModule `json:"-" gorm:"foreignKey:ModuleId"`
}

type RoleTables struct {
	Id              uint `gorm:"primaryKey" json:"id"`
	RoleModulesId   uint `json:"role_modules_id"`
	TableId         int  `json:"table_id"`
	Read            bool `json:"read"`
	Delete          bool `json:"delete"`
	Update          bool `json:"update"`
	Download        bool `json:"download"`
	Write           bool `json:"write"`
	ModelMasterForm `json:"-"`
	RoleModules     RoleModules `json:"-" gorm:"foreignKey:RoleModulesId"`
	Table           MasterTable `json:"-" gorm:"foreignKey:TableId"`
}

type CreateRoleModules struct {
	RolesId  uint  `json:"roles_id"`
	ModuleId []int `json:"module_id"`
}

type ShowRoleModules struct {
	Id       uint   `json:"id"`
	ModuleId int    `json:"module_id"`
	Module   string `json:"module"`
}
