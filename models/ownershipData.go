package models

type OwnershipData struct {
	Id                   int     `json:"id" gorm:"primaryKey"`
	GeneralInformationId int     `json:"general_information_id"`
	OwnershipType        int     `json:"ownership_type"`
	Name                 string  `json:"name"`
	NoIdentity           string  `json:"no_identity"`
	IdCardAddress        string  `json:"id_card_address"`
	City                 string  `json:"city"`
	ZipCode              string  `json:"zip_code"`
	HomeOwnership        string  `json:"home_ownership"`
	Remark               string  `json:"remark"`
	CifManager           string  `json:"cif_manager"`
	BirthDate            string  `json:"birth_date"`
	LastEducation        string  `json:"last_education"`
	NPWP                 string  `json:"npwp"`
	JobTitle             string  `json:"job_title"`
	Experince            string  `json:"experience"`
	OwnershipMarket      float64 `json:"ownership_market"`
	CitizenshipStatus    string  `json:"citizenship_status"`
	Gender               string  `json:"gender"`
	MaritalStatus        string  `json:"marital_status"`
	NumberOfChildren     int     `json:"number_of_children"`
	StartDate            string  `json:"start_date"`
	KeyPerson            bool    `json:"key_person"`
	Removed              bool    `json:"removed"`
	Status               string  `json:"status"`
}

type CreateOwnershipData struct {
	GeneralInformationId int     `json:"general_information_id"`
	OwnershipType        int     `json:"ownership_type"`
	Name                 string  `json:"name"`
	NoIdentity           string  `json:"no_identity"`
	IdCardAddress        string  `json:"id_card_address"`
	City                 string  `json:"city"`
	ZipCode              string  `json:"zip_code"`
	HomeOwnership        string  `json:"home_ownership"`
	Remark               string  `json:"remark"`
	CifManager           string  `json:"cif_manager"`
	BirthDate            string  `json:"birth_date"`
	LastEducation        string  `json:"last_education"`
	NPWP                 string  `json:"npwp"`
	JobTitle             string  `json:"job_title"`
	Experince            string  `json:"experience"`
	OwnershipMarket      float64 `json:"ownership_market"`
	CitizenshipStatus    string  `json:"citizenship_status"`
	Gender               string  `json:"gender"`
	MaritalStatus        string  `json:"marital_status"`
	NumberOfChildren     int     `json:"number_of_children"`
	StartDate            string  `json:"start_date"`
	KeyPerson            bool    `json:"key_person"`
	Removed              bool    `json:"removed"`
}

type CreateRelationWithBank struct {
	Name       string `json:"name"`
	NoIdCard   string `json:"no_id_card"`
	NPWP       int    `json:"npwp"`
	KeyPerson  string `json:"key_person"`
	Remark     string `json:"remark"`
	Giro       string `json:"giro"`
	Tabungan   string `json:"tabungan"`
	NoRekening int    `json:"no_rekening"`
	Debitur    string `json:"debitur"`
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
	Id                   int    `json:"id" gorm:"primaryKey"`
	GeneralInformationId int    `json:"general_information_id"`
	OwnershipDataId      int    `json:"ownership_data_id"`
	NoRekening           string `json:"no_rekening"`
	Remark               string `json:"remark"`
	Status               string `json:"status"`
}

type ShowRekeningDebitur struct {
	Id                   int    `json:"id" gorm:"primaryKey"`
	// GeneralInformationId int    `json:"general_information_id"`
	OwnershipDataId      int    `json:"ownership_data_id"`
	Name                 string `json:"name"`
	NoIdentity        string `json:"no_identity"`
	NPWP                 string `json:"npwp"`
	KeyPerson            bool   `json:"key_person"`
	Pemilik              string `json:"pemilik"`
	NoRekening           string `json:"no_rekening"`
	Remark               string `json:"remark"`
	Status               string `json:"status"`
}
