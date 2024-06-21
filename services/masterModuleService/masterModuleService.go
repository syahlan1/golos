package masterModuleService

import (
	"errors"
	"log"
	"time"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func CreateMasterModule(claims string, data models.MasterModule) (err error) {

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return err
	}

	data.CreatedBy = user.Username

	if err := connection.DB.Create(&data).Error; err != nil {
		return errors.New("failed to create Master Module")
	}

	return
}

func ShowAllMasterModule() (result []models.MasterModule) {
	connection.DB.Find(&result)

	return result
}

func ShowMasterModuleDetail(id string) (result models.MasterModule, err error) {
	if err := connection.DB.Where("id = ?", id).First(&result).Error; err != nil {
		return result, err
	}

	return
}

func UpdateMasterModule(claims string, masterModuleId string, data models.MasterModule) (result models.MasterModule, err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
	}

	var masterModule models.MasterModule
	if err := connection.DB.First(&masterModule, masterModuleId).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	masterModule.ModuleName = data.ModuleName
	masterModule.Description = data.Description
	masterModule.EnglishDescription = data.EnglishDescription
	masterModule.IsActive = data.IsActive
	masterModule.UseBranch = data.UseBranch
	masterModule.UsePeriod = data.UsePeriod
	masterModule.UpdatedBy = user.Username
	masterModule.UpdatedAt = time.Now()

	if err := connection.DB.Model(&result).Where("id = ?", masterModuleId).Save(&masterModule).Error; err != nil {
		return result, err
	}

	return masterModule, nil
}

func DeleteMasterModule(claims, masterModuleId string) (err error) {

	if err := connection.DB.Where("id = ?", masterModuleId).Delete(&models.MasterModule{}).Error; err != nil {
		return  err
	}

	return
}
