package models

type Validation struct {
	Id          int    `json:"id" gorm:"autoIncrement"`
	Name        string `json:"name" gorm:"unique"`
	Validation  string `json:"validation"`
	Description string `json:"description"`
}
