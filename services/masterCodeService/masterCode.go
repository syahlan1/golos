package masterCodeService

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
)

func ShowDetailMasterCode(groupId, groupName string) (result []models.MasterCode, err error) {
	// Periksa apakah groupId tidak kosong
	if groupId != "" {
		// Konversi groupId ke tipe data integer
		groupIdInt, err := strconv.Atoi(groupId)
		if err != nil {
			return result, errors.New("invalid group ID")
		}

		// Gunakan groupIdInt dalam kondisi Where
		connection.DB.Where("code_group_id = ?", groupIdInt).Find(&result)
	} else if groupName != "" {
		// Jika groupId kosong, gunakan groupName
		connection.DB.Where("code_group = ?", groupName).Find(&result)
	} else {
		return result, errors.New("missing parameter")
	}

	return result, nil
}

func ShowMasterCode() (result []models.MasterCode) {
	var masterCode []models.MasterCode

	connection.DB.Find(&masterCode)

	return masterCode
}

func ShowMasterCodeGroup() (result []models.MasterCodeGroup) {
	var masterCodeGroup []models.MasterCodeGroup

	connection.DB.Find(&masterCodeGroup)

	return masterCodeGroup
}

func CreateMasterCode(claims string, data models.MasterCode) (err error) {

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return errors.New("user not found")
	}

	data.CreatedBy = user.Username

	// newMasterCode := models.MasterCode{
	// 	// Authoriser:         user.Username,
	// 	// AuthorizeDate:      timeNow,
	// 	// CreatedBy:          user.Username,
	// 	// CreatedDate:        timeNow,
	// 	Code:               data.Code,
	// 	CodeGroupId:        data.CodeGroupId,
	// 	// Status:             "L",
	// 	Description:        data.Description,
	// 	EnglishDescription: data.EnglishDescription,
	// 	Sequence:           data.Sequence,
	// 	CodeGroup:          data.CodeGroup,
	// }

	if err := connection.DB.Create(&data).Error; err != nil {
		return errors.New("failed to create Master Code")
	}

	return nil
}

func CreateMasterCodeGroup(claims string, data models.MasterCodeGroup) (err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return err
	}

	//check existing name code group
	var existingCodeGroup models.MasterCodeGroup
	if err := connection.DB.Where("code_group = ?", data.CodeGroup).First(&existingCodeGroup).Error; err == nil {
		return errors.New("code Group Name Already Exists")
	}

	data.CreatedBy = user.Username

	// newMasterCodeGroup := models.MasterCodeGroup{
	// 	// Authoriser:         user.Username,
	// 	// AuthorizeDate:      timeNow,
	// 	// CreatedBy:          user.Username,
	// 	// CreatedDate:        timeNow,
	// 	// Status:             "L",
	// 	CodeGroup:          data.CodeGroup,
	// 	Description:        data.Description,
	// 	EnglishDescription: data.EnglishDescription,
	// }

	if err := connection.DB.Create(&data).Error; err != nil {
		return errors.New("failed to Master Code Group")
	}

	return nil
}

func UpdateMasterCode(claims, masterCodeId string, updatedMasterCode models.MasterCode) (result models.MasterCode, err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return result, err
	}

	var masterCode models.MasterCode
	if err := connection.DB.First(&masterCode, masterCodeId).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	masterCode.UpdatedBy = user.Username
	masterCode.UpdatedAt = time.Now()
	masterCode.Code = updatedMasterCode.Code
	// masterCode.CodeGroupId = updatedMasterCode.CodeGroupId
	masterCode.Description = updatedMasterCode.Description
	masterCode.EnglishDescription = updatedMasterCode.EnglishDescription
	masterCode.Sequence = updatedMasterCode.Sequence
	masterCode.CodeGroup = updatedMasterCode.CodeGroup

	if err := connection.DB.Save(&masterCode).Error; err != nil {
		return result, errors.New("failed to update Master Code")
	}

	return masterCode, nil
}

func UpdateMasterCodeGroup(claims string, masterCodeGroupId string, updatedMasterCodeGroup models.MasterCodeGroup) (result models.MasterCodeGroup, err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return result, err
	}

	var masterCodeGroup models.MasterCodeGroup
	if err := connection.DB.First(&masterCodeGroup, masterCodeGroupId).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	masterCodeGroup.UpdatedBy = user.Username
	masterCodeGroup.UpdatedAt = time.Now()
	masterCodeGroup.CodeGroup = updatedMasterCodeGroup.CodeGroup
	masterCodeGroup.Description = updatedMasterCodeGroup.Description
	masterCodeGroup.EnglishDescription = updatedMasterCodeGroup.EnglishDescription

	if err := connection.DB.Save(&masterCodeGroup).Error; err != nil {
		return result, errors.New("failed to update Master Code Group")
	}

	return masterCodeGroup, nil
}

func DeleteMasterCode(claims, masterCodeId string) (err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return err
	}

	var masterCode models.MasterCode

	masterCode.ModelMasterForm = utils.SoftDelete(user.Username)
	return connection.DB.Model(&masterCode).Where("id = ?", masterCodeId).Updates(&masterCode).Error
}

func DeleteMasterCodeGroup(claims string, masterCodeGroupId string) (err error) {
	// var user models.Users
	// if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
	// 	log.Println("Error retrieving user:", err)
	// 	return result, err
	// }

	// var masterCodeGroup models.MasterCodeGroup
	// if err := connection.DB.First(&masterCodeGroup, masterCodeGroupId).Error; err != nil {
	// 	return result,errors.New("Data Not Found")
	// }

	// masterCodeGroup.UpdatedBy = user.Username
	// masterCodeGroup.UpdatedAt = time.Now()
	// masterCodeGroup.Status = "D"

	// if err := connection.DB.Save(&masterCodeGroup).Error; err != nil {
	// 	return result,errors.New("Failed to delete Master Code Group")
	// }

	return connection.DB.Where("id = ?", masterCodeGroupId).Delete(&models.MasterCodeGroup{}).Error
}
