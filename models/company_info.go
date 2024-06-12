package models

type CompanyInfo struct {
	Id          int   `json:"id" gorm:"primary_key"`
	InfoType int    `json:"info_type"`
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name"`
	LastName    string `json:"last_name"`
	IdCardNumber string `json:"id_card_number"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	Address3    string `json:"address3"`
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
	HomeOwnershipId int `json:"home_ownership_id"`
	HomeOwnership string `json:"home_ownership" gorm:"-:migration"`
	Remark       string `json:"remark"`
	Cif string `json:"cif"`
	BirthDate   string `json:"birth_date"`
	LastEducation string `json:"last_education"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Position    string `json:"position"`
}