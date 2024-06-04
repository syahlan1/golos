package creditTermsService

import (
	"errors"
	"log"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
	"gorm.io/gorm"
)

func GetSubmissionType() (result []models.Dropdown, err error) {

	if err = connection.DB.Select("id", "name").
		Model(&models.SubmissionType{}).
		Find(&result).Error; err != nil {
		return nil, err
	}

	result = utils.Prepend(result, models.Dropdown{Id: 0, Name: "- PILIH -"})

	return
}

func GetCreditType() (result []models.Dropdown, err error) {

	if err = connection.DB.Select("id", "CONCAT(code, ' - ', name) AS name").
		Model(&models.CreditType{}).
		Find(&result).Error; err != nil {
		return nil, err
	}

	result = utils.Prepend(result, models.Dropdown{Id: 0, Name: "- PILIH -"})

	return
}

func GetCreditPurpose() (result []models.Dropdown, err error) {
	if err = connection.DB.Select("id", "CONCAT(code, ' - ', name) AS name").
		Model(&models.Purpose{}).
		Find(&result).Error; err != nil {
		return nil, err
	}

	result = utils.Prepend(result, models.Dropdown{Id: 0, Name: "- PILIH -"})

	return
}

func GetCollateralType() (result []models.Dropdown, err error) {
	if err = connection.DB.Select("id", "CONCAT(code, ' - ', name) AS name").
		Order("code").
		Model(&models.CollateralType{}).
		Find(&result).Error; err != nil {
		return nil, err
	}

	result = utils.Prepend(result, models.Dropdown{Id: 0, Name: "- PILIH -"})

	return
}

func GetProofOfOwnership() (result []models.Dropdown, err error) {
	if err = connection.DB.Select("id", "name").
		Model(&models.ProofOfOwnership{}).
		Find(&result).Error; err != nil {
		return nil, err
	}

	result = utils.Prepend(result, models.Dropdown{Id: 0, Name: "- PILIH -"})

	return
}

func GetFormOfBinding() (result []models.Dropdown, err error) {
	if err = connection.DB.Select("id", "name").
		Model(&models.FormOfBinding{}).
		Find(&result).Error; err != nil {
		return nil, err
	}

	result = utils.Prepend(result, models.Dropdown{Id: 0, Name: "- PILIH -"})

	return
}

func GetCollateralClassification() (result []models.Dropdown, err error) {
	if err = connection.DB.Select("id", "name").
		Model(&models.CollateralClassification{}).
		Find(&result).Error; err != nil {
		return nil, err
	}

	result = utils.Prepend(result, models.Dropdown{Id: 0, Name: "- PILIH -"})

	return
}

func GetCreditCurrency() (result []models.Dropdown, err error) {
	if err = connection.DB.Select("id", "CONCAT(code, ' - ', name) AS name").
		Model(&models.Currency{}).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return
}

func GetAssessmentBy() (result []models.Assessment, err error) {
	if err = connection.DB.Select("id", "name").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return
}

func CreateCreditTerms(data *models.CreditTerms) (err error) {

	// creditTerms := data.CreditTerms
	data.Status = "L"

	if err := connection.DB.Create(&data).Error; err != nil {
		return err
	}

	// check if LoanNew is not empty
	if data.LoanNew != nil {

		loanNew := data.LoanNew
		loanNew.CreditId = data.Id
		loanNew.Status = "L"

		if err := connection.DB.Create(&loanNew).Error; err != nil {
			return err
		}

		// check if Collateral is not empty
		if loanNew.CollateralStatus {

			collateral := data.LoanNew.Collateral
			collateral.LoanId = loanNew.Id
			collateral.Status = "L"

			if err := connection.DB.Create(&collateral).Error; err != nil {
				return err
			}
		}
		// check if LoanRenewal is not empty
	} else if data.LoanRenewal != nil {

		loanNew := data.LoanRenewal
		loanNew.CreditId = data.Id
		loanNew.Status = "L"

		if err := connection.DB.Create(&loanNew).Error; err != nil {
			return err
		}
	}

	return nil
}

