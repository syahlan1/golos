package masterModuleService

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
)

func CreateMasterModule(claims string, data models.MasterModule) (err error) {

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return err
	}

	var checkDatabaseName int64
	connection.DB.Model(&models.MasterModule{}).
		Where("database_name = ?", data.DatabaseName).
		Count(&checkDatabaseName)
	if checkDatabaseName != 0 {
		return fmt.Errorf("database_name : %s already exist", data.DatabaseName)
	}

	data.CreatedBy = user.Username

	if err := connection.DB.Create(&data).Error; err != nil {
		return errors.New("failed to create Master Module")
	}

	return
}

func ShowAllMasterModule(active bool) (result []models.MasterModule) {

	if active {
		connection.DB.Where("is_active = ?", true).Find(&result)
	} else {
		connection.DB.Find(&result)
	}

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

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return err
	}

	var masterModule models.MasterModule

	masterModule.ModelMasterForm = utils.SoftDelete(user.Username)
	return connection.DB.Model(&masterModule).Where("id = ?", masterModuleId).Updates(&masterModule).Error
}
