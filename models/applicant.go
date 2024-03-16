package models

type Applicant struct {
	Id                  int    `json:"id" gorm:"primaryKey"`
	TitleBeforeName     string `json:"title_before_name"`
	CustomerName        string `json:"customer_name"`
	TitleAfterName      string `json:"title_after_name"`
	NickName            string `json:"nickname"`
	HomeAddress         string `json:"home_address"`
	District            string `json:"district"`
	City                string `json:"city"`
	ZipCode             int    `json:"zip_code"`
	HomeStatus          string `json:"home_status"`
	StaySince           string `json:"stay_since"`
	NoTelp              int    `json:"no_telp"`
	NoFax               int    `json:"no_fax"`
	BirthPlace          string `json:"birth_place"`
	BirthDate           string `json:"birth_date"`
	MaritalStatus       string `json:"marital_status"`
	Gender              string `json:"gender"`
	Nationality         string `json:"nationality"`
	NumberOfChildren    int    `json:"number_of_children"`
	NoKartuKeluarga     int    `json:"no_kartu_keluarga"`
	SpouseName          string `json:"spouse_name"`
	SpouseIdCard        int    `json:"spouse_id_card"`
	SpouseAddress       string `json:"spouse_address"`
	IdCardIssuedDate    string `json:"id_card_issued_date"`
	IdCard              int    `json:"id_card"`
	IdCardExpireDate    string `json:"id_card_expire_date"`
	IdCardAddress       string `json:"id_card_address"`
	IdCardDistrict      string `json:"id_card_district"`
	IdCardCity          string `json:"id_card_city"`
	IdCardZipCode       int    `json:"id_card_zip_code"`
	AddressType         string `json:"address_type"`
	Education           string `json:"education"`
	JobPosition         string `json:"job_position"`
	BusinessSector      string `json:"business_sector"`
	EstablishDate       string `json:"establish_date"`
	NPWP                int    `json:"npwp"`
	GrossIncomePerMonth int    `json:"gross_income_per_month"`
	NumberOfEmployees   int    `json:"number_of_employees"`
	MotherName          string `json:"mother_name"`
	NamaPelaporan       string `json:"nama_pelaporan"`
	NegaraDomisili      string `json:"negara_domisili"`
	NamaInstansi        string `json:"nama_instansi"`
	KodeInstansi        string `json:"kode_instansi"`
	NoPegawai           int    `json:"no_pegawai"`
	ApproveStatus       int    `json:"approve_status"`
}

type HomeStatus struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}

type ApplicantAddressType struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}

type Education struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}

type JobPosition struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}

type BusinessSector struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}

type KodeInstansi struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}

type Negara struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}

type ZipCode struct {
	Id          int    `json: "id" gorm:"primaryKey"`
	ZipCode     string `json: "zip_code"`
	Subdistrict string `json:"subdistrict"`
	District    string `json: "district"`
	City        string `json: "city"`
	Province    string `json: "province"`
}
