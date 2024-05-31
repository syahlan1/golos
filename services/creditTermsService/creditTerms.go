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

func GetCreditUsage() (result []models.Dropdown, err error) {
	if err = connection.DB.Select("id", "CONCAT(code, ' - ', name) AS name").
		Model(&models.Usage{}).
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

		// check if Guarantee is not empty
		if loanInformation.GuaranteeStatus {

			guarantee := data.LoanInformation.Guarantee
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

		if !result[i].LoanInformation.GuaranteeStatus {
			result[i].LoanInformation.Guarantee = nil
			continue
		}

		if err = connection.DB.Where("loan_id = ?", result[i].LoanInformation.Id).Not("status", "D").First(&result[i].LoanInformation.Guarantee).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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

	if err = connection.DB.Where("loan_id = ?", id).Not("status", "D").First(&result.LoanInformation.Guarantee).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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
		result.LoanInformation.Usage = data.LoanInformation.Usage
		result.LoanInformation.Description = data.LoanInformation.Description
		result.LoanInformation.GuaranteeStatus = data.LoanInformation.GuaranteeStatus
		result.LoanInformation.DeptorTransfer = data.LoanInformation.DeptorTransfer

		loanInformation := result.LoanInformation

		if err := connection.DB.Save(&loanInformation).Error; err != nil {
			return result, err
		}

		// check if Guarantee is not empty
		if loanInformation.GuaranteeStatus {

			result.LoanInformation.Guarantee.GuaranteeType = data.LoanInformation.Guarantee.GuaranteeType
			result.LoanInformation.Guarantee.Description = data.LoanInformation.Guarantee.Description
			result.LoanInformation.Guarantee.IdCoreCollateral = data.LoanInformation.Guarantee.IdCoreCollateral
			result.LoanInformation.Guarantee.ProofOfOwnership = data.LoanInformation.Guarantee.ProofOfOwnership
			result.LoanInformation.Guarantee.FormOfBinding = data.LoanInformation.Guarantee.FormOfBinding
			result.LoanInformation.Guarantee.GuaranteeClassification = data.LoanInformation.Guarantee.GuaranteeClassification
			result.LoanInformation.Guarantee.CurrencyId = data.LoanInformation.Guarantee.CurrencyId
			result.LoanInformation.Guarantee.ExchangeRate = data.LoanInformation.Guarantee.ExchangeRate
			result.LoanInformation.Guarantee.BankValue = data.LoanInformation.Guarantee.BankValue
			result.LoanInformation.Guarantee.MarketValue = data.LoanInformation.Guarantee.MarketValue
			result.LoanInformation.Guarantee.InsuranceValue = data.LoanInformation.Guarantee.InsuranceValue
			result.LoanInformation.Guarantee.BindingValue = data.LoanInformation.Guarantee.BindingValue
			result.LoanInformation.Guarantee.PPADeductionValue = data.LoanInformation.Guarantee.PPADeductionValue

			guarantee := result.LoanInformation.Guarantee

			if err := connection.DB.Save(&guarantee).Error; err != nil {
				return result, err
			}
		} else {
			// if empty then delete
			guarantee := result.LoanInformation.Guarantee
			guarantee.Status = "D"

			if err := connection.DB.Save(&guarantee).Error; err != nil {
				return result, err
			}
			result.LoanInformation.Guarantee = nil
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
