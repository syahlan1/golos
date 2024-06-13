package ownershipDataService

import (
	"errors"
	"strconv"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
)

func CreateOwnershipData(generalInformationId string, data models.CreateOwnershipData) (err error) {

	if generalInformationId == "" {
		return errors.New("generalInformationId cannot be empty")
	}
	generalInformationIdInt, err := strconv.Atoi(generalInformationId)
	if err != nil {
		return err
	}

	ownership := models.OwnershipData{
		GeneralInformationId: generalInformationIdInt,
		OwnershipType:        data.OwnershipType,
		Name:                 data.Name,
		NoIdentity:           data.NoIdentity,
		IdCardAddress:        data.IdCardAddress,
		City:                 data.City,
		ZipCode:              data.ZipCode,
		HomeOwnership:        data.HomeOwnership,
		Remark:               data.Remark,
		CifManager:           data.CifManager,
		BirthDate:            data.BirthDate,
		LastEducation:        data.LastEducation,
		NPWP:                 data.NPWP,
		JobTitle:             data.JobTitle,
		Experince:            data.Experince,
		OwnershipMarket:      data.OwnershipMarket,
		CitizenshipStatus:    data.CitizenshipStatus,
		Gender:               data.Gender,
		MaritalStatus:        data.MaritalStatus,
		NumberOfChildren:     data.NumberOfChildren,
		StartDate:            data.StartDate,
		KeyPerson:            data.KeyPerson,
		Removed:              data.Removed,
		Status:               "L",
	}

	var totalOwnershipMarket float64
	if err := connection.DB.Select("SUM(ownership_market)").
		Table("ownership_data").
		Where("general_information_id = ? AND status = ?", generalInformationId, "L").
		Scan(&totalOwnershipMarket).Error; err != nil {
		return err
	}

	if totalOwnershipMarket+data.OwnershipMarket > 100 {
		return errors.New("total ownership market cannot more than 100%")
	}

	if err := connection.DB.Create(&ownership).Error; err != nil {
		return err
	}

	return nil
}

func ShowOwnershipData(generalInformationId string) (result []models.OwnershipData) {
	var ownershipData []models.OwnershipData

	if generalInformationId != "" {
		connection.DB.Where("status = ? AND general_information_id = ?", "L", generalInformationId).Find(&ownershipData)
	} else {
		connection.DB.Where("status = ?", "L").Find(&ownershipData)
	}

	return ownershipData
}

func ShowOwnershipName(generalInformationId string) (result []models.OwnershipDataDropdown) {
	connection.DB.Select("id, name, no_identity, npwp, key_person").
		Model(models.OwnershipData{}).
		Where("status = ? AND general_information_id = ?", "L", generalInformationId).Find(&result)

	result = utils.Prepend(result, models.OwnershipDataDropdown{Name: "- SELECT -"})

	return

}

func EditOwnershipData(id string, updatedOwnershipData models.OwnershipData) (result models.OwnershipData, err error) {
	var ownershipData models.OwnershipData
	if err := connection.DB.First(&ownershipData, id).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	ownershipData.Name = updatedOwnershipData.Name
	ownershipData.NoIdentity = updatedOwnershipData.NoIdentity
	ownershipData.IdCardAddress = updatedOwnershipData.IdCardAddress
	ownershipData.City = updatedOwnershipData.City
	ownershipData.ZipCode = updatedOwnershipData.ZipCode
	ownershipData.HomeOwnership = updatedOwnershipData.HomeOwnership
	ownershipData.Remark = updatedOwnershipData.Remark
	ownershipData.CifManager = updatedOwnershipData.CifManager
	ownershipData.BirthDate = updatedOwnershipData.BirthDate
	ownershipData.LastEducation = updatedOwnershipData.LastEducation
	ownershipData.NPWP = updatedOwnershipData.NPWP
	ownershipData.JobTitle = updatedOwnershipData.JobTitle
	ownershipData.Experince = updatedOwnershipData.Experince
	ownershipData.OwnershipMarket = updatedOwnershipData.OwnershipMarket
	ownershipData.CitizenshipStatus = updatedOwnershipData.CitizenshipStatus
	ownershipData.Gender = updatedOwnershipData.Gender
	ownershipData.MaritalStatus = updatedOwnershipData.MaritalStatus
	ownershipData.NumberOfChildren = updatedOwnershipData.NumberOfChildren
	ownershipData.StartDate = updatedOwnershipData.StartDate
	ownershipData.KeyPerson = updatedOwnershipData.KeyPerson
	ownershipData.Removed = updatedOwnershipData.Removed
	ownershipData.Status = updatedOwnershipData.Status

	if err := connection.DB.Save(&ownershipData).Error; err != nil {
		return result, errors.New("failed to update the user data")
	}

	return ownershipData, nil
}

func DeleteOwnershipData(ownershipId string) (result models.OwnershipData, err error) {
	var ownership models.OwnershipData
	if err := connection.DB.First(&ownership, ownershipId).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	ownership.Status = "D"

	if err := connection.DB.Save(&ownership).Error; err != nil {
		return result, errors.New("failed to delete")
	}

	return ownership, nil
}

