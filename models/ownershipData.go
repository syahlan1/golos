package models

import "github.com/syahlan1/golos/utils/formatTime"

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
	OwnershipMarket      float64 `json:"ownership_market" gorm:"type:float"`
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
	OwnershipMarket      float64 `json:"ownership_market" gorm:"type:float"`
	CitizenshipStatus    string  `json:"citizenship_status"`
	Gender               string  `json:"gender"`
	MaritalStatus        string  `json:"marital_status"`
	NumberOfChildren     int     `json:"number_of_children"`
	StartDate            string  `json:"start_date"`
	KeyPerson            bool    `json:"key_person"`
	Removed              bool    `json:"removed"`
}
type RelationWithBank struct {
	Id                   int                 `json:"id" gorm:"primaryKey;autoIncrement"`
	GeneralInformationId int                 `json:"general_information_id"`
	Giro                 formatTime.WrapDate `json:"giro" gorm:"type:date"`
	Tabungan             formatTime.WrapDate `json:"tabungan" gorm:"type:date"`
	NoRekening           string              `json:"no_rekening"`
	Debitur              formatTime.WrapDate `json:"debitur" gorm:"type:date"`
	Status               string              `json:"status"`
}

type CustomerLoanInfo struct {
	Id                   int    `json:"id" gorm:"primaryKey"`
	GeneralInformationId int    `json:"general_information_id"`
	AAStatus             int    `json:"aa_status" gorm:"-"`
	AANo                 string `json:"aa_no"`
	FacilityId           int    `json:"facility_id"`
	FacilitySequence     string `json:"facility_sequence"`
	ChannelingFacilty    bool   `json:"channeling_facility"`
	ProductId            int    `json:"product_id"`
	NoRekening           string `json:"no_rekening"`
	Status               string `json:"status"`
}

type CustomerAA struct {
	Id                   int    `json:"id"`
	GeneralInformationId int    `json:"general_information_id"`
	AANo                 string `json:"aa_no"`
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
	Id int `json:"id" gorm:"primaryKey"`
	// GeneralInformationId int    `json:"general_information_id"`
	OwnershipDataId int    `json:"ownership_data_id"`
	Name            string `json:"name"`
	NoIdentity      string `json:"no_identity"`
	NPWP            string `json:"npwp"`
	KeyPerson       bool   `json:"key_person"`
	Pemilik         string `json:"pemilik"`
	NoRekening      string `json:"no_rekening"`
	Remark          string `json:"remark"`
	Status          string `json:"status"`
}
