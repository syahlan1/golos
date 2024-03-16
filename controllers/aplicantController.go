package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func ApplicantShow(c *fiber.Ctx) error {
	var applicant []models.Applicant

	connection.DB.Find(&applicant)

	return c.JSON(applicant)
}

func ApplicantCreate(c *fiber.Ctx) error {
	// userRole, err := getUserRoleFromToken(c)
	// if err != nil {
	//     return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//         "error": err.Error(),
	//     })
	// }

	// // Cek apakah peran pengguna diizinkan untuk membuat entri
	// if userRole != "admin" {
	//     return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
	//         "error": "You do not have permission to access this resource",
	//     })
	// }

	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Buat objek bisnis dengan nilai-nilai yang diberikan
	applicant := models.Applicant{
		TitleBeforeName:     data["title_before_name"].(string),
		CustomerName:        data["customer_name"].(string),
		TitleAfterName:      data["title_after_name"].(string),
		NickName:            data["nickname"].(string),
		HomeAddress:         data["home_address"].(string),
		District:            data["district"].(string),
		City:                data["city"].(string),
		ZipCode:             int(data["zip_code"].(float64)),
		HomeStatus:          data["home_status"].(string),
		StaySince:           data["stay_since"].(string),
		NoTelp:              int(data["no_telp"].(float64)),
		NoFax:               int(data["no_fax"].(float64)),
		BirthPlace:          data["birth_place"].(string),
		BirthDate:           data["birth_date"].(string),
		MaritalStatus:       data["marital_status"].(string),
		Gender:              data["gender"].(string),
		Nationality:         data["nationality"].(string),
		NumberOfChildren:    int(data["number_of_children"].(float64)),
		NoKartuKeluarga:     int(data["no_kartu_keluarga"].(float64)),
		SpouseName:          data["spouse_name"].(string),
		SpouseIdCard:        int(data["spouse_id_card"].(float64)),
		SpouseAddress:       data["spouse_address"].(string),
		IdCardIssuedDate:    data["id_card_issued_date"].(string),
		IdCard:              int(data["id_card"].(float64)),
		IdCardExpireDate:    data["id_card_expire_date"].(string),
		IdCardAddress:       data["id_card_address"].(string),
		IdCardDistrict:      data["id_card_district"].(string),
		IdCardCity:          data["id_card_city"].(string),
		IdCardZipCode:       int(data["id_card_zip_code"].(float64)),
		AddressType:         data["address_type"].(string),
		Education:           data["education"].(string),
		JobPosition:         data["job_position"].(string),
		BusinessSector:      data["business_sector"].(string),
		EstablishDate:       data["establish_date"].(string),
		NPWP:                int(data["npwp"].(float64)),
		GrossIncomePerMonth: int(data["gross_income_per_month"].(float64)),
		NumberOfEmployees:   int(data["number_of_employees"].(float64)),
		MotherName:          data["mother_name"].(string),
		NamaPelaporan:       data["nama_pelaporan"].(string),
		NegaraDomisili:      data["negara_domisili"].(string),
		NamaInstansi:        data["nama_instansi"].(string),
		KodeInstansi:        data["kode_instansi"].(string),
		NoPegawai:           int(data["no_pegawai"].(float64)),
		ApproveStatus:       1,
	}

	// Buat data bisnis ke database
	if err := connection.DB.Create(&applicant).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "insert sukses",
	})
}

