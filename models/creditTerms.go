package models

type CreditTerms struct {
	Id                   int              `json:"id" gorm:"primaryKey"`
	GeneralInformationId int              `json:"general_information_id"`
	NewApplication       int              `json:"new_application"`
	ChannelingCompany    int              `json:"channeling_company"`
	TakeoverBank         bool             `json:"takeover_bank"`
	AccountAccomodation  int              `json:"account_accommodation"`
	NCLProduct           int              `json:"ncl_product"`
	Facility             int              `json:"facility"`
	AccountNumber        string           `json:"account_number"`
	ApplicationType      string           `json:"application_type"`
	Project              int              `json:"project"`
	Status               string           `json:"status"`
	LoanInformation      *LoanInformation `json:"loan_information" gorm:"-:all"`
}

// type CreateCreditTerms struct {
// 	*CreditTerms
// 	LoanInformation *LoanInformation `json:"loan_information"`
// 	Collateral       *Collateral       `json:"collateral"`
// }

type LoanInformation struct {
	Id               int         `json:"id" gorm:"primaryKey"`
	CreditId         int         `json:"-"`
	SubmissionType   int         `json:"submission_type"`
	CreditType       int         `json:"credit_type"`
	Limit            int         `json:"limit"`
	ExchangeRate     int         `json:"exchange_rate"`
	LimitRp          int         `json:"limit_rp"`
	TimePeriod       int         `json:"time_period"`
	PeriodType       string      `json:"period_type"`
	Purpose          int         `json:"purpose"`
	Description      string      `json:"description" gorm:"type:text"`
	CollateralStatus bool        `json:"collateral_status"`
	Collateral       *Collateral `json:"collateral" gorm:"-:all"`
	DeptorTransfer   bool        `json:"depositor_transfer"`
	Status           string      `json:"status"`
}

type Collateral struct {
	Id                       int    `json:"id" gorm:"primaryKey"`
	LoanId                   int    `json:"-"`
	CollateralType           int    `json:"collateral_type"`
	Description              string `json:"description"`
	IdCoreCollateral         int    `json:"id_core_collateral"`
	ProofOfOwnership         int    `json:"proof_of_ownership"`
	FormOfBinding            int    `json:"form_of_binding"`
	CollateralClassification int    `json:"collateral_classification"`
	CurrencyId               int    `json:"currency_id"`
	ExchangeRate             int    `json:"exchange_rate"`
	BankValue                int    `json:"bank_value"`
	MarketValue              int    `json:"market_value"`
	InsuranceValue           int    `json:"insurance_value"`
	BindingValue             int    `json:"binding_value"`
	PPADeductionValue        int    `json:"ppa_deduction_value"`
	LiquidationValue         int    `json:"liquidation_value"`
	AssessmentDate           string `json:"assessment_date"`
	AssessmentBy             int    `json:"assessment_by"`
	Status                   string `json:"status"`
}

type Dropdown struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
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
	Flag       *string `json:"flag"`
	RatingCode *string `json:"rating_code"`
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
