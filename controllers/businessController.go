package controllers

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func ShowBusinessApplicant(c *fiber.Ctx) error {
	var business []models.Business
	var applicant []models.Applicant
	connection.DB.Find(&business)
	connection.DB.Find(&applicant)

	return c.JSON(fiber.Map{
		"business":  business,
		"applicant": applicant,
	})
}

func BusinessShow(c *fiber.Ctx) error {
	var businesses []models.Business

	connection.DB.Find(&businesses)

	return c.JSON(businesses)
}

func BusinessShowDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	var business models.Business

	connection.DB.Where("id = ?", id).Find(&business)

	return c.JSON(business)
}

func BusinessCreate(c *fiber.Ctx) error {
	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	document := models.Document{
		DocumentFile:      getStringValue(data, "document_file"),
		DocumentPath:      getStringValue(data, "document_path"),
		Status:            getStringValue(data, "status"),
		NoCreditSalesForm: getStringValue(data, "no_credit_sales_form"),
		DateOfLetter:      getStringValue(data, "date_of_letter"),
		DateOfReceipt:     getStringValue(data, "date_of_receipt"),
	}

	generalIformation := models.GeneralInformation{
		BankName:              getStringValue(data, "bank_name"),
		KCP:                   getStringValue(data, "kcp"),
		SubProgram:            getStringValue(data, "sub_program"),
		Analisis:              getStringValue(data, "analisis"),
		CabangPencairan:       getStringValue(data, "cabang_pencairan"),
		CabangAdmin:           getStringValue(data, "cabang_admin"),
		TglAplikasi:           getStringValue(data, "tgl_aplikasi"),
		TglPenerusan:          getStringValue(data, "tgl_penerusan"),
		Segmen:                getStringValue(data, "segmen"),
		NoAplikasi:            getIntValue(data, "no_aplikasi"),
		MarketInterestRate:    getIntValue(data, "masket_interest_rate"),
		RequestedInterestRate: getIntValue(data, "requested_interest_rate"),
		Status:                "L",
	}

	if err := connection.DB.Create(&generalIformation).Error; err != nil {
		return err
	}
	if err := connection.DB.Create(&document).Error; err != nil {
		return err
	}

	companyFirstName := getStringValue(data, "company_first_name")
	companyName := getStringValue(data, "company_type")

	// Buat objek bisnis dengan nilai-nilai yang diberikan
	business := models.Business{
		Cif:                   getStringValue(data, "cif"),
		CompanyFirstName:      getStringValue(data, "company_first_name"),
		CompanyName:           getStringValue(data, "company_name"),
		CompanyType:           getStringValue(data, "company_type"),
		CustomerName:          companyFirstName + ". " + companyName,
		EstablishDate:         getStringValue(data, "establishment_date"),
		EstablishPlace:        getStringValue(data, "establish_place"),
		CompanyAddress:        getStringValue(data, "company_address"),
		District:              getStringValue(data, "district"),
		City:                  getStringValue(data, "city"),
		ZipCode:               getStringValue(data, "zip_code"),
		AddressType:           getStringValue(data, "address_type"),
		EternalRatingCompany:  getStringValue(data, "eternal_rating_company"),
		RatingClass:           getStringValue(data, "rating_class"),
		RatingDate:            getStringValue(data, "rating_date"),
		ListingBursaCode:      getStringValue(data, "listing_bursa_code"),
		ListingBursaDate:      getStringValue(data, "listing_bursa_date"),
		BusinessType:          getStringValue(data, "business_type"),
		AktaPendirian:         getStringValue(data, "akta_pendirian"),
		TglTerbit:             getStringValue(data, "tgl_terbit"),
		AktaLastChange:        getStringValue(data, "akta_last_change"),
		LastChangeDate:        getStringValue(data, "last_change_date"),
		NotarisName:           getStringValue(data, "notaris_name"),
		JumlahKaryawan:        getIntValue(data, "jumlah_karyawan"),
		NoTelp:                getStringValue(data, "no_telp"),
		NoFax:                 getStringValue(data, "no_fax"),
		NPWP:                  getStringValue(data, "npwp"),
		TDP:                   getStringValue(data, "tdp"),
		TglPenerbitan:         getStringValue(data, "tgl_penerbitan"),
		TglJatuhTempo:         getStringValue(data, "tgl_jatuh_tempo"),
		ContactPerson:         getStringValue(data, "contact_person"),
		BankName:              getStringValue(data, "bank_name"),
		KCP:                   getStringValue(data, "kcp"),
		SubProgram:            getStringValue(data, "sub_program"),
		Analisis:              getStringValue(data, "analisis"),
		CabangPencairan:       getStringValue(data, "cabang_pencairan"),
		CabangAdmin:           getStringValue(data, "cabang_admin"),
		TglAplikasi:           getStringValue(data, "tgl_aplikasi"),
		TglPenerusan:          getStringValue(data, "tgl_penelusuran"),
		Segmen:                getStringValue(data, "segmen"),
		NoAplikasi:            getIntValue(data, "no_aplikasi"),
		MarketInterestRate:    getIntValue(data, "market_interest_rate"),
		RequestedInterestRate: getIntValue(data, "requested_interest_rate"),
		DocumentFile:          getStringValue(data, "document_file"),
		Status:                "L",
		DocumentId:            document.Id,
		GeneralInformationId:  generalIformation.Id,
	}

	// Buat data bisnis ke database
	if err := connection.DB.Create(&business).Error; err != nil {
		return err
	}

	createdBy, err := TakeUsername(c)
	if err != nil {
		return err
	}

	//generate id
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	approval := models.Approval{
		Id:                id.String(),
		DisplayData:       "Data badan usaha " + business.CompanyName,
		Data:              BusinessToJson(business, document, generalIformation),
		ApprovalSettingID: 1,
		CurrentProcess:    7,
		ApprovalStatus:    "draft",
		CreatedDate:       time.Now(),
		CreatedBy:         createdBy,
	}

	// Buat data approval ke database
	if err := connection.DB.Create(&approval).Error; err != nil {
		return err
	}

	historyId := ulid.MustNew(ulid.Timestamp(t), entropy)

	history := models.ApprovalHistory{
		Id:             historyId.String(),
		ApprovalID:     approval.Id,
		Date:           approval.CreatedDate,
		UserID:         approval.CreatedBy,
		Status:         approval.ApprovalStatus,
		CurrentProcess: approval.CurrentProcess,
		Data:           BusinessToJson(business, document, generalIformation),
	}
	if err := connection.DB.Create(&history).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "insert sukses",
	})
}

