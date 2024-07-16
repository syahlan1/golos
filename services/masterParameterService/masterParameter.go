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

	// masterParameter := models.MasterParameter{
	// 	CreatedBy:          user.Username,
	// 	CreatedDate:        timeNow,
	// 	Status:             "L",
	// 	Description:        data.Description,
	// 	EnglishDescription: data.EnglishDescription,
	// 	ParamKey:           data.ParamKey,
	// 	ParamValue:         data.ParamValue,
	// }

	if err := connection.DB.Create(&data).Error; err != nil {
		return errors.New("failed to create Parameter")
	}

	return nil
}

func ShowAllParameter() (result []models.ShowAllMasterParameter, err error) {

	rows, err := connection.DB.
		Select("md.id as module_id, COALESCE(md.module_name, 'Default') as module_name").
		Joins("LEFT JOIN master_modules md ON md.id = master_parameters.module_id AND md.is_active IS TRUE").
		Model(&models.MasterParameter{}).
		Group("md.id").
		Order("(case when md.id is null then 1 else 0 end) desc").
		Where("is_encrypted <> ? and md.deleted_at IS NULL", 1).Rows()
	if err != nil {
		return result, err
	}

	defer rows.Close()

	var data models.ShowAllMasterParameter
	for rows.Next() {
		if err := connection.DB.ScanRows(rows, &data); err != nil {
			return result, err
		}

		if data.ModuleId != nil {
			if err := connection.DB.Where("module_id = ?", *data.ModuleId).Find(&data.Parameter, "is_encrypted <> ?", 1).Error; err != nil {
				return result, err
			}
		} else {
			if err := connection.DB.Where("module_id IS NULL").Find(&data.Parameter, "is_encrypted <> ?", 1).Error; err != nil {
				return result, err
			}
		}

		result = append(result, data)
	}

	return
}

func ShowParameterDetail(parameterId string) (result models.MasterParameter, err error) {

	if err := connection.DB.Where("id = ?", parameterId).First(&result).Error; err != nil {
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
