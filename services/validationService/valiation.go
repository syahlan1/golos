package validationService

import (
	"errors"
	"log"
	"time"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func ShowAllValidations() (result []models.MasterValidation, err error) {
	var validations []models.MasterValidation

	if err := connection.DB.Where("status = ? AND is_active = ?", "L", 1).Find(&validations).Error; err != nil {
		return result, errors.New("failed to fetch validations")
	}
	return validations, nil
}

func ShowDetailValidation(validationId string) (result models.MasterValidation, err error) {
	var validations models.MasterValidation

	if err := connection.DB.Where("id = ?", validationId).Find(&validations).Error; err != nil {
		return result, errors.New("failed to fetch validations")
	}

	return validations, nil
}

func ShowValidationByColumn(columnId string) (result []models.MasterValidation, err error) {
	var validations []models.MasterValidation

	if err := connection.DB.Where("status = ? AND is_active = ? AND column_id = ?", "L", 1, columnId).Find(&validations).Error; err != nil {
		return result, errors.New("failed to fetch validations")
	}

	return validations, nil
}

func CreateValidation(claims string, columnId int, data models.CreateValidation) (err error) {

	timeNow := time.Now()

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	newMasterValidation := models.MasterValidation{
		CreatedBy:          user.Username,
		CreatedDate:        timeNow,
		Status:             "L",
		Description:        data.Description,
		EnglishDescription: data.EnglishDescription,
		MessageType:        data.MessageType,
		ValidationFunction: data.ValidationFunction,
		IsActive:           data.IsActive,
		MasterCodeId:       data.MasterCodeId,
		ColumnId:           columnId,
	}

	if err := connection.DB.Create(&newMasterValidation).Error; err != nil {
		return errors.New("failed to create Master Validation")
	}

	return nil
}

func UpdateValidation(claims, masterValidationId string, updatedMasterValidation models.MasterValidation) (result models.MasterValidation, err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return result, err
	}

	var masterValidation models.MasterValidation
	if err := connection.DB.First(&masterValidation, masterValidationId).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	masterValidation.UpdatedBy = user.Username
	masterValidation.UpdatedDate = time.Now()
	masterValidation.Description = updatedMasterValidation.Description
	masterValidation.EnglishDescription = updatedMasterValidation.EnglishDescription
	masterValidation.MessageType = updatedMasterValidation.MessageType
	masterValidation.ValidationFunction = updatedMasterValidation.ValidationFunction
	masterValidation.IsActive = updatedMasterValidation.IsActive

	if err := connection.DB.Save(&masterValidation).Error; err != nil {
		return result, errors.New("failed to update Master Validation")
	}

	return masterValidation, nil
}

func DeleteValidation(claims, masterValidateId string) (result models.MasterCode, err error){
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return result, err
	}

	var masterValidate models.MasterCode
	if err := connection.DB.First(&masterValidate, masterValidateId).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	masterValidate.UpdatedBy = user.Username
	masterValidate.UpdatedDate = time.Now()
	masterValidate.Status = "D"

	if err := connection.DB.Save(&masterValidate).Error; err != nil {
		return result, errors.New("failed to delete Master Validation")
	}

	return masterValidate,  nil
}