func ShowCreditTerms(id string) (result []models.CreditTerms, err error) {

	if err = connection.DB.Where("status != ? AND general_information_id = ?", "D", id).Find(&result).Error; err != nil {
		return result, err
	}

	for i, value := range result {

		// check if LoanNew and LoanRenewal is exist
		var checkLoan models.CheckLoan
		if err = connection.DB.Select(`
			CASE WHEN ln.id IS NOT NULL THEN TRUE ELSE FALSE END AS loan_new, 
			CASE WHEN lr.id IS NOT NULL THEN TRUE ELSE FALSE END AS loan_renewal`).
			Table("credit_terms ct").
			Joins("LEFT JOIN loan_new ln ON ln.credit_id = ct.id").
			Joins("LEFT JOIN loan_renewals lr ON lr.credit_id = ct.id").
			Where("ct.id = ?", value.Id).
			Where("ln.status <> ? OR lr.status <> ?", "D", "D").
			Find(&checkLoan).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return result, err
		}

		switch {
		case checkLoan.LoanNew:

			if err = connection.DB.Select("ln.*, st.name AS submission_type, CONCAT(ct.code, ' - ', ct.name) AS credit_type, CONCAT(p.code, ' - ', p.name) AS purpose").
				Table("loan_new AS ln").
				Joins("JOIN submission_types AS st ON st.id = ln.submission_type_id").
				Joins("JOIN credit_types AS ct ON ct.id = ln.credit_type_id").
				Joins("JOIN purposes AS p ON p.id = ln.purpose_id").
				Where("credit_id = ?", value.Id).Not("status", "D").First(&result[i].LoanNew).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return result, err
			}

			if !result[i].LoanNew.CollateralStatus {
				result[i].LoanNew.Collateral = nil
				continue
			}

			if err = connection.DB.Select("c.*, CONCAT(ct.code, ' - ', ct.name) AS collateral_type, po.name AS proof_of_ownership, fob.name AS form_of_binding, cc.name AS collateral_classification, CONCAT(cur.code, ' - ', cur.name) AS currency, asmt.name AS assessment_by").
				Table("collaterals AS c").
				Joins("JOIN collateral_types AS ct ON ct.id = c.collateral_type_id").
				Joins("JOIN proof_of_ownerships AS po ON po.id = c.proof_of_ownership_id").
				Joins("JOIN form_of_bindings AS fob ON fob.id = c.form_of_binding_id").
				Joins("JOIN collateral_classifications AS cc ON cc.id = c.collateral_classification_id").
				Joins("JOIN currencies AS cur ON cur.id = c.currency_id").
				Joins("JOIN assessments AS asmt ON asmt.id = c.assessment_by_id").
				Where("loan_id = ?", result[i].LoanNew.Id).Not("status", "D").First(&result[i].LoanNew.Collateral).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return result, err
			}

		case checkLoan.LoanRenewal:

			if err = connection.DB.Select("lr.*, st.name AS submission_type, CONCAT(p.code, ' - ', p.name) AS purpose").
				Table("loan_renewals AS lr").
				Joins("JOIN submission_types AS st ON st.id = lr.submission_type_id").
				Joins("JOIN purposes AS p ON p.id = lr.purpose_id").
				Where("credit_id = ?", value.Id).Not("status", "D").First(&result[i].LoanRenewal).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return result, err
			}
		}

	}

	return result, nil
}

