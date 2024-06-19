package idCardService

import (
	"errors"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func CreateIdCard(data *models.IdCard) (err error) {
	return connection.DB.Create(&data).Error
}

func ShowIdCardById(id int) (result models.IdCard, err error) {

	if err = connection.DB.First(&result, "id = ?", id).Error; err != nil {
		return
	}
	return
}

func UpdateIdCard(id int, data models.IdCard) (err error) {

	var idCard models.IdCard
	if err := connection.DB.First(&idCard, id).Error; err != nil {
		return errors.New("id card not found")
	}
	updatedIdCard := data

	idCard.IdCardIssuedDate = updatedIdCard.IdCardIssuedDate
	idCard.IdCard = updatedIdCard.IdCard
	idCard.IdCardExpireDate = updatedIdCard.IdCardExpireDate
	idCard.IdCardAddress = updatedIdCard.IdCardAddress
	idCard.IdCardDistrict = updatedIdCard.IdCardDistrict
	idCard.IdCardCity = updatedIdCard.IdCardCity
	idCard.IdCardZipCode = updatedIdCard.IdCardZipCode
	idCard.AddressType = updatedIdCard.AddressType

	if err := connection.DB.Save(&idCard).Error; err != nil {
		return errors.New("failed to update the id card data")
	}

	return
}
