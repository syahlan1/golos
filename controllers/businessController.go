package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func BusinessShow(c *fiber.Ctx) error {
	var businesses []models.Business

	connection.DB.Find(&businesses)

	return c.JSON(businesses)
}

func BusinessCreate(c *fiber.Ctx) error {
	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Buat objek bisnis dengan nilai-nilai yang diberikan
	business := models.Business{
		Cif:                  int(data["cif"].(float64)),
		CompanyName:          data["company_name"].(string),
		CompanyType:          data["company_type"].(string),
		EstablishDate:        data["establishment_date"].(string),
		EstablishPlace:       data["establish_place"].(string),
		CompanyAddress:       data["company_address"].(string),
		District:             data["district"].(string),
		City:                 data["city"].(string),
		ZipCode:              int(data["zip_code"].(float64)),
		AddressType:          data["address_type"].(string),
		EternalRatingCompany: data["eternal_rating_company"].(string),
		RatingClass:          data["rating_class"].(string),
		RatingDate:           data["rating_date"].(string),
		ListingBursaCode:     int(data["listing_bursa_code"].(float64)),
		ListingBursaDate:     data["listing_bursa_date"].(string),
		BusinessType:         data["business_type"].(string),
		AktaPendirian:        data["akta_pendirian"].(string),
		TglTerbit:            data["tgl_terbit"].(string),
		AktaLastChange:       data["akta_last_change"].(string),
		LastChangeDate:       data["last_change_date"].(string),
		NotarisName:          data["notaris_name"].(string),
		JumlahKaryawan:       int(data["jumlah_karyawan"].(float64)),
		NoTelp:               int(data["no_telp"].(float64)),
		NoFax:                int(data["no_fax"].(float64)),
		NPWP:                 int(data["npwp"].(float64)),
		TDP:                  data["tdp"].(string),
		TglPenerbitan:        data["tgl_penerbitan"].(string),
		TglJatuhTempo:        data["tgl_jatuh_tempo"].(string),
		ContactPerson:        data["contact_person"].(string),
	}

	// Buat data bisnis ke database
	if err := connection.DB.Create(&business).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "insert sukses",
	})
}

func BusinessUpdate(c *fiber.Ctx) error {
	businessID := c.Params("id")

	var businesses models.Business
	if err := connection.DB.First(&businesses, businessID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "User Not Found",
		})
	}

	var updatedBusiness models.Business
	if err := c.BodyParser(&updatedBusiness); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid User Data",
		})
	}

	businesses.Cif = updatedBusiness.Cif
	businesses.CompanyName = updatedBusiness.CompanyName
	businesses.CompanyType = updatedBusiness.CompanyType
	businesses.EstablishDate = updatedBusiness.EstablishDate
	businesses.EstablishPlace = updatedBusiness.EstablishPlace
	businesses.CompanyAddress = updatedBusiness.CompanyAddress
	businesses.District = updatedBusiness.District
	businesses.City = updatedBusiness.City
	businesses.ZipCode = updatedBusiness.ZipCode
	businesses.AddressType = updatedBusiness.AddressType
	businesses.EternalRatingCompany = updatedBusiness.EternalRatingCompany
	businesses.RatingClass = updatedBusiness.RatingClass
	businesses.RatingDate = updatedBusiness.RatingDate
	businesses.ListingBursaCode = updatedBusiness.ListingBursaCode
	businesses.ListingBursaDate = updatedBusiness.ListingBursaDate
	businesses.BusinessType = updatedBusiness.BusinessType
	businesses.AktaPendirian = updatedBusiness.AktaPendirian
	businesses.TglTerbit = updatedBusiness.TglTerbit
	businesses.AktaLastChange = updatedBusiness.AktaLastChange
	businesses.LastChangeDate = updatedBusiness.LastChangeDate
	businesses.NotarisName = updatedBusiness.NotarisName
	businesses.JumlahKaryawan = updatedBusiness.JumlahKaryawan
	businesses.NoTelp = updatedBusiness.NoTelp
	businesses.NoFax = updatedBusiness.NoFax
	businesses.NPWP = updatedBusiness.NPWP
	businesses.TDP = updatedBusiness.TDP
	businesses.TglPenerbitan = updatedBusiness.TglPenerbitan
	businesses.TglJatuhTempo = updatedBusiness.TglJatuhTempo
	businesses.ContactPerson = updatedBusiness.ContactPerson

	if err := connection.DB.Save(&businesses).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to update the user data",
		})
	}

	return c.JSON(fiber.Map{
		"data":   businesses,
		"status": "Updated!",
	})
}

func BusinessDelete(c *fiber.Ctx) error {
	businessID := c.Params("id")

	var businesses models.Business

	if err := connection.DB.Delete(&businesses, businessID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Data Not Found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "item deleted",
	})
}