func ApplicantUpdate(c *fiber.Ctx) error {
	applicantID := c.Params("id")

	var applicant models.Applicant
	if err := connection.DB.First(&applicant, applicantID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "User Not Found",
		})
	}

	var updatedApplicant models.Applicant
	if err := c.BodyParser(&updatedApplicant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid User Data",
		})
	}
	applicant.TitleBeforeName = updatedApplicant.TitleBeforeName
	applicant.CustomerName = updatedApplicant.CustomerName
	applicant.TitleAfterName = updatedApplicant.TitleAfterName
	applicant.NickName = updatedApplicant.NickName
	applicant.HomeAddress = updatedApplicant.HomeAddress
	applicant.District = updatedApplicant.District
	applicant.City = updatedApplicant.City
	applicant.ZipCode = updatedApplicant.ZipCode
	applicant.HomeStatus = updatedApplicant.HomeStatus
	applicant.StaySince = updatedApplicant.StaySince
	applicant.NoTelp = updatedApplicant.NoTelp
	applicant.NoFax = updatedApplicant.NoFax
	applicant.BirthPlace = updatedApplicant.BirthPlace
	applicant.BirthDate = updatedApplicant.BirthDate
	applicant.MaritalStatus = updatedApplicant.MaritalStatus
	applicant.Gender = updatedApplicant.Gender
	applicant.Nationality = updatedApplicant.Nationality
	applicant.NumberOfChildren = updatedApplicant.NumberOfChildren
	applicant.NoKartuKeluarga = updatedApplicant.NoKartuKeluarga
	applicant.SpouseName = updatedApplicant.SpouseName
	applicant.SpouseIdCard = updatedApplicant.SpouseIdCard
	applicant.SpouseAddress = updatedApplicant.SpouseAddress
	applicant.IdCardIssuedDate = updatedApplicant.IdCardIssuedDate
	applicant.IdCard = updatedApplicant.IdCard
	applicant.IdCardExpireDate = updatedApplicant.IdCardExpireDate
	applicant.IdCardAddress = updatedApplicant.IdCardAddress
	applicant.IdCardDistrict = updatedApplicant.IdCardDistrict
	applicant.IdCardCity = updatedApplicant.IdCardCity
	applicant.IdCardZipCode = updatedApplicant.IdCardZipCode
	applicant.AddressType = updatedApplicant.AddressType
	applicant.Education = updatedApplicant.Education
	applicant.JobPosition = updatedApplicant.JobPosition
	applicant.BusinessSector = updatedApplicant.BusinessSector
	applicant.EstablishDate = updatedApplicant.EstablishDate
	applicant.NPWP = updatedApplicant.NPWP
	applicant.GrossIncomePerMonth = updatedApplicant.GrossIncomePerMonth
	applicant.NumberOfEmployees = updatedApplicant.NumberOfEmployees
	applicant.MotherName = updatedApplicant.MotherName
	applicant.NamaPelaporan = updatedApplicant.NamaPelaporan
	applicant.NegaraDomisili = updatedApplicant.NegaraDomisili
	applicant.NamaInstansi = updatedApplicant.NamaInstansi
	applicant.KodeInstansi = updatedApplicant.KodeInstansi
	applicant.NoPegawai = updatedApplicant.NoPegawai

	if err := connection.DB.Save(&applicant).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to update the user data",
		})
	}

	return c.JSON(fiber.Map{
		"data":   applicant,
		"status": "Updated!",
	})
}

func ApplicantDelete(c *fiber.Ctx) error {
	applicantID := c.Params("id")

	var applicant models.Applicant

	if err := connection.DB.Delete(&applicant, applicantID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Data Not Found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "item deleted",
	})
}

func ShowHomeStatus(c *fiber.Ctx) error {
	var HomeStatus []string

	if err := connection.DB.Model(&models.HomeStatus{}).Pluck("name", &HomeStatus).Error; err != nil {
		return err
	}

	return c.JSON(HomeStatus)
}

func ShowApplicantAddressType(c *fiber.Ctx) error {
	var applicantAddressType []string

	if err := connection.DB.Model(&models.ApplicantAddressType{}).Pluck("name", &applicantAddressType).Error; err != nil {
		return err
	}

	return c.JSON(applicantAddressType)
}

func ShowEducation(c *fiber.Ctx) error {
	var education []string

	if err := connection.DB.Model(&models.Education{}).Pluck("name", &education).Error; err != nil {
		return err
	}

	return c.JSON(education)
}

func ShowJobPosition(c *fiber.Ctx) error {
	var jobPosition []string

	if err := connection.DB.Model(&models.Education{}).Pluck("name", &jobPosition).Error; err != nil {
		return err
	}

	return c.JSON(jobPosition)
}

func ShowBusinessSector(c *fiber.Ctx) error {
	var businessSector []string

	if err := connection.DB.Model(&models.BusinessSector{}).Pluck("name", &businessSector).Error; err != nil {
		return err
	}

	return c.JSON(businessSector)
}

func ShowKodeInstansi(c *fiber.Ctx) error {
	var kodeInstansi []string

	if err := connection.DB.Model(&models.KodeInstansi{}).Pluck("name", &kodeInstansi).Error; err != nil {
		return err
	}

	return c.JSON(kodeInstansi)
}

func ShowNegara(c *fiber.Ctx) error {
	var negara []string

	if err := connection.DB.Model(&models.Negara{}).Pluck("name", &negara).Error; err != nil {
		return err
	}

	return c.JSON(negara)
}
