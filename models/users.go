package models

import (
	"time"
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
	ModelMasterForm    `json:"-"`
}

type UserPermission struct {
	Id       uint      `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	RoleId   uint      `json:"-"`
	Role     ShowRoles `json:"role"`
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
