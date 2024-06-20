package models

type Applicant struct {
	Id                  int    `json:"id" gorm:"primaryKey"`
	ExistingNasabah     bool   `json:"existing_nasabah"`
	Cif                 string `json:"cif"`
	TitleBeforeName     string `json:"title_before_name"`
	CustomerName        string `json:"customer_name"`
	TitleAfterName      string `json:"title_after_name"`
	NickName            string `json:"nickname"`
	HomeAddress         string `json:"home_address"`
	District            string `json:"district"`
	City                string `json:"city"`
	ZipCode             string `json:"zip_code"`
	HomeStatusId        int    `json:"home_status_id"`
	StaySince           string `json:"stay_since"`
	NoTelp              string `json:"no_telp"`
	NoFax               string `json:"no_fax"`
	BirthPlace          string `json:"birth_place"`
	BirthDate           string `json:"birth_date"`
	MaritalStatusId     int    `json:"marital_status_id"`
	GenderId            int    `json:"gender_id"`
	NationalityId       int    `json:"nationality_id"`
	NumberOfChildren    int    `json:"number_of_children"`
	NoKartuKeluarga     string `json:"no_kartu_keluarga"`
	EducationId         int    `json:"education_id"`
	JobPositionId       int    `json:"job_position_id"`
	BusinessSectorId    int    `json:"business_sector_id"`
	EstablishDate       string `json:"establish_date"`
	NPWP                string `json:"npwp"`
	GrossIncomePerMonth int    `json:"gross_income_per_month"`
	NumberOfEmployees   int    `json:"number_of_employees"`
	MotherName          string `json:"mother_name"`
	NamaPelaporan       string `json:"nama_pelaporan"`
	NegaraDomisiliId    int    `json:"negara_domisili_id"`
	NamaInstansi        string `json:"nama_instansi"`
	KodeInstansi        string `json:"kode_instansi"`
	NoPegawai           string `json:"no_pegawai"`
	// Status               string `json:"status"`
	IdCard               int `json:"id_card"`
	GeneralInformationId int `json:"general_information_id"`
	DocumentId           int `json:"document_id"`
	SpouseId             int `json:"spouse_id"`
	SectorEconomyId      int `json:"sector_economy_id"`
	Model                `json:"-"`
	HomeStatus           HomeStatus     `json:"-" gorm:"foreignKey:HomeStatusId"`
	MaritalStatus        MaritalStatus  `json:"-" gorm:"foreignKey:MaritalStatusId"`
	Gender               Gender         `json:"-" gorm:"foreignKey:GenderId"`
	Nationality          Nationality    `json:"-" gorm:"foreignKey:NationalityId"`
	Education            Education      `json:"-" gorm:"foreignKey:EducationId"`
	JobPosition          JobPosition    `json:"-" gorm:"foreignKey:JobPositionId"`
	BusinessSector       BusinessSector `json:"-" gorm:"foreignKey:BusinessSectorId"`
	NegaraDomisili       Negara         `json:"-" gorm:"foreignKey:NegaraDomisiliId"`
	SectorEconomy        SectorEconomy  `json:"-" gorm:"foreignKey:SectorEconomyId"`
	Spouse               SpouseData     `json:"-" gorm:"foreignKey:SpouseId"`
	IdCards              IdCard             `json:"-" gorm:"foreignKey:IdCard;references:Id"`
	GeneralInformation GeneralInformation `json:"-" gorm:"foreignKey:GeneralInformationId"`
	Document           Document           `json:"-" gorm:"foreignKey:DocumentId"`
}

type ShowApplicant struct {
	Id                  int    `json:"id" gorm:"primaryKey"`
	ExistingNasabah     bool   `json:"existing_nasabah"`
	Cif                 string `json:"cif"`
	TitleBeforeName     string `json:"title_before_name"`
	CustomerName        string `json:"customer_name"`
	TitleAfterName      string `json:"title_after_name"`
	NickName            string `json:"nickname"`
	HomeAddress         string `json:"home_address"`
	District            string `json:"district"`
	City                string `json:"city"`
	ZipCode             string `json:"zip_code"`
	HomeStatusId        int    `json:"home_status_id"`
	HomeStatus          string `json:"home_status"`
	StaySince           string `json:"stay_since"`
	NoTelp              string `json:"no_telp"`
	NoFax               string `json:"no_fax"`
	BirthPlace          string `json:"birth_place"`
	BirthDate           string `json:"birth_date"`
	MaritalStatusId     int    `json:"marital_status_id"`
	MaritalStatus       string `json:"marital_status"`
	GenderId            int    `json:"gender_id"`
	Gender              string `json:"gender"`
	NationalityId       int    `json:"nationality_id"`
	Nationality         string `json:"nationality"`
	NumberOfChildren    int    `json:"number_of_children"`
	NoKartuKeluarga     string `json:"no_kartu_keluarga"`
	EducationId         int    `json:"education_id"`
	Education           string `json:"education"`
	JobPositionId       int    `json:"job_position_id"`
	JobPosition         string `json:"job_position"`
	BusinessSectorId    int    `json:"business_sector_id"`
	BusinessSector      string `json:"business_sector"`
	EstablishDate       string `json:"establish_date"`
	NPWP                string `json:"npwp"`
	GrossIncomePerMonth int    `json:"gross_income_per_month"`
	NumberOfEmployees   int    `json:"number_of_employees"`
	MotherName          string `json:"mother_name"`
	NamaPelaporan       string `json:"nama_pelaporan"`
	NegaraDomisiliId    int    `json:"negara_domisili_id"`
	NegaraDomisili      string `json:"negara_domisili"`
	NamaInstansi        string `json:"nama_instansi"`
	KodeInstansi        string `json:"kode_instansi"`
	NoPegawai           string `json:"no_pegawai"`
	// Status               string `json:"status"`
	IdCard               int `json:"id_card"`
	GeneralInformationId int `json:"-"`
	DocumentId           int `json:"-"`
	SpouseId             int `json:"-"`
	SectorEconomyId      int `json:"-"`
}

type CreateApplicant struct {
	Applicant          Applicant          `json:"applicant"`
	Spouse             SpouseData         `json:"spouse"`
	SectorEconomy      SectorEconomy      `json:"sector_economy"`
	IdCard             IdCard             `json:"id_card"`
	Document           Document           `json:"document"`
	GeneralInformation GeneralInformation `json:"general_information"`
}

type ApplicantDetail struct {
	ShowApplicant
	SpouseData         SpouseData             `json:"spouse"`
	SectorEconomy      ShowSectorEconomy      `json:"sector_economy"`
	IdCard             ShowIdCard             `json:"id_card"`
	Document           Document               `json:"document"`
	GeneralInformation ShowGeneralInformation `json:"general_information"`
}

type HomeStatus struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Sibs string `json:"sibs" gorm:"type:varchar(250)"`
	Model
}

type ApplicantAddressType struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Model
}

type Education struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type JobPosition struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type BusinessSector struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type KodeInstansi struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type Negara struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type SektorEkonomi struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Model
}

type MaritalStatus struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type Gender struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type Nationality struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type ZipCode struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	ZipCode     string `json:"zip_code"`
	Subdistrict string `json:"subdistrict"`
	District    string `json:"district"`
	City        string `json:"city"`
	Province    string `json:"province"`
	Model
}
