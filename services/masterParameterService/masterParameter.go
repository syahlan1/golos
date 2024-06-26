package masterParameterService

import (
	"errors"
	"log"
	"time"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func CreateParameter(claims string, data models.CreateMasterParameter) (err error) {
	timeNow := time.Now()
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	masterParameter := models.MasterParameter{
		CreatedBy:          user.Username,
		CreatedDate:        timeNow,
		Status:             "L",
		Description:        data.Description,
		EnglishDescription: data.EnglishDescription,
		ParamKey:           data.ParamKey,
		ParamValue:         data.ParamValue,
	}

	if err := connection.DB.Create(&masterParameter).Error; err != nil {
		return errors.New("failed to create Parameter")
	}

	return nil
}

func ShowAllParameter() (result []models.MasterParameter, err error) {
	var masterParameters []models.MasterParameter

	if err := connection.DB.Where("status = ?", "L").Find(&masterParameters).Error; err != nil {
		return result, err
	}

	return masterParameters, nil
}

func ShowParameterDetail(parameterId string) (result models.MasterParameter, err error) {
	var masterParameter models.MasterParameter

	if err := connection.DB.Where("id = ?", parameterId).First(&masterParameter).Error; err != nil {
		return result, err
	}

	return masterParameter, nil
}

func UpdateMasterParameter(claims string, parameterId string, updatedMasterParameter models.MasterParameter) (result models.MasterParameter ,err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return result,  err
	}

	var masterParameter models.MasterParameter
	if err := connection.DB.First(&masterParameter, parameterId).Error; err != nil {
		return result ,errors.New("data Not Found")
	}

	masterParameter.UpdatedBy = user.Username
	masterParameter.UpdatedDate = time.Now()
	masterParameter.Description = updatedMasterParameter.Description
	masterParameter.EnglishDescription = updatedMasterParameter.EnglishDescription
	masterParameter.ParamKey = updatedMasterParameter.ParamKey
	masterParameter.ParamValue = updatedMasterParameter.ParamValue

	if err := connection.DB.Save(&masterParameter).Error; err != nil {
		return result ,errors.New("failed to update Master Parameter")
	}

	return masterParameter, nil
}

func DeleteMasterParameter(claims, parameterId string) (result models.MasterParameter, err error) {

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return result, err
	}

	var masterParameter models.MasterParameter
	if err := connection.DB.First(&masterParameter, parameterId).Error; err != nil {
		return result ,errors.New("data Not Found")
	}

	masterParameter.UpdatedBy = user.Username
	masterParameter.UpdatedDate = time.Now()
	masterParameter.Status = "D"

	if err := connection.DB.Save(&masterParameter).Error; err != nil {
		return result ,errors.New("Failed to Delete Master Parameter")
	}

	return masterParameter, nil
}
