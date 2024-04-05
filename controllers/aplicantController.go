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
		ZipCode:             data["zip_code"].(string),
		HomeStatus:          data["home_status"].(string),
		StaySince:           data["stay_since"].(string),
		NoTelp:              data["no_telp"].(string),
		NoFax:               data["no_fax"].(string),
		BirthPlace:          data["birth_place"].(string),
		BirthDate:           data["birth_date"].(string),
		MaritalStatus:       data["marital_status"].(string),
		Gender:              data["gender"].(string),
		Nationality:         data["nationality"].(string),
		NumberOfChildren:    int(data["number_of_children"].(float64)),
		NoKartuKeluarga:     data["no_kartu_keluarga"].(string),
		SpouseName:          data["spouse_name"].(string),
		SpouseIdCard:        data["spouse_id_card"].(string),
		SpouseAddress:       data["spouse_address"].(string),
		GroupNasabah:        data["group_nasabah"].(string),
		SektorEkonomi1:      data["sektor_ekonomi_1"].(string),
		SektorEkonomi2:      data["sektor_ekonomi_2"].(string),
		SektorEkonomi3:      data["sektor_ekonomi_3"].(string),
		SektorEkonomiOjk:    data["sektor_ekonomi_ojk"].(string),
		NetIncome:           int(data["net_income"].(float64)),
		LokasiPabrik:        data["lokasi_pabrik"].(string),
		KeyPerson:           data["key_person"].(string),
		LokasiDati2:         data["lokasi_dati_2"].(string),
		HubunganNasabahBank: data["hubungan_nasabah_bank"].(string),
		HubunganKeluarga:    data["hubungan_keluarga"].(string),
		IdCardIssuedDate:    data["id_card_issued_date"].(string),
		IdCard:              data["id_card"].(string),
		IdCardExpireDate:    data["id_card_expire_date"].(string),
		IdCardAddress:       data["id_card_address"].(string),
		IdCardDistrict:      data["id_card_district"].(string),
		IdCardCity:          data["id_card_city"].(string),
		IdCardZipCode:       data["id_card_zip_code"].(string),
		AddressType:         data["address_type"].(string),
		Education:           data["education"].(string),
		JobPosition:         data["job_position"].(string),
		BusinessSector:      data["business_sector"].(string),
		EstablishDate:       data["establish_date"].(string),
		NPWP:                data["npwp"].(string),
		GrossIncomePerMonth: int(data["gross_income_per_month"].(float64)),
		NumberOfEmployees:   int(data["number_of_employees"].(float64)),
		MotherName:          data["mother_name"].(string),
		NamaPelaporan:       data["nama_pelaporan"].(string),
		NegaraDomisili:      data["negara_domisili"].(string),
		NamaInstansi:        data["nama_instansi"].(string),
		KodeInstansi:        data["kode_instansi"].(string),
		NoPegawai:           data["no_pegawai"].(string),
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
	applicant.GroupNasabah = updatedApplicant.GroupNasabah
	applicant.SektorEkonomi1 = updatedApplicant.SektorEkonomi1
	applicant.SektorEkonomi2 = updatedApplicant.SektorEkonomi2
	applicant.SektorEkonomi3 = updatedApplicant.SektorEkonomi3
	applicant.SektorEkonomiOjk = updatedApplicant.SektorEkonomiOjk
	applicant.NetIncome = updatedApplicant.NetIncome
	applicant.LokasiPabrik = updatedApplicant.LokasiPabrik
	applicant.KeyPerson = updatedApplicant.KeyPerson
	applicant.LokasiDati2 = updatedApplicant.LokasiDati2
	applicant.HubunganNasabahBank = updatedApplicant.HubunganNasabahBank
	applicant.HubunganKeluarga = updatedApplicant.HubunganKeluarga
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

	if err := connection.DB.Model(&models.JobPosition{}).Pluck("name", &jobPosition).Error; err != nil {
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

func ShowSektorEkonomi(c *fiber.Ctx) error {
	var sektorEkonomi []string

	if err := connection.DB.Model(&models.SektorEkonomi{}).Pluck("name", &sektorEkonomi).Error; err != nil {
		return err
	}

	return c.JSON(sektorEkonomi)
}

func ShowHubunganNasabah(c *fiber.Ctx) error {
	var hubunganNasabah []string

	if err := connection.DB.Model(&models.HubunganNasabah{}).Pluck("name", &hubunganNasabah).Error; err != nil {
		return err
	}

	return c.JSON(hubunganNasabah)
}

func ShowHubunganKeluarga(c *fiber.Ctx) error {
	var hubunganKeluarga []string

	if err := connection.DB.Model(&models.HubunganKeluarga{}).Pluck("name", &hubunganKeluarga).Error; err != nil {
		return err
	}

	return c.JSON(hubunganKeluarga)
}

func ShowLokasiPabrik(c *fiber.Ctx) error {
	var lokasiPabrik []string

	if err := connection.DB.Model(&models.LokasiPabrik{}).Pluck("name", &lokasiPabrik).Error; err != nil {
		return err
	}

	return c.JSON(lokasiPabrik)
}

func ShowMaritalStatus(c *fiber.Ctx) error {
	var maritalStatus []string

	if err := connection.DB.Model(&models.MaritalStatus{}).Pluck("name", &maritalStatus).Error; err != nil {
		return err
	}

	return c.JSON(maritalStatus)
}
