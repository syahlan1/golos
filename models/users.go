package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id                 uint      `json:"id" gorm:"primaryKey"`
	Username           string    `json:"username" gorm:"unique"`
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	Email              string    `json:"email" gorm:"unique"`
	IsActive           int       `json:"is_active"`
	Status             string    `json:"status"`
	IsLogin            int       `json:"is_login"`
	Password           []byte    `json:"-"`
	PasswordHistory    []string  `json:"password_history" gorm:"type:json"` // Menyimpan hash dari kata sandi sebelumnya
	LastPasswordChange time.Time `json:"last_password_change"`
	FailedAttempts     int       `json:"failed_attempts" gorm:"default:0"`
	LastLogin          time.Time `json:"last_login"`
	RoleId             uint
	Role               Roles
	gorm.Model
}

type Register struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsActive  int    `json:"is_active"`
	Password  string `json:"password"`
}

type Login struct {
	Username string `json:"username"`
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