func CreateRelationWithBank(generalInformationId string, data *models.RelationWithBank) (err error) {

	if generalInformationId == "" {
		return errors.New("generalInformationId cannot be empty")
	}
	generalInformationIdInt, err := strconv.Atoi(generalInformationId)
	if err != nil {
		return err
	}

	data.GeneralInformationId = generalInformationIdInt
	data.Status = "L"

	if err := connection.DB.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func CreateCustomerLoanInfo(generalInformationId string, data *models.CustomerLoanInfo) (err error) {

	if generalInformationId == "" {
		return errors.New("generalInformationId cannot be empty")
	}

	generalInformationIdInt, err := strconv.Atoi(generalInformationId)
	if err != nil {
		return err
	}
	data.GeneralInformationId = generalInformationIdInt
	data.Status = "L"

	if data.AAStatus == 2 {
		var checkAa int64

		if err := connection.DB.Select("id").
			Table("customer_loan_infos").
			Where("aa_no =?", data.AANo).
			Count(&checkAa).Error; err != nil {
			return err
		}

		if checkAa > 0 {
			return errors.New("aa_no already exist")
		}
	}

	if err := connection.DB.Create(&data).Error; err != nil {
		return err
	}

	return
}



func CreateRekeningDebitur(generalInformationId string, data *models.DataRekeningDebitur) (err error) {

	if generalInformationId == "" {
		return errors.New("generalInformationId cannot be empty")
	}

	generalInformationIdInt, err := strconv.Atoi(generalInformationId)
	if err != nil {
		return err
	}

	data.GeneralInformationId = generalInformationIdInt
	data.Status = "L"
	// log.Println(data)
	if err := connection.DB.Create(&data).Error; err != nil {
		return err
	}
	return
}

func UpdateRekeningDebitur(id string, updatedRekeningDebitur models.DataRekeningDebitur) (result models.DataRekeningDebitur, err error) {
	var rekeningDebitur models.DataRekeningDebitur
	if err := connection.DB.First(&rekeningDebitur, id).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	rekeningDebitur.NoRekening = updatedRekeningDebitur.NoRekening
	rekeningDebitur.Remark = updatedRekeningDebitur.Remark

	if err := connection.DB.Save(&rekeningDebitur).Error; err != nil {
		return result, errors.New("failed to update the user data")
	}

	return rekeningDebitur, nil
}

func UpdateRelationWithBank(id string, updateRelationBank models.RelationWithBank) (result models.RelationWithBank, err error) {
	var relationBank models.RelationWithBank
	if err := connection.DB.First(&relationBank, id).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	relationBank.Giro = updateRelationBank.Giro
	relationBank.Tabungan = updateRelationBank.Tabungan
	relationBank.NoRekening = updateRelationBank.NoRekening
	relationBank.Debitur = updateRelationBank.Debitur

	if err := connection.DB.Save(&relationBank).Error; err != nil {
		return result, errors.New("failed to update the user data")
	}

	return relationBank, nil
}

func UpdateCustomerLoanInfo(id string, updatedCustomerLoanInfo models.CustomerLoanInfo) (result models.CustomerLoanInfo, err error) {
	
	var customerLoanInfo models.CustomerLoanInfo
	if err := connection.DB.First(&customerLoanInfo, id).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	customerLoanInfo.NoRekening = updatedCustomerLoanInfo.NoRekening


	if err := connection.DB.Save(&customerLoanInfo).Error; err != nil {
		return result, errors.New("failed to update the user data")
	}

	return customerLoanInfo, nil
}

func DeleteRelationWithBank(Id string) (result models.RelationWithBank, err error) {
	var relation models.RelationWithBank
	if err := connection.DB.First(&relation, Id).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	relation.Status = "D"

	if err := connection.DB.Save(&relation).Error; err != nil {
		return result, errors.New("failed to delete")
	}

	return relation, nil
}

func DeleteRekeningDebitur(Id string) (result models.DataRekeningDebitur, err error) {
	var rekeningDebitur models.DataRekeningDebitur
	if err := connection.DB.First(&rekeningDebitur, Id).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	rekeningDebitur.Status = "D"

	if err := connection.DB.Save(&rekeningDebitur).Error; err != nil {
		return result, errors.New("failed to delete")
	}

	return rekeningDebitur, nil
}

func DeleteCustomerLoanInfo(id string) (result models.CustomerLoanInfo, err error) {
	var customerLoanInfo models.CustomerLoanInfo
	if err := connection.DB.First(&customerLoanInfo, id).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	customerLoanInfo.Status = "D"

	if err := connection.DB.Save(&customerLoanInfo).Error; err != nil {
		return result, errors.New("failed to delete")
	}

	return
}

func ShowRelationWithBank() (result []models.RelationWithBank) {
	var relationBank []models.RelationWithBank

	connection.DB.Where("status = ?", "L").Find(&relationBank)

	return relationBank
}

func ShowRekeningDebitur(generalInformationId string) (result []models.ShowRekeningDebitur) {
	connection.DB.Select("rd.*, od.name, od.no_identity, od.npwp, od.key_person, CASE WHEN od.key_person = TRUE THEN 'Yes' ELSE 'No' END as pemilik").
		Table("data_rekening_debiturs AS rd").
		Joins("JOIN ownership_data AS od ON od.id = rd.ownership_data_id").
		Where("rd.status = ? AND od.status = ? AND rd.general_information_id = ?", "L", "L", generalInformationId).
		Order("od.key_person DESC").
		Find(&result)

	return
}

func ShowCustomerLoanInfo(generalInformationId string) (result []models.CustomerLoanInfo) {

	return
}

func ShowFacilityNo() (result []models.Dropdown) {
	connection.DB.Select("id, code AS name").
		Table("credit_types").
		Where("status = ?", "L").Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowProduct() (result []models.Dropdown) {
	connection.DB.Select("id, CONCAT(code, ' - ', name) AS name").
		Table("credit_types").
		Where("status = ?", "L").Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowCustomerAA(generalInformationId string) (result []models.Dropdown) {
	connection.DB.Select("id, aa_no AS name").
		Table("customer_loan_infos").
		Where("status = ? AND general_information_id = ?", "L", generalInformationId).Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}
