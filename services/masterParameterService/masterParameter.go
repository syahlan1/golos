package masterParameterService

import (
	"errors"
	"log"
	"time"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
)

func CreateParameter(claims string, data models.MasterParameter) (err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	if data.IsEncrypted == 1 {
		hashed, err := utils.Encrypt([]byte(data.ParamValue))
		if err != nil {
			return err
		}

		data.ParamValue = string(hashed)
	}
	data.CreatedBy = user.Username

	if err := connection.DB.Create(&data).Error; err != nil {
		return errors.New("failed to create Parameter")
	}

	return nil
}

func ShowAllParameter() (result []models.ShowAllMasterParameter, err error) {

	rows, err := connection.DB.Raw(`
	select 0 as module_id, 'Default' as module_name 
	union
	select id as module_id, module_name from master_modules
	where is_active is true and deleted_at is null
	order by module_id`).Rows()
	if err != nil {
		return result, err
	}

	defer rows.Close()

	var data models.ShowAllMasterParameter
	for rows.Next() {
		if err := connection.DB.ScanRows(rows, &data); err != nil {
			return result, err
		}
		var param []models.MasterParameter
		if err := connection.DB.Debug().Where("module_id = ?", data.ModuleId).Find(&param, "is_encrypted <> ?", 1).Error; err != nil {
			return result, err
		}

		data.Parameter = param
		result = append(result, data)
	}

	return
}

func ShowParameterDetail(parameterId string) (result models.MasterParameter, err error) {

	if err := connection.DB.
		Select("master_parameters.*, COALESCE(md.module_name, 'Default') as module_name").
		Joins("LEFT JOIN master_modules md ON md.id = master_parameters.module_id").
		Where("master_parameters.id = ?", parameterId).First(&result).Error; err != nil {
		return result, err
	}

	if result.IsEncrypted == 1 {
		hashed, err := utils.Decrypt(result.ParamValue)
		if err != nil {
			return result, err
		}
		result.ParamValue = string(hashed)
	}

	return
}

func UpdateMasterParameter(claims string, parameterId string, updatedMasterParameter models.MasterParameter) (result models.MasterParameter, err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return result, err
	}

	var masterParameter models.MasterParameter
	if err := connection.DB.First(&masterParameter, parameterId).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	if updatedMasterParameter.IsEncrypted == 1 {
		hashed, err := utils.Encrypt([]byte(updatedMasterParameter.ParamValue))
		if err != nil {
			return result, err
		}

		updatedMasterParameter.ParamValue = string(hashed)
	}

	masterParameter.UpdatedBy = user.Username
	masterParameter.UpdatedAt = time.Now()
	masterParameter.Description = updatedMasterParameter.Description
	masterParameter.EnglishDescription = updatedMasterParameter.EnglishDescription
	masterParameter.ParamKey = updatedMasterParameter.ParamKey
	masterParameter.ParamValue = updatedMasterParameter.ParamValue

	if err := connection.DB.Save(&masterParameter).Error; err != nil {
		return result, errors.New("failed to update Master Parameter")
	}

	return masterParameter, nil
}

func DeleteMasterParameter(claims, parameterId string) (err error) {

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return err
	}

	var masterParameter models.MasterParameter

	masterParameter.ModelMasterForm = utils.SoftDelete(user.Username)
	return connection.DB.Model(&masterParameter).Where("id = ?", parameterId).Updates(&masterParameter).Error
}
