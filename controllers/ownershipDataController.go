package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
)

func CreateOwnershipData(c *fiber.Ctx) error {
	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	ownership := models.OwnershipData{
		Name:              utils.GetStringValue(data, "name"),
		NoIdentity:        utils.GetStringValue(data, "no_identity"),
		IdCardAddress:     utils.GetStringValue(data, "id_card_identity"),
		City:              utils.GetStringValue(data, "city"),
		ZipCode:           utils.GetStringValue(data, "zip_code"),
		HomeOwnership:     utils.GetStringValue(data, "home_ownership"),
		Remark:            utils.GetStringValue(data, "remark"),
		CifManager:        utils.GetStringValue(data, "cif_manager"),
		BirthDate:         utils.GetStringValue(data, "birth_date"),
		LastEducation:     utils.GetStringValue(data, "last_education"),
		NPWP:              utils.GetStringValue(data, "npwp"),
		JobTitle:          utils.GetStringValue(data, "job_title"),
		Experince:         utils.GetStringValue(data, "experience"),
		OwnershipMarket:   utils.GetIntValue(data, "ownership_market"),
		CitizenshipStatus: utils.GetStringValue(data, "citizenship_status"),
		Gender:            utils.GetStringValue(data, "gender"),
		MaritalStatus:     utils.GetStringValue(data, "marital_status"),
		NumberOfChildren:  utils.GetIntValue(data, "number_of_children"),
		StartDate:         utils.GetStringValue(data, "start_date"),
		KeyPerson:         utils.GetStringValue(data, "key_person"),
		Removed:           utils.GetStringValue(data, "removed"),
		Status:            "L",
	}

	if err := connection.DB.Create(&ownership).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "insert sukses",
	})
}

func ShowOwnershipData(c *fiber.Ctx) error {
	var ownershipData []models.OwnershipData

	connection.DB.Where("status = ?", "L").Find(&ownershipData)

	return c.JSON(ownershipData)
}

func EditOwnershipData(c *fiber.Ctx) error {
	id := c.Params("id")

	var ownershipData models.OwnershipData
	if err := connection.DB.First(&ownershipData, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Data Not Found",
		})
	}

	var updatedOwnershipData models.OwnershipData
	if err := c.BodyParser(&updatedOwnershipData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid Data",
		})
	}

	ownershipData.Name = updatedOwnershipData.Name
	ownershipData.NoIdentity = updatedOwnershipData.NoIdentity
	ownershipData.IdCardAddress = updatedOwnershipData.IdCardAddress
	ownershipData.City = updatedOwnershipData.City
	ownershipData.ZipCode = updatedOwnershipData.ZipCode
	ownershipData.HomeOwnership = updatedOwnershipData.HomeOwnership
	ownershipData.Remark = updatedOwnershipData.Remark
	ownershipData.CifManager = updatedOwnershipData.CifManager
	ownershipData.BirthDate = updatedOwnershipData.BirthDate
	ownershipData.LastEducation = updatedOwnershipData.LastEducation
	ownershipData.NPWP = updatedOwnershipData.NPWP
	ownershipData.JobTitle = updatedOwnershipData.JobTitle
	ownershipData.Experince = updatedOwnershipData.Experince
	ownershipData.OwnershipMarket = updatedOwnershipData.OwnershipMarket
	ownershipData.CitizenshipStatus = updatedOwnershipData.CitizenshipStatus
	ownershipData.Gender = updatedOwnershipData.Gender
	ownershipData.MaritalStatus = updatedOwnershipData.MaritalStatus
	ownershipData.NumberOfChildren = updatedOwnershipData.NumberOfChildren
	ownershipData.StartDate = updatedOwnershipData.StartDate
	ownershipData.KeyPerson = updatedOwnershipData.KeyPerson
	ownershipData.Removed = updatedOwnershipData.Removed
	ownershipData.Status = updatedOwnershipData.Status

	if err := connection.DB.Save(&ownershipData).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to update the user data",
		})
	}

	return c.JSON(fiber.Map{
		"data":   ownershipData,
		"status": "Updated!",
	})
}

func DeleteOwnershipData(c *fiber.Ctx) error {
	ownershipId := c.Params("id")

	var ownership models.OwnershipData
	if err := connection.DB.First(&ownership, ownershipId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Data Not Found",
		})
	}

	ownership.Status = "D"

	if err := connection.DB.Save(&ownership).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to delete",
		})
	}

	return c.JSON(fiber.Map{
		"data":   ownership,
		"status": "Deleted!",
	})

}

