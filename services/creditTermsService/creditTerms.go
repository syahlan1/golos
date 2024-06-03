package creditTermsService

import (
	"errors"
	"log"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
	"gorm.io/gorm"
)

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

	// check if LoanInformation is not empty
	if data.LoanInformation != nil {

		loanInformation := data.LoanInformation
		loanInformation.CreditId = data.Id
		loanInformation.Status = "L"

		if err := connection.DB.Create(&loanInformation).Error; err != nil {
			return err
		}

		// check if Collateral is not empty
		if loanInformation.CollateralStatus {

			guarantee := data.LoanInformation.Collateral
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

		if err = connection.DB.Where("credit_id = ?", value.Id).Not("status", "D").First(&result[i].LoanInformation).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return result, err
		}

		if result[i].LoanInformation.Id == 0 {
			result[i].LoanInformation = nil
			continue
		}

		if !result[i].LoanInformation.CollateralStatus {
			result[i].LoanInformation.Collateral = nil
			continue
		}

		if err = connection.DB.Where("loan_id = ?", result[i].LoanInformation.Id).Not("status", "D").First(&result[i].LoanInformation.Collateral).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return result, err
		}
	}

	return result, nil
}

func UpdateCreditTerms(id string, data models.CreditTerms) (result models.CreditTerms, err error) {

	if err = connection.DB.Where("id = ?", id).Not("status", "D").First(&result).Error; err != nil {
		return result, err
	}

	if err = connection.DB.Where("credit_id = ?", id).Not("status", "D").First(&result.LoanInformation).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("error1", err)
		return result, err
	}

	if err = connection.DB.Where("loan_id = ?", id).Not("status", "D").First(&result.LoanInformation.Collateral).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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

	// check if LoanInformation is not empty
	if data.LoanInformation != nil {

		result.LoanInformation.SubmissionType = data.LoanInformation.SubmissionType
		result.LoanInformation.CreditType = data.LoanInformation.CreditType
		result.LoanInformation.Limit = data.LoanInformation.Limit
		result.LoanInformation.ExchangeRate = data.LoanInformation.ExchangeRate
		result.LoanInformation.LimitRp = data.LoanInformation.LimitRp
		result.LoanInformation.TimePeriod = data.LoanInformation.TimePeriod
		result.LoanInformation.PeriodType = data.LoanInformation.PeriodType
		result.LoanInformation.Purpose = data.LoanInformation.Purpose
		result.LoanInformation.Description = data.LoanInformation.Description
		result.LoanInformation.CollateralStatus = data.LoanInformation.CollateralStatus
		result.LoanInformation.DeptorTransfer = data.LoanInformation.DeptorTransfer

		loanInformation := result.LoanInformation

		if err := connection.DB.Save(&loanInformation).Error; err != nil {
			return result, err
		}

		// check if Collateral is not empty
		if loanInformation.CollateralStatus {

			result.LoanInformation.Collateral.CollateralType = data.LoanInformation.Collateral.CollateralType
			result.LoanInformation.Collateral.Description = data.LoanInformation.Collateral.Description
			result.LoanInformation.Collateral.IdCoreCollateral = data.LoanInformation.Collateral.IdCoreCollateral
			result.LoanInformation.Collateral.ProofOfOwnership = data.LoanInformation.Collateral.ProofOfOwnership
			result.LoanInformation.Collateral.FormOfBinding = data.LoanInformation.Collateral.FormOfBinding
			result.LoanInformation.Collateral.CollateralClassification = data.LoanInformation.Collateral.CollateralClassification
			result.LoanInformation.Collateral.CurrencyId = data.LoanInformation.Collateral.CurrencyId
			result.LoanInformation.Collateral.ExchangeRate = data.LoanInformation.Collateral.ExchangeRate
			result.LoanInformation.Collateral.BankValue = data.LoanInformation.Collateral.BankValue
			result.LoanInformation.Collateral.MarketValue = data.LoanInformation.Collateral.MarketValue
			result.LoanInformation.Collateral.InsuranceValue = data.LoanInformation.Collateral.InsuranceValue
			result.LoanInformation.Collateral.BindingValue = data.LoanInformation.Collateral.BindingValue
			result.LoanInformation.Collateral.PPADeductionValue = data.LoanInformation.Collateral.PPADeductionValue
			result.LoanInformation.Collateral.LiquidationValue = data.LoanInformation.Collateral.LiquidationValue
			result.LoanInformation.Collateral.AssessmentDate = data.LoanInformation.Collateral.AssessmentDate
			result.LoanInformation.Collateral.AssessmentBy = data.LoanInformation.Collateral.AssessmentBy

			guarantee := result.LoanInformation.Collateral

			if err := connection.DB.Save(&guarantee).Error; err != nil {
				return result, err
			}
		} else {
			// if empty then delete
			guarantee := result.LoanInformation.Collateral
			guarantee.Status = "D"

			if err := connection.DB.Save(&guarantee).Error; err != nil {
				return result, err
			}
			result.LoanInformation.Collateral = nil
		}
	} else {
		// if empty then delete
		loanInformation := result.LoanInformation
		loanInformation.Status = "D"

		if err := connection.DB.Save(&loanInformation).Error; err != nil {
			return result, err
		}
		result.LoanInformation = nil
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
