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
	AddressType      string `json:"address_type"`
}
