package models

type GeneralInformation struct {
	Id                    int    `json:"id"`
	BankName              string `json:"bank_name"`
	KCP                   string `json:"kcp"`
	SubProgramId          int    `json:"sub_program_id"`
	Analisis              string `json:"analisis"`
	CabangPencairanId     int    `json:"cabang_pencairan_id"`
	CabangAdminId         int    `json:"cabang_admin_id"`
	TglAplikasi           string `json:"tgl_aplikasi"`
	TglPenerusan          string `json:"tgl_penerusan"`
	SegmenId              int    `json:"segmen_id"`
	NoAplikasi            string `json:"no_aplikasi"`
	NoReferensi           string `json:"no_referensi"`
	MarketInterestRate    int    `json:"market_interest_rate"`
	RequestedInterestRate int    `json:"requested_interest_rate"`
	DocumentFile          string `json:"document_file"`
	Status                string `json:"status"`
	Model                 `json:"-"`
	SubProgram            Program `json:"-" gorm:"foreignKey:SubProgramId"`
	CabangPencairan       Cabang  `json:"-" gorm:"foreignKey:CabangPencairanId"`
	CabangAdmin           Cabang  `json:"-" gorm:"foreignKey:CabangAdminId"`
	Segmen                Segment `json:"-" gorm:"foreignKey:SegmenId"`
}

type ShowGeneralInformation struct {
	Id                    int    `json:"id"`
	BankName              string `json:"bank_name"`
	KCP                   string `json:"kcp"`
	SubProgramId          int    `json:"sub_program_id"`
	SubProgram            string `json:"sub_program"`
	Analisis              string `json:"analisis"`
	CabangPencairanId     int    `json:"cabang_pencairan_id"`
	CabangPencairan       string `json:"cabang_pencairan"`
	CabangAdminId         int    `json:"cabang_admin_id"`
	CabangAdmin           string `json:"cabang_admin"`
	TglAplikasi           string `json:"tgl_aplikasi"`
	TglPenerusan          string `json:"tgl_penerusan"`
	SegmenId              int    `json:"segmen_id"`
	Segmen                string `json:"segmen"`
	NoAplikasi            string `json:"no_aplikasi"`
	NoReferensi           string `json:"no_referensi"`
	MarketInterestRate    int    `json:"market_interest_rate"`
	RequestedInterestRate int    `json:"requested_interest_rate"`
}

type Cabang struct {
	Id              int    `json:"id" gorm:"primaryKey"`
	BranchCode      string `json:"branch_code"`
	Name            string `json:"name"`
	CBCCode         string `json:"cbc_code"`
	Address         string `json:"address"`
	Cabang          bool   `json:"cabang"`
	CabangPencairan bool   `json:"cabang_pencairan"`
	Model           `json:"-"`
}

type Program struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Model `json:"-"`
}

type Segment struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Model `json:"-"`
}
