package ownershipDataService

import (
	"errors"
	"strconv"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
)

func CreateOwnershipData(data models.CreateOwnershipData) (err error) {
	ownership := models.OwnershipData{
		GeneralInformationId: data.GeneralInformationId,
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

func CreateRelationWithBank(ownershipIdInt int, data models.CreateRelationWithBank) (err error) {
	relation := models.RelationWithBank{
		Giro:            data.Giro,
		Tabungan:        data.Tabungan,
		NoRekening:      data.NoRekening,
		Debitur:         data.Debitur,
		Status:          "L",
		OwnershipDataId: ownershipIdInt,
	}

	// debitur := models.DataRekeningDebitur{
	// 	NoIdCard:        data.NoIdCard,
	// 	NPWP:            data.NPWP,
	// 	KeyPerson:       data.KeyPerson,
	// 	NoRekening:      data.NoRekening,
	// 	Remark:          data.Remark,
	// 	Status:          "L",
	// 	OwnershipDataId: ownershipIdInt,
	// }

	if err := connection.DB.Create(&relation).Error; err != nil {
		return err
	}

	// if err := connection.DB.Create(&debitur).Error; err != nil {
	// 	return err
	// }

	return nil
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
		return result, errors.New("Data Not Found")
	}

	rekeningDebitur.Status = "D"

	if err := connection.DB.Save(&rekeningDebitur).Error; err != nil {
		return result, errors.New("Failed to delete")
	}

	return rekeningDebitur, nil
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
