package masterWorkflowService

import (
	"errors"
	"log"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
)

func CreateMasterWorkflow(claims string, data models.MasterWorkflow) (err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return err
	}

	data.CreatedBy = user.Username

	if err := connection.DB.Create(&data).Error; err != nil {
		return errors.New("failed to create Master Code")
	}

	// Return success response
	return nil
}

func ShowMasterWorkflow() (result []models.MasterWorkflow) {
	connection.DB.Find(&result)
	return
}

func UpdateMasterWorkflow(claims string, masterWorkflowId string, data models.MasterWorkflow) (result models.MasterWorkflow, err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
	}

	var masterWorkflow models.MasterWorkflow
	if err := connection.DB.First(&masterWorkflow, masterWorkflowId).Error; err != nil {
		return result, errors.New("data Not Found")
	}

	masterWorkflow.StatusDescription = data.StatusDescription
	masterWorkflow.StatusEnglishDescription = data.StatusEnglishDescription
	masterWorkflow.StatusName = data.StatusName
	masterWorkflow.UpdatedBy = user.Username

	if err := connection.DB.Save(&masterWorkflow).Error; err != nil {
		return result, err
	}

	return masterWorkflow, nil
}

func DeleteMasterWorkflow(claims, masterWorkflowId string) (err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return err
	}

	var masterWorkflow models.MasterWorkflow

	masterWorkflow.ModelMasterForm = utils.SoftDelete(user.Username)
	return connection.DB.Model(&masterWorkflow).Where("id = ?", masterWorkflowId).Updates(&masterWorkflow).Error
}
