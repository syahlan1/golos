package controllers

type Business struct {
	Cif                  string `json:"cif"`
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
	ListingBursaCode     string `json:"listing_bursa_code"`
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
}

type Applicant struct {
	TitleBeforeName  string `json:"title_before_name"`
	CustomerName     string `json:"customer_name"`
	TitleAfterName   string `json:"title_after_name"`
	NickName         string `json:"nickname"`
	HomeAddress      string `json:"home_address"`
	District         string `json:"district"`
	City             string `json:"city"`
	ZipCode          string `json:"zip_code"`
	HomeStatus       string `json:"home_status"`
	StaySince        string `json:"stay_since"`
	NoTelp           int    `json:"no_telp"`
	NoFax            int    `json:"no_fax"`
	BirthPlace       string `json:"birth_place"`
	BirthDate        string `json:"birth_date"`
	MaritalStatus    string `json:"marital_status"`
	Gender           string `json:"gender"`
	Nationality      string `json:"nationality"`
	NumberOfChildren int    `json:"number_of_children"`
	NoKartuKeluarga  int    `json:"no_kartu_keluarga"`
}
