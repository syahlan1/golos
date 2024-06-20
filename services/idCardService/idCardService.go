package idCardService

import (
	"errors"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
)

func ShowAddressType() (result []models.Dropdown) {
	connection.DB.Select("id, name").
		Model(&models.AddressType{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func CreateIdCard(data *models.IdCard) (err error) {
	return connection.DB.Create(&data).Error
}

func ShowIdCardById(id int) (result models.ShowIdCard, err error) {

	if err = connection.DB.Select("id_cards.*, at.name AS address_type").
		Joins("JOIN address_types at ON at.id = id_cards.address_type_id").
		Model(&models.IdCard{}).
		First(&result, "id_cards.id = ?", id).Error; err != nil {
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
	idCard.IdCardNo = updatedIdCard.IdCardNo
	idCard.IdCardExpireDate = updatedIdCard.IdCardExpireDate
	idCard.IdCardAddress = updatedIdCard.IdCardAddress
	idCard.IdCardDistrict = updatedIdCard.IdCardDistrict
	idCard.IdCardCity = updatedIdCard.IdCardCity
	idCard.IdCardZipCode = updatedIdCard.IdCardZipCode
	idCard.AddressTypeId = updatedIdCard.AddressTypeId

	if err := connection.DB.Save(&idCard).Error; err != nil {
		return errors.New("failed to update the id card data")
	}

	return
}
