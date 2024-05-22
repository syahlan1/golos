package models

type OwnershipData struct {
	Id                int    `json:"id" gorm:"primaryKey"`
	Name              string `json:"name"`
	NoIdentity        string `json:"no_identity"`
	IdCardAddress     string `json:"id_card_identity"`
	City              string `json:"city"`
	ZipCode           string `json:"zip_code"`
	HomeOwnership     string `json:"home_ownership"`
	Remark            string `json:"remark"`
	CifManager        string `json:"cif_manager"`
	BirthDate         string `json:"birth_date"`
	LastEducation     string `json:"last_education"`
	NPWP              string `json:"npwp"`
	JobTitle          string `json:"job_title"`
	Experince         string `json:"experience"`
	OwnershipMarket   int    `json:"ownership_market"`
	CitizenshipStatus string `json:"citizenship_status"`
	Gender            string `json:"gender"`
	MaritalStatus     string `json:"marital_status"`
	NumberOfChildren  int    `json:"number_of_children"`
	StartDate         string `json:"start_date"`
	KeyPerson         string `json:"key_person"`
	Removed           string `json:"removed"`
	BusinessId        int    `json:"business_id"`
	Status            string `json:"status"`
}

type RelationWithBank struct {
	Id              int `json:"id" gorm:"primaryKey"`
	Giro            string
	Tabungan        string
	NoRekening      int `json:"no_rekening"`
	Debitur         string
	OwnershipDataId int
	Status          string
}

type DataRekeningDebitur struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	NoIdCard        string `json:"no_id_card"`
	NPWP            int    `json:"npwp"`
	KeyPerson       string `json:"key_person"`
	NoRekening      int    `json:"no_rekening"`
	Remark          string `json:"remark"`
	OwnershipDataId int
	Status          string
}
