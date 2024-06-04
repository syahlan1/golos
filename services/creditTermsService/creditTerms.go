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

		loanInformation := data.LoanNew
		loanInformation.CreditId = data.Id
		loanInformation.Status = "L"

		if err := connection.DB.Create(&loanInformation).Error; err != nil {
			return err
		}

		// check if Collateral is not empty
		if loanInformation.CollateralStatus {

			guarantee := data.LoanNew.Collateral
			guarantee.LoanId = loanInformation.Id
			guarantee.Status = "L"

			if err := connection.DB.Create(&guarantee).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func ShowCreditTerms(id string) (result []models.CreditTerms, err error) {

	if err = connection.DB.Where("status != ? AND general_information_id = ?", "D", id).Find(&result).Error; err != nil {
		return result, err
	}

	for i, value := range result {

		if err = connection.DB.Where("credit_id = ?", value.Id).Not("status", "D").First(&result[i].LoanNew).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return result, err
		}

		if result[i].LoanNew.Id == 0 {
			result[i].LoanNew = nil
			continue
		}

		if !result[i].LoanNew.CollateralStatus {
			result[i].LoanNew.Collateral = nil
			continue
		}

		if err = connection.DB.Where("loan_id = ?", result[i].LoanNew.Id).Not("status", "D").First(&result[i].LoanNew.Collateral).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return result, err
		}
	}

	return result, nil
}

func UpdateCreditTerms(id string, data models.CreditTerms) (result models.CreditTerms, err error) {

	if err = connection.DB.Where("id = ?", id).Not("status", "D").First(&result).Error; err != nil {
		return result, err
	}

	if err = connection.DB.Where("credit_id = ?", id).Not("status", "D").First(&result.LoanNew).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("error1", err)
		return result, err
	}

	if err = connection.DB.Where("loan_id = ?", id).Not("status", "D").First(&result.LoanNew.Collateral).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("error2", err)
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

	// check if LoanNew is not empty
	if data.LoanNew != nil {

		// result.LoanNew.SubmissionType = data.LoanNew.SubmissionType
		result.LoanNew.CreditType = data.LoanNew.CreditType
		result.LoanNew.Limit = data.LoanNew.Limit
		result.LoanNew.ExchangeRate = data.LoanNew.ExchangeRate
		result.LoanNew.LimitRp = data.LoanNew.LimitRp
		result.LoanNew.TimePeriod = data.LoanNew.TimePeriod
		result.LoanNew.PeriodType = data.LoanNew.PeriodType
		result.LoanNew.Purpose = data.LoanNew.Purpose
		result.LoanNew.Description = data.LoanNew.Description
		result.LoanNew.CollateralStatus = data.LoanNew.CollateralStatus
		result.LoanNew.DeptorTransfer = data.LoanNew.DeptorTransfer

		loanInformation := result.LoanNew

		if err := connection.DB.Save(&loanInformation).Error; err != nil {
			return result, err
		}

		// check if Collateral is not empty
		if loanInformation.CollateralStatus {

			result.LoanNew.Collateral.CollateralType = data.LoanNew.Collateral.CollateralType
			result.LoanNew.Collateral.Description = data.LoanNew.Collateral.Description
			result.LoanNew.Collateral.IdCoreCollateral = data.LoanNew.Collateral.IdCoreCollateral
			result.LoanNew.Collateral.ProofOfOwnership = data.LoanNew.Collateral.ProofOfOwnership
			result.LoanNew.Collateral.FormOfBinding = data.LoanNew.Collateral.FormOfBinding
			result.LoanNew.Collateral.CollateralClassification = data.LoanNew.Collateral.CollateralClassification
			result.LoanNew.Collateral.CurrencyId = data.LoanNew.Collateral.CurrencyId
			result.LoanNew.Collateral.ExchangeRate = data.LoanNew.Collateral.ExchangeRate
			result.LoanNew.Collateral.BankValue = data.LoanNew.Collateral.BankValue
			result.LoanNew.Collateral.MarketValue = data.LoanNew.Collateral.MarketValue
			result.LoanNew.Collateral.InsuranceValue = data.LoanNew.Collateral.InsuranceValue
			result.LoanNew.Collateral.BindingValue = data.LoanNew.Collateral.BindingValue
			result.LoanNew.Collateral.PPADeductionValue = data.LoanNew.Collateral.PPADeductionValue
			result.LoanNew.Collateral.LiquidationValue = data.LoanNew.Collateral.LiquidationValue
			result.LoanNew.Collateral.AssessmentDate = data.LoanNew.Collateral.AssessmentDate
			result.LoanNew.Collateral.AssessmentBy = data.LoanNew.Collateral.AssessmentBy

			guarantee := result.LoanNew.Collateral

			if err := connection.DB.Save(&guarantee).Error; err != nil {
				return result, err
			}
		} else {
			// if empty then delete
			guarantee := result.LoanNew.Collateral
			guarantee.Status = "D"

			if err := connection.DB.Save(&guarantee).Error; err != nil {
				return result, err
			}
			result.LoanNew.Collateral = nil
		}
	} else {
		// if empty then delete
		loanInformation := result.LoanNew
		loanInformation.Status = "D"

		if err := connection.DB.Save(&loanInformation).Error; err != nil {
			return result, err
		}
		result.LoanNew = nil
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
