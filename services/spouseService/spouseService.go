package spouseService

import (
	"errors"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func CreateSpouse(data *models.SpouseData) (err error) {
	return connection.DB.Create(&data).Error
}

func ShowSpouseById(id int) (result models.SpouseData, err error) {
	if err = connection.DB.First(&result, "id = ?", id).Error; err != nil {
		return 
	}
	return
}

func UpdateSpouse(id int, data models.SpouseData) (err error) {

	var spouse models.SpouseData
	if err := connection.DB.First(&spouse, id).Error; err != nil {
		return errors.New("spouse not found")
	}
	updatedSpouse := data

	spouse.SpouseName = updatedSpouse.SpouseName
	spouse.SpouseIdCard = updatedSpouse.SpouseIdCard
	spouse.SpouseAddress = updatedSpouse.SpouseAddress
	spouse.SpouseIdDate = updatedSpouse.SpouseIdDate

	if err := connection.DB.Save(&spouse).Error; err != nil {
		return errors.New("failed to update the spouse data")
	}

	return
}
