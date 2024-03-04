package models

type Users struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Roles    string `json:"roles"`
	Status   int    `json:"status"`
	Password []byte `json:"-"`
}
