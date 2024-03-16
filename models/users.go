package models

type Users struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	IsLogin  int    `json:"is_login"`
	Password []byte `json:"-"`
}

type Roles struct {
	Id       int    `json:"id"`
	RoleName string `json:"role_name"`
}

type UserRole struct {
	UserId int   `json:"user_id"`
	RoleId int   `json:"role_id"`
	Role   Roles `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	User   Users `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
