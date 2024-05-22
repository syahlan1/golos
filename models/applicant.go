package models

type Applicant struct {
	Id                   int    `json:"id" gorm:"primaryKey"`
	Cif                  string `json:"cif"`
	TitleBeforeName      string `json:"title_before_name"`
	CustomerName         string `json:"customer_name"`
	TitleAfterName       string `json:"title_after_name"`
	NickName             string `json:"nickname"`
	HomeAddress          string `json:"home_address"`
	District             string `json:"district"`
	City                 string `json:"city"`
	ZipCode              string `json:"zip_code"`
	HomeStatus           string `json:"home_status"`
	StaySince            string `json:"stay_since"`
	NoTelp               string `json:"no_telp"`
	NoFax                string `json:"no_fax"`
	BirthPlace           string `json:"birth_place"`
	BirthDate            string `json:"birth_date"`
	MaritalStatus        string `json:"marital_status"`
	Gender               string `json:"gender"`
	Nationality          string `json:"nationality"`
	NumberOfChildren     int    `json:"number_of_children"`
	NoKartuKeluarga      string `json:"no_kartu_keluarga"`
	Education            string `json:"education"`
	JobPosition          string `json:"job_position"`
	BusinessSector       string `json:"business_sector"`
	EstablishDate        string `json:"establish_date"`
	NPWP                 string `json:"npwp"`
	GrossIncomePerMonth  int    `json:"gross_income_per_month"`
	NumberOfEmployees    int    `json:"number_of_employees"`
	MotherName           string `json:"mother_name"`
	NamaPelaporan        string `json:"nama_pelaporan"`
	NegaraDomisili       string `json:"negara_domisili"`
	NamaInstansi         string `json:"nama_instansi"`
	KodeInstansi         string `json:"kode_instansi"`
	NoPegawai            string `json:"no_pegawai"`
	Status               string `json:"status"`
	IdCard               int    `json:"id_card"`
	GeneralInformationId int    `json:"general_information_id"`
	DocumentId           int    `json:"document_id"`
	SpouseId             int    `json:"spouse_id"`
}


type CreateApplicant struct {
	Applicant          Applicant          `json:"applicant"`
	Spouse             SpouseData         `json:"spouse"`
	IdCard             IdCard             `json:"id_card"`
	Document           Document           `json:"document"`
	GeneralInformation GeneralInformation `json:"general_information"`
}


type ApplicantDetail struct {
	Applicant
	SpouseData         SpouseData         `json:"spouse"`
	IdCard             IdCard             `json:"id_card"`
	Document           Document           `json:"document"`
	GeneralInformation GeneralInformation `json:"general_information"`
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

type SektorEkonomi struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}

type HubunganNasabah struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}

type HubunganKeluarga struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}

type LokasiPabrik struct {
	Id   int    `json: "id" gorm:"primaryKey"`
	Name string `json: "name"`
}

type MaritalStatus struct {
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