func CreateRelationWithBank(c *fiber.Ctx) error {
	ownershipId := c.Params("id")

	ownershipIdInt, err := strconv.Atoi(ownershipId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ownership ID"})
	}

	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	relation := models.RelationWithBank{
		Giro:            utils.GetStringValue(data, "giro"),
		Tabungan:        utils.GetStringValue(data, "tabungan"),
		NoRekening:      utils.GetIntValue(data, "no_rekening"),
		Debitur:         utils.GetStringValue(data, "debitur"),
		Status:          "L",
		OwnershipDataId: ownershipIdInt,
	}

	debitur := models.DataRekeningDebitur{
		Name:            utils.GetStringValue(data, "name"),
		NoIdCard:        utils.GetStringValue(data, "no_id_card"),
		NPWP:            utils.GetIntValue(data, "npwp"),
		KeyPerson:       utils.GetStringValue(data, "key_person"),
		NoRekening:      utils.GetIntValue(data, "no_rekening"),
		Remark:          utils.GetStringValue(data, "remark"),
		Status:          "L",
		OwnershipDataId: ownershipIdInt,
	}

	if err := connection.DB.Create(&relation).Error; err != nil {
		return err
	}

	if err := connection.DB.Create(&debitur).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "insert sukses",
	})
}

func UpdateRekeningDebitur(c *fiber.Ctx) error {
	id := c.Params("id")

	var rekeningDebitur models.DataRekeningDebitur
	if err := connection.DB.First(&rekeningDebitur, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Data Not Found",
		})
	}

	var updatedRekeningDebitur models.DataRekeningDebitur
	if err := c.BodyParser(&rekeningDebitur); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid Data",
		})
	}

	rekeningDebitur.Name = updatedRekeningDebitur.Name
	rekeningDebitur.NoIdCard = updatedRekeningDebitur.NoIdCard
	rekeningDebitur.NPWP = updatedRekeningDebitur.NPWP
	rekeningDebitur.KeyPerson = updatedRekeningDebitur.KeyPerson
	rekeningDebitur.NoRekening = updatedRekeningDebitur.NoRekening
	rekeningDebitur.Remark = updatedRekeningDebitur.Remark

	if err := connection.DB.Save(&rekeningDebitur).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to update the user data",
		})
	}

	return c.JSON(fiber.Map{
		"data":   rekeningDebitur,
		"status": "Updated!",
	})
}

func UpdateRelationWithBank(c *fiber.Ctx) error {
	id := c.Params("id")

	var relationBank models.RelationWithBank
	if err := connection.DB.First(&relationBank, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Data Not Found",
		})
	}

	var updateRelationBank models.RelationWithBank
	if err := c.BodyParser(&updateRelationBank); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid Data",
		})
	}

	relationBank.Giro = updateRelationBank.Giro
	relationBank.Tabungan = updateRelationBank.Tabungan
	relationBank.NoRekening = updateRelationBank.NoRekening
	relationBank.Debitur = updateRelationBank.Debitur

	if err := connection.DB.Save(&relationBank).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to update the user data",
		})
	}

	return c.JSON(fiber.Map{
		"data":   relationBank,
		"status": "Updated!",
	})
}

func DeleteRelationWithBank(c *fiber.Ctx) error {
	Id := c.Params("id")

	var relation models.RelationWithBank
	if err := connection.DB.First(&relation, Id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Data Not Found",
		})
	}

	relation.Status = "D"

	if err := connection.DB.Save(&relation).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to delete",
		})
	}

	return c.JSON(fiber.Map{
		"data":   relation,
		"status": "Deleted!",
	})

}

func DeleteRekeningDebitur(c *fiber.Ctx) error {
	Id := c.Params("id")

	var rekeningDebitur models.DataRekeningDebitur
	if err := connection.DB.First(&rekeningDebitur, Id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Data Not Found",
		})
	}

	rekeningDebitur.Status = "D"

	if err := connection.DB.Save(&rekeningDebitur).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to delete",
		})
	}

	return c.JSON(fiber.Map{
		"data":   rekeningDebitur,
		"status": "Deleted!",
	})

}

func ShowRelationWithBank(c *fiber.Ctx) error {
	var relationBank []models.RelationWithBank

	connection.DB.Where("status = ?", "L").Find(&relationBank)

	return c.JSON(relationBank)
}

func ShowRekeningDebitur(c *fiber.Ctx) error {
	var rekeningDebitur []models.OwnershipData

	connection.DB.Where("status = ?", "L").Find(&rekeningDebitur)

	return c.JSON(rekeningDebitur)
}