func UpdateCreditTerms(id string, data models.CreditTerms) (result models.CreditTerms, err error) {

	if err = connection.DB.Where("id = ?", id).Not("status", "D").First(&result).Error; err != nil {
		return result, err
	}

	result.NewApplication = data.NewApplication
	result.ChannelingCompany = data.ChannelingCompany
	result.TakeoverBank = data.TakeoverBank
	result.AccountAccomodation = data.AccountAccomodation
	result.NCLProduct = data.NCLProduct
	result.Facility = data.Facility
	result.AccountNumber = data.AccountNumber
	result.ApplicationType = data.ApplicationType
	result.Project = data.Project
	result.Status = data.Status

	if err := connection.DB.Save(&data).Error; err != nil {
		return result, err
	}

	var checkLoan models.CheckLoan
	if err = connection.DB.Select(`
		CASE WHEN ln.id IS NOT NULL THEN TRUE ELSE FALSE END AS loan_new, 
		CASE WHEN lr.id IS NOT NULL THEN TRUE ELSE FALSE END AS loan_renewal`).
		Table("credit_terms ct").
		Joins("LEFT JOIN loan_new ln ON ln.credit_id = ct.id").
		Joins("LEFT JOIN loan_renewals lr ON lr.credit_id = ct.id").
		Where("ct.id = ?", id).
		Where("ln.status <> ? OR lr.status <> ?", "D", "D").
		Find(&checkLoan).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return result, err
	}

	switch {
	case checkLoan.LoanNew:

		// check if LoanNew is not empty
		if data.LoanNew != nil {

			if err = connection.DB.Where("credit_id = ?", id).Not("status", "D").First(&result.LoanNew).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				log.Println("error1", err)
				return result, err
			}

			// result.LoanNew.SubmissionType = data.LoanNew.SubmissionType
			result.LoanNew.CreditTypeId = data.LoanNew.CreditTypeId
			result.LoanNew.Limit = data.LoanNew.Limit
			result.LoanNew.ExchangeRate = data.LoanNew.ExchangeRate
			result.LoanNew.LimitRp = data.LoanNew.LimitRp
			result.LoanNew.TimePeriod = data.LoanNew.TimePeriod
			result.LoanNew.PeriodType = data.LoanNew.PeriodType
			result.LoanNew.PurposeId = data.LoanNew.PurposeId
			result.LoanNew.Description = data.LoanNew.Description
			result.LoanNew.CollateralStatus = data.LoanNew.CollateralStatus
			result.LoanNew.DeptorTransfer = data.LoanNew.DeptorTransfer
			result.LoanNew.OldCifNo = data.LoanNew.OldCifNo
			result.LoanNew.OldAccountNo = data.LoanNew.OldAccountNo

			loanNew := result.LoanNew

			if err := connection.DB.Save(&loanNew).Error; err != nil {
				return result, err
			}

			// check if Collateral is not empty
			if loanNew.CollateralStatus {

				if err = connection.DB.Where("loan_id = ?", result.LoanNew.Id).Not("status", "D").First(&result.LoanNew.Collateral).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					log.Println("error2", err)
					return result, err
				}

				result.LoanNew.Collateral.CollateralTypeId = data.LoanNew.Collateral.CollateralTypeId
				result.LoanNew.Collateral.Description = data.LoanNew.Collateral.Description
				result.LoanNew.Collateral.IdCoreCollateral = data.LoanNew.Collateral.IdCoreCollateral
				result.LoanNew.Collateral.ProofOfOwnershipId = data.LoanNew.Collateral.ProofOfOwnershipId
				result.LoanNew.Collateral.FormOfBindingId = data.LoanNew.Collateral.FormOfBindingId
				result.LoanNew.Collateral.CollateralClassificationId = data.LoanNew.Collateral.CollateralClassificationId
				result.LoanNew.Collateral.CurrencyId = data.LoanNew.Collateral.CurrencyId
				result.LoanNew.Collateral.ExchangeRate = data.LoanNew.Collateral.ExchangeRate
				result.LoanNew.Collateral.BankValue = data.LoanNew.Collateral.BankValue
				result.LoanNew.Collateral.MarketValue = data.LoanNew.Collateral.MarketValue
				result.LoanNew.Collateral.InsuranceValue = data.LoanNew.Collateral.InsuranceValue
				result.LoanNew.Collateral.BindingValue = data.LoanNew.Collateral.BindingValue
				result.LoanNew.Collateral.PPADeductionValue = data.LoanNew.Collateral.PPADeductionValue
				result.LoanNew.Collateral.LiquidationValue = data.LoanNew.Collateral.LiquidationValue
				result.LoanNew.Collateral.AssessmentDate = data.LoanNew.Collateral.AssessmentDate
				result.LoanNew.Collateral.AssessmentById = data.LoanNew.Collateral.AssessmentById

				collateral := result.LoanNew.Collateral

				if err := connection.DB.Save(&collateral).Error; err != nil {
					return result, err
				}
			} else {
				// if collateral empty then delete
				collateral := result.LoanNew.Collateral
				collateral.Status = "D"

				if err := connection.DB.Save(&collateral).Error; err != nil {
					return result, err
				}
				result.LoanNew.Collateral = nil
			}
		} else {
			// if LoanNew empty then delete
			loanNew := result.LoanNew
			loanNew.Status = "D"

			if err := connection.DB.Save(&loanNew).Error; err != nil {
				return result, err
			}
			result.LoanNew = nil
		}

	case checkLoan.LoanRenewal:

		// check if LoanRenewal is not empty
		if data.LoanRenewal != nil {
			if err = connection.DB.Where("credit_id = ?", id).Not("status", "D").First(&result.LoanRenewal).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				log.Println("error1", err)
				return result, err
			}

			result.LoanRenewal.SubmissionTypeId = data.LoanRenewal.SubmissionTypeId
			result.LoanRenewal.AccountAccomodation = data.LoanRenewal.AccountAccomodation
			result.LoanRenewal.AccountNumber = data.LoanRenewal.AccountNumber
			result.LoanRenewal.FacilityNumber = data.LoanRenewal.FacilityNumber
			result.LoanRenewal.CreditType = data.LoanRenewal.CreditType
			result.LoanRenewal.Limit = data.LoanRenewal.Limit
			result.LoanRenewal.TimePeriod = data.LoanRenewal.TimePeriod
			result.LoanRenewal.PeriodType = data.LoanRenewal.PeriodType
			result.LoanRenewal.PurposeId = data.LoanRenewal.PurposeId
			result.LoanRenewal.TimePeriodRequest = data.LoanRenewal.TimePeriodRequest
			result.LoanRenewal.LimitRequest = data.LoanRenewal.LimitRequest
			result.LoanRenewal.ExchangeRate = data.LoanRenewal.ExchangeRate
			result.LoanRenewal.LimitRp = data.LoanRenewal.LimitRp
			result.LoanRenewal.Description = data.LoanRenewal.Description

			LoanRenewal := result.LoanRenewal

			if err := connection.DB.Save(&LoanRenewal).Error; err != nil {
				return result, err
			}

		} else {
			// if LoanRenewal empty then delete
			LoanRenewal := result.LoanRenewal
			LoanRenewal.Status = "D"

			if err := connection.DB.Save(&LoanRenewal).Error; err != nil {
				return result, err
			}
			result.LoanRenewal = nil
		}

	}

	return result, nil
}

func DeleteCreditTerms(id string) (result models.CreditTerms, err error) {
	connection.DB.First(&result, id)

	result.Status = "D"

	if err = connection.DB.Save(&result).Error; err != nil {
		return
	}
	return
}
