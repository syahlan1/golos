package models

type CreditTerms struct {
	Id                   int             `json:"id" gorm:"primaryKey"` 
	GeneralInformationId int             `json:"general_information_id"`
	NewApplication       int             `json:"new_application"` //permohonan baru
	ChannelingCompany    int             `json:"channeling_company"` //perusahaan channeling
	TakeoverBank         bool            `json:"takeover_bank"` //nasabah ambil alih bank lain
	AccountAccomodation  int             `json:"account_accommodation"` //no. akomodasi rekening
	NCLProduct           int             `json:"ncl_product"` //produk ncl
	Facility             int             `json:"facility"` //fasilitas
	AccountNumber        string          `json:"account_number"` //no. rekening
	ApplicationType      string          `json:"application_type"` //jenis permohonan
	Project              int             `json:"project"` //proyek
	Status               string          `json:"status"`
	LoanNew              *LoanNew        `json:"loan_new,omitempty" gorm:"-:all"`
	LoanRenewal          *LoanRenewal    `json:"loan_renewal,omitempty" gorm:"-:all"`
	LoanWithdrawal       *LoanWithdrawal `json:"loan_withdrawal,omitempty" gorm:"-:all"`
	LoanPostFin          *LoanPostFin    `json:"loan_post_fin,omitempty" gorm:"-:all"`
}

type Loan struct {
	Id               int         `json:"id" gorm:"primaryKey"`
	CreditId         int         `json:"-"`
	SubmissionTypeId int         `json:"submission_type_id"`
	SubmissionType   string      `json:"submission_type" gorm:"-:migration"`
	
}

type CheckLoan struct {
	LoanNew     bool
	LoanRenewal bool
	// LoanWithdrawal bool
	// LoanPostFin bool
}

// type CreateCreditTerms struct {
// 	*CreditTerms
// 	LoanInformation *LoanInformation `json:"loan_new"`
// 	Collateral       *Collateral       `json:"collateral"`
// }

type LoanNew struct {
	Id               int         `json:"id" gorm:"primaryKey"`
	CreditId         int         `json:"-"`
	SubmissionTypeId int         `json:"submission_type_id"`
	SubmissionType   string      `json:"submission_type" gorm:"-:migration"`
	CreditTypeId     int         `json:"credit_type_id"`
	CreditType       string      `json:"credit_type" gorm:"-:migration"`
	Limit            int         `json:"limit"`
	ExchangeRate     int         `json:"exchange_rate"`
	LimitRp          int         `json:"limit_rp"`
	TimePeriod       int         `json:"time_period"`
	PeriodType       string      `json:"period_type"`
	PurposeId        int         `json:"purpose_id"`
	Purpose          string      `json:"purpose" gorm:"-:migration"`
	Description      string      `json:"description" gorm:"type:text"`
	CollateralStatus bool        `json:"collateral_status"`
	Collateral       *Collateral `json:"collateral" gorm:"-:all"`
	DeptorTransfer   bool        `json:"depositor_transfer"`
	OldCifNo         *string     `json:"old_cif_no" gorm:"default:null"`
	OldAccountNo     *string     `json:"old_account_no" gorm:"default:null"`
	Status           string      `json:"status"`
}

func (LoanNew) TableName() string {
	return "loan_new"
}

type LoanRenewal struct {
	Id                  int    `json:"id" gorm:"primaryKey"`
	CreditId            int    `json:"-"`
	SubmissionTypeId    int    `json:"submission_type_id"`
	SubmissionType      string `json:"submission_type" gorm:"-:migration"`
	AccountAccomodation int    `json:"account_accommodation"`
	AccountNumber       string `json:"account_number"`
	FacilityNumber      string `json:"facility_number"`
	CreditType          string `json:"credit_type"`
	Limit               int    `json:"limit"`
	TimePeriod          int    `json:"time_period"`
	PeriodType          string `json:"period_type"`
	PurposeId           int    `json:"purpose_id"`
	Purpose             string `json:"purpose" gorm:"-:migration"`
	TimePeriodRequest   int    `json:"time_period_request"`
	LimitRequest        int    `json:"limit_request"`
	ExchangeRate        int    `json:"exchange_rate"`
	LimitRp             int    `json:"limit_rp"`
	Description         string `json:"description" gorm:"type:text"`
	Status              string `json:"status"`
}

type LoanWithdrawal struct {
}

type LoanPostFin struct {
}

type Collateral struct {
	Id                         int    `json:"id" gorm:"primaryKey"`
	LoanId                     int    `json:"-"`
	CollateralTypeId           int    `json:"collateral_type_id"`
	CollateralType             string `json:"collateral_type" gorm:"-:migration"`
	Description                string `json:"description"`
	IdCoreCollateral           int    `json:"id_core_collateral"`
	ProofOfOwnershipId         int    `json:"proof_of_ownership_id"`
	ProofOfOwnership           string `json:"proof_of_ownership" gorm:"-:migration"`
	FormOfBindingId            int    `json:"form_of_binding_id"`
	FormOfBinding              string `json:"form_of_binding" gorm:"-:migration"`
	CollateralClassificationId int    `json:"collateral_classification_id"`
	CollateralClassification   string `json:"collateral_classification" gorm:"-:migration"`
	CurrencyId                 int    `json:"currency_id"`
	Currency                   string `json:"currency" gorm:"-:migration"`
	ExchangeRate               int    `json:"exchange_rate"`
	BankValue                  int    `json:"bank_value"`
	MarketValue                int    `json:"market_value"`
	InsuranceValue             int    `json:"insurance_value"`
	BindingValue               int    `json:"binding_value"`
	PPADeductionValue          int    `json:"ppa_deduction_value"`
	LiquidationValue           int    `json:"liquidation_value"`
	PercentUse                 int    `json:"percent_use"`
	AssessmentDate             string `json:"assessment_date"`
	AssessmentById             int    `json:"assessment_by_id"`
	AssessmentBy               string `json:"assessment_by" gorm:"-:migration"`
	Status                     string `json:"status"`
}

type Dropdown struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type SubmissionType struct {
	Id         int      `json:"id" gorm:"primaryKey"`
	Code       string   `json:"code"`
	Name       string   `json:"name"`
	SibsCode   *string  `json:"sibs_code" gorm:"default:null"`
	Scoring    *string  `json:"scoring" gorm:"default:null"`
	StopTrack  *string  `json:"stop_track" gorm:"default:null"`
	GoToTrack  *string  `json:"go_to_track" gorm:"default:null"`
	SibsLimit  *float64 `json:"sibs_limit" gorm:"default:null"`
	Channeling *string  `json:"channeling" gorm:"default:null"`
}

type CreditType struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type Purpose struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type CollateralType struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	LinkTable    string `json:"link_table"`
	CodeIbs      string `json:"code_ibs"`
	ReqAppraisal string `json:"req_appraisal"`
	RatingCode   string `json:"rating_code"`
}

type ProofOfOwnership struct {
	Id         int     `json:"id" gorm:"primaryKey"`
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	Flag       *string `json:"flag" gorm:"default:null"`
	RatingCode *string `json:"rating_code" gorm:"default:null"`
}

type FormOfBinding struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type CollateralClassification struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type Currency struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type Assessment struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
