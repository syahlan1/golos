package models

type IdCard struct {
	Id               int    `json:"id"`
	IdCardIssuedDate string `json:"id_card_issued_date"`
	IdCard           string `json:"id_card"`
	IdCardExpireDate string `json:"id_card_expire_date"`
	IdCardAddress    string `json:"id_card_address"`
	IdCardDistrict   string `json:"id_card_district"`
	IdCardCity       string `json:"id_card_city"`
	IdCardZipCode    string `json:"id_card_zip_code"`
	AddressTypeId    int    `json:"address_type_id"`
}

type ShowIdCard struct {
	Id               int    `json:"id"`
	IdCardIssuedDate string `json:"id_card_issued_date"`
	IdCard           string `json:"id_card"`
	IdCardExpireDate string `json:"id_card_expire_date"`
	IdCardAddress    string `json:"id_card_address"`
	IdCardDistrict   string `json:"id_card_district"`
	IdCardCity       string `json:"id_card_city"`
	IdCardZipCode    string `json:"id_card_zip_code"`
	AddressTypeId    int    `json:"address_type_id"`
	AddressType      string `json:"address_type"`
}

type AddressType struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}