func BusinessToJson(business models.Business, document models.Document, generalInformation models.GeneralInformation) string {
	data := map[string]interface{}{
		"business":           business,
		"document":           document,
		"generalInformation": generalInformation,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error converting business data to JSON:", err)
		return "{}"
	}
	return string(jsonData)
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
	businesses.CompanyFirstName = updatedBusiness.CompanyFirstName
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
	businesses.BankName = updatedBusiness.BankName
	businesses.KCP = updatedBusiness.KCP
	businesses.SubProgram = updatedBusiness.SubProgram
	businesses.Analisis = updatedBusiness.Analisis
	businesses.CabangPencairan = updatedBusiness.CabangPencairan
	businesses.CabangAdmin = updatedBusiness.CabangAdmin
	businesses.TglAplikasi = updatedBusiness.TglAplikasi
	businesses.TglPenerusan = updatedBusiness.TglPenerusan
	businesses.Segmen = updatedBusiness.Segmen
	businesses.NoAplikasi = updatedBusiness.NoAplikasi
	businesses.MarketInterestRate = updatedBusiness.MarketInterestRate
	businesses.RequestedInterestRate = updatedBusiness.RequestedInterestRate
	businesses.DocumentFile = updatedBusiness.DocumentFile

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
	if err := connection.DB.First(&businesses, businessID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "User Not Found",
		})
	}

	businesses.Status = "D"

	if err := connection.DB.Save(&businesses).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to delete",
		})
	}

	return c.JSON(fiber.Map{
		"data":   businesses,
		"status": "Deleted!",
	})
}

