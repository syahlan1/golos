package models

type CreditTerms struct {
	Id                   int    `json:"id" gorm:"primaryKey"`
	GeneralInformationId int    `json:"general_information_id"`
	NewApplication       int    `json:"new_application"`
	ChannelingCompany    int    `json:"channeling_company"`
	TakeoverBank         bool   `json:"takeover_bank"`
	AccountAccomodation  int    `json:"account_accommodation"`
	NCLProduct           int    `json:"ncl_product"`
	Facility             int    `json:"facility"`
	AccountNumber        string `json:"account_number"`
	ApplicationType      string `json:"application_type"`
	Project              int    `json:"project"`
}

type LoanInformation struct {
	Id             int    `json:"id" gorm:"primaryKey"`
	CreditId       int    `json:"credit_id"`
	SubmissionType int    `json:"submission_type"`
	CreditType     int    `json:"credit_type"`
	Limit          int    `json:"limit"`
	ExchangeRate   int    `json:"exchange_rate"`
	LimitRp        int    `json:"limit_rp"`
	TimePeriod     int    `json:"time_period"`
	PeriodType     string `json:"period_type"`
	Usage          int    `json:"usage"`
	Description    string `json:"description" gorm:"type:text"`
	Guarantee      bool   `json:"guarantee"`
	DeptorTransfer bool   `json:"depositor_transfer"`
}

type Guarantee struct {
	Id                      int    `json:"id" gorm:"primaryKey"`
	LoanId                  int    `json:"loan_id"`
	GuaranteeType           int    `json:"guarantee_type"`
	Description             string `json:"description"`
	IdCoreCollateral        int    `json:"id_core_collateral"`
	ProofOfOwnership        int    `json:"proof_of_ownership"`
	FormOfBinding           int    `json:"form_of_binding"`
	GuaranteeClassification int    `json:"guarantee_classification"`
	CurrencyId              int    `json:"currency_id"`
	ExchangeRate            int    `json:"exchange_rate"`
	BankValue               int    `json:"bank_value"`
	MarketValue             int    `json:"market_value"`
	InsuranceValue          int    `json:"insurance_value"`
	BindingValue            int    `json:"binding_value"`
	PPADeductionValue       int    `json:"ppa_deduction_value"`
}

type CreditType struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type Usage struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type GuaranteeType struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type ProofOfOwnership struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type FormOfBinding struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type GuaranteeClassification struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type Currency struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code"`
	Name string `json:"name"`
}
