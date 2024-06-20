package models

type Business struct {
	Id                      int    `json:"id" gorm:"primaryKey"`
	ExistingNasabah         bool   `json:"existing_nasabah"`
	Cif                     string `json:"cif"`
	CompanyFirstNameId      int    `json:"company_first_name_id"`
	CompanyName             string `json:"company_name"`
	CompanyTypeId           int    `json:"company_type_id"`
	CustomerName            string `json:"customer_name"`
	EstablishDate           string `json:"establishment_date"`
	EstablishPlace          string `json:"establish_place"`
	CompanyAddress          string `json:"company_address"`
	District                string `json:"district"`
	City                    string `json:"city"`
	ZipCode                 string `json:"zip_code"`
	AddressTypeId           int    `json:"address_type_id"`
	ExternalRatingCompanyId int    `json:"external_rating_company_id"`
	RatingClassId           int    `json:"rating_class_id"`
	RatingDate              string `json:"rating_date"`
	ListingBursaCodeId      int    `json:"listing_bursa_code_id"`
	ListingBursaDate        string `json:"listing_bursa_date"`
	BusinessTypeId          int    `json:"business_type_id"`
	AktaPendirian           string `json:"akta_pendirian"`
	TglTerbit               string `json:"tgl_terbit"`
	AktaLastChange          string `json:"akta_last_change"`
	LastChangeDate          string `json:"last_change_date"`
	NotarisName             string `json:"notaris_name"`
	JumlahKaryawan          int    `json:"jumlah_karyawan"`
	NoTelp                  string `json:"no_telp"`
	NoFax                   string `json:"no_fax"`
	NPWP                    string `json:"npwp"`
	TDP                     string `json:"tdp"`
	TglPenerbitan           string `json:"tgl_penerbitan"`
	TglJatuhTempo           string `json:"tgl_jatuh_tempo"`
	ContactPerson           string `json:"contact_person"`
	ApproveStatus           string `json:"approve_status"`
	// Status                  string `json:"status"`
	DocumentId           int `json:"document_id"`
	GeneralInformationId int `json:"general_information_id"`
	SectorEconomyId      int `json:"sector_economy_id"`
	Model                `json:"-"`
}

type ShowBusiness struct {
	Id                      int    `json:"id" gorm:"primaryKey"`
	ExistingNasabah         bool   `json:"existing_nasabah"`
	Cif                     string `json:"cif"`
	CompanyFirstNameId      int    `json:"company_first_name_id"`
	CompanyName             string `json:"company_name"`
	CompanyTypeId           int    `json:"company_type_id"`
	CompanyType             string `json:"company_type"`
	CustomerName            string `json:"customer_name"`
	EstablishDate           string `json:"establishment_date"`
	EstablishPlace          string `json:"establish_place"`
	CompanyAddress          string `json:"company_address"`
	District                string `json:"district"`
	City                    string `json:"city"`
	ZipCode                 string `json:"zip_code"`
	AddressTypeId           int    `json:"address_type_id"`
	AddressType             string `json:"address_type"`
	ExternalRatingCompanyId int    `json:"external_rating_company_id"`
	ExternalRatingCompany   string `json:"external_rating_company"`
	RatingClassId           int    `json:"rating_class_id"`
	RatingClass             string `json:"rating_class"`
	RatingDate              string `json:"rating_date"`
	ListingBursaCodeId      int    `json:"listing_bursa_code_id"`
	ListingBursaCode        string `json:"listing_bursa_code"`
	ListingBursaDate        string `json:"listing_bursa_date"`
	BusinessTypeId          int    `json:"business_type_id"`
	BusinessType            string `json:"business_type"`
	AktaPendirian           string `json:"akta_pendirian"`
	TglTerbit               string `json:"tgl_terbit"`
	AktaLastChange          string `json:"akta_last_change"`
	LastChangeDate          string `json:"last_change_date"`
	NotarisName             string `json:"notaris_name"`
	JumlahKaryawan          int    `json:"jumlah_karyawan"`
	NoTelp                  string `json:"no_telp"`
	NoFax                   string `json:"no_fax"`
	NPWP                    string `json:"npwp"`
	TDP                     string `json:"tdp"`
	TglPenerbitan           string `json:"tgl_penerbitan"`
	TglJatuhTempo           string `json:"tgl_jatuh_tempo"`
	ContactPerson           string `json:"contact_person"`
	ApproveStatus           string `json:"approve_status"`
	DocumentId              int    `json:"-"`
	GeneralInformationId    int    `json:"-"`
	SectorEconomyId         int    `json:"-"`
}

type BusinessApplicant struct {
	Business  []Business  `json:"business"`
	Applicant []Applicant `json:"applicant"`
}

type CreateBusiness struct {
	Document           Document           `json:"document"`
	GeneralInformation GeneralInformation `json:"general_information"`
	Business           Business           `json:"business"`
	SectorEconomy      SectorEconomy      `json:"sector_economy"`
}

type BusinessDetail struct {
	ShowBusiness
	Document           Document               `json:"document"`
	GeneralInformation ShowGeneralInformation `json:"general_information"`
	SectorEconomy      ShowSectorEconomy      `json:"sector_economy"`
}
type CompanyFirstName struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type CompanyType struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type BusinessAddressType struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type ExternalRatingCompany struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type RatingClass struct {
	Id               int    `json:"id" gorm:"primaryKey"`
	ExternalRatingId int    `json:"external_rating_id"`
	Code             string `json:"code" gorm:"type:varchar(10)"`
	Name             string `json:"name" gorm:"type:varchar(250)"`
	Sibs             string `json:"sibs" gorm:"type:varchar(250)"`
	Model
}

type KodeBursa struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type BusinessType struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}