func BusinessApproveUpdate(c *fiber.Ctx) error {
	businessID := c.Params("id")

	var business models.Business
	if err := connection.DB.First(&business, businessID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Id Not Found",
		})
	}

	if err := connection.DB.Save(&business).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to update the data",
		})
	}

	return c.JSON(fiber.Map{
		"data":   business,
		"status": "ApproveStatus berhasil diperbarui!",
	})
}

func ShowCompanyFirstName(c *fiber.Ctx) error {
	var companyFirstNames []string

	if err := connection.DB.Model(&models.CompanyFirstName{}).Pluck("name", &companyFirstNames).Error; err != nil {
		return err
	}

	return c.JSON(companyFirstNames)
}

func ShowCompanyType(c *fiber.Ctx) error {
	var companyType []string

	if err := connection.DB.Model(&models.CompanyType{}).Pluck("name", &companyType).Error; err != nil {
		return err
	}

	return c.JSON(companyType)
}

func ShowBusinessAddressType(c *fiber.Ctx) error {
	var businessAddressType []string

	if err := connection.DB.Model(&models.BusinessAddressType{}).Pluck("name", &businessAddressType).Error; err != nil {
		return err
	}

	return c.JSON(businessAddressType)
}

func ShowEternalRatingCompany(c *fiber.Ctx) error {
	var eternalRatingCompany []string

	if err := connection.DB.Model(&models.EternalRatingCompany{}).Pluck("name", &eternalRatingCompany).Error; err != nil {
		return err
	}

	return c.JSON(eternalRatingCompany)
}

func ShowRatingClass(c *fiber.Ctx) error {
	var ratingClass []string

	if err := connection.DB.Model(&models.RatingClass{}).Pluck("name", &ratingClass).Error; err != nil {
		return err
	}

	return c.JSON(ratingClass)
}

func ShowKodeBursa(c *fiber.Ctx) error {
	var kodeBursa []string

	if err := connection.DB.Model(&models.KodeBursa{}).Pluck("name", &kodeBursa).Error; err != nil {
		return err
	}

	return c.JSON(kodeBursa)
}

func ShowBusinessType(c *fiber.Ctx) error {
	var businessType []string

	if err := connection.DB.Model(&models.BusinessType{}).Pluck("name", &businessType).Error; err != nil {
		return err
	}

	return c.JSON(businessType)
}

// zipcode
func GetProvinces(c *fiber.Ctx) error {
	var provinces []string
	if err := connection.DB.Model(&models.ZipCode{}).Distinct("province").Pluck("province", &provinces).Error; err != nil {
		return err
	}
	return c.JSON(provinces)
}

func GetCitiesByProvince(c *fiber.Ctx) error {
	province := c.Query("province")
	var cities []string
	if err := connection.DB.Model(&models.ZipCode{}).Distinct("city").Where("province = ?", province).Pluck("city", &cities).Error; err != nil {
		return err
	}
	return c.JSON(cities)
}

func GetDistrictByCity(c *fiber.Ctx) error {
	city := c.Query("city")
	var district []string
	if err := connection.DB.Model(&models.ZipCode{}).Distinct("district").Where("city = ?", city).Pluck("district", &district).Error; err != nil {
		return err
	}
	return c.JSON(district)
}

func GetSubdistrictByDistrict(c *fiber.Ctx) error {
	district := c.Query("district")
	var subdistrict []string
	if err := connection.DB.Model(&models.ZipCode{}).Distinct("subdistrict").Where("district = ?", district).Pluck("subdistrict", &subdistrict).Error; err != nil {
		return err
	}
	return c.JSON(subdistrict)
}

func GetZipCodesBySubdistrict(c *fiber.Ctx) error {
	subdistrict := c.Query("subdistrict")
	var zipCodes []string
	if err := connection.DB.Model(&models.ZipCode{}).Where("subdistrict = ?", subdistrict).Pluck("zip_code", &zipCodes).First(&zipCodes).Error; err != nil {
		return err
	}
	return c.JSON(zipCodes)
}
