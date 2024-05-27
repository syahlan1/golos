package models

import "gorm.io/gorm"

type Users struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	IsLogin  int    `json:"is_login"`
	Password []byte `json:"-"`
	RoleId   uint
	Role     Roles
	gorm.Model
}

type Register struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Roles struct {
	Id          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"unique"`
	Description string       `json:"description"`
	CreatedBy   string       `json:"created_by"`
	UpdatedBy   string       `json:"updated_by"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
	Users       []Users      `gorm:"foreignKey:RoleId"`
	gorm.Model
}

type CreateRole struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

type Permission struct {
	Id   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique"`
	gorm.Model
}

type RolePermission struct {
	RolesId      uint `json:"roles_id"`
	PermissionId uint `json:"permission_id"`
}
