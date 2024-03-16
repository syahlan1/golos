package models

type Business struct {
	Id                   int    `json:"id" gorm:"primaryKey"`
	Cif                  int    `json:"cif"`
	CompanyFirstName     string `json:"company_first_name"`
	CompanyName          string `json:"company_name"`
	CompanyType          string `json:"company_type"`
	EstablishDate        string `json:"establishment_date"`
	EstablishPlace       string `json:"establish_place"`
	CompanyAddress       string `json:"company_address"`
	District             string `json:"district"`
	City                 string `json:"city"`
	ZipCode              int    `json:"zip_code"`
	AddressType          string `json:"address_type"`
	EternalRatingCompany string `json:"eternal_rating_company"`
	RatingClass          string `json:"rating_class"`
	RatingDate           string `json:"rating_date"`
	ListingBursaCode     int    `json:"listing_bursa_code"`
	ListingBursaDate     string `json:"listing_bursa_date"`
	BusinessType         string `json:"business_type"`
	AktaPendirian        string `json:"akta_pendirian"`
	TglTerbit            string `json:"tgl_terbit"`
	AktaLastChange       string `json:"akta_last_change"`
	LastChangeDate       string `json:"last_change_date"`
	NotarisName          string `json:"notaris_name"`
	JumlahKaryawan       int    `json:"jumlah_karyawan"`
	NoTelp               int    `json:"no_telp"`
	NoFax                int    `json:"no_fax"`
	NPWP                 int    `json:"npwp"`
	TDP                  string `json:"tdp"`
	TglPenerbitan        string `json:"tgl_penerbitan"`
	TglJatuhTempo        string `json:"tgl_jatuh_tempo"`
	ContactPerson        string `json:"contact_person"`
	ApproveStatus        int    `json:"approve_status"`
}

type CompanyFirstName struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}

type CompanyType struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}

type BusinessAddressType struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}

type EternalRatingCompany struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}

type RatingClass struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}

type KodeBursa struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}

type BusinessType struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}
