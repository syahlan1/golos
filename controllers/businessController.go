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
	"github.com/syahlan1/golos/utils"
)

func ShowBusinessApplicant(c *fiber.Ctx) error {
	var business []models.Business
	var applicant []models.Applicant
	connection.DB.Find(&business)
	connection.DB.Find(&applicant)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data: fiber.Map{
			"business":   business,
			"applicants": applicant},
	})
}

func BusinessShow(c *fiber.Ctx) error {
	var businesses []models.Business

	connection.DB.Find(&businesses)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    businesses,
	})
}

func BusinessShowDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	var business models.Business

	connection.DB.Where("id = ?", id).Find(&business)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    business,
	})
}

func BusinessCreate(c *fiber.Ctx) error {
	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	document := models.Document{
		DocumentFile:      utils.GetStringValue(data, "document_file"),
		DocumentPath:      utils.GetStringValue(data, "document_path"),
		Status:            utils.GetStringValue(data, "status"),
		NoCreditSalesForm: utils.GetStringValue(data, "no_credit_sales_form"),
		DateOfLetter:      utils.GetStringValue(data, "date_of_letter"),
		DateOfReceipt:     utils.GetStringValue(data, "date_of_receipt"),
	}

	generalIformation := models.GeneralInformation{
		BankName:              utils.GetStringValue(data, "bank_name"),
		KCP:                   utils.GetStringValue(data, "kcp"),
		SubProgram:            utils.GetStringValue(data, "sub_program"),
		Analisis:              utils.GetStringValue(data, "analisis"),
		CabangPencairan:       utils.GetStringValue(data, "cabang_pencairan"),
		CabangAdmin:           utils.GetStringValue(data, "cabang_admin"),
		TglAplikasi:           utils.GetStringValue(data, "tgl_aplikasi"),
		TglPenerusan:          utils.GetStringValue(data, "tgl_penerusan"),
		Segmen:                utils.GetStringValue(data, "segmen"),
		NoAplikasi:            utils.GetIntValue(data, "no_aplikasi"),
		MarketInterestRate:    utils.GetIntValue(data, "masket_interest_rate"),
		RequestedInterestRate: utils.GetIntValue(data, "requested_interest_rate"),
		Status:                "L",
	}

	if err := connection.DB.Create(&generalIformation).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}) 
	}
	if err := connection.DB.Create(&document).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	companyFirstName := utils.GetStringValue(data, "company_first_name")
	companyName := utils.GetStringValue(data, "company_name")

	// Buat objek bisnis dengan nilai-nilai yang diberikan
	business := models.Business{
		Cif:                   utils.GetStringValue(data, "cif"),
		CompanyFirstName:      utils.GetStringValue(data, "company_first_name"),
		CompanyName:           utils.GetStringValue(data, "company_name"),
		CompanyType:           utils.GetStringValue(data, "company_type"),
		CustomerName:          companyFirstName + ". " + companyName,
		EstablishDate:         utils.GetStringValue(data, "establishment_date"),
		EstablishPlace:        utils.GetStringValue(data, "establish_place"),
		CompanyAddress:        utils.GetStringValue(data, "company_address"),
		District:              utils.GetStringValue(data, "district"),
		City:                  utils.GetStringValue(data, "city"),
		ZipCode:               utils.GetStringValue(data, "zip_code"),
		AddressType:           utils.GetStringValue(data, "address_type"),
		EternalRatingCompany:  utils.GetStringValue(data, "eternal_rating_company"),
		RatingClass:           utils.GetStringValue(data, "rating_class"),
		RatingDate:            utils.GetStringValue(data, "rating_date"),
		ListingBursaCode:      utils.GetStringValue(data, "listing_bursa_code"),
		ListingBursaDate:      utils.GetStringValue(data, "listing_bursa_date"),
		BusinessType:          utils.GetStringValue(data, "business_type"),
		AktaPendirian:         utils.GetStringValue(data, "akta_pendirian"),
		TglTerbit:             utils.GetStringValue(data, "tgl_terbit"),
		AktaLastChange:        utils.GetStringValue(data, "akta_last_change"),
		LastChangeDate:        utils.GetStringValue(data, "last_change_date"),
		NotarisName:           utils.GetStringValue(data, "notaris_name"),
		JumlahKaryawan:        utils.GetIntValue(data, "jumlah_karyawan"),
		NoTelp:                utils.GetStringValue(data, "no_telp"),
		NoFax:                 utils.GetStringValue(data, "no_fax"),
		NPWP:                  utils.GetStringValue(data, "npwp"),
		TDP:                   utils.GetStringValue(data, "tdp"),
		TglPenerbitan:         utils.GetStringValue(data, "tgl_penerbitan"),
		TglJatuhTempo:         utils.GetStringValue(data, "tgl_jatuh_tempo"),
		ContactPerson:         utils.GetStringValue(data, "contact_person"),
		BankName:              utils.GetStringValue(data, "bank_name"),
		KCP:                   utils.GetStringValue(data, "kcp"),
		SubProgram:            utils.GetStringValue(data, "sub_program"),
		Analisis:              utils.GetStringValue(data, "analisis"),
		CabangPencairan:       utils.GetStringValue(data, "cabang_pencairan"),
		CabangAdmin:           utils.GetStringValue(data, "cabang_admin"),
		TglAplikasi:           utils.GetStringValue(data, "tgl_aplikasi"),
		TglPenerusan:          utils.GetStringValue(data, "tgl_penelusuran"),
		Segmen:                utils.GetStringValue(data, "segmen"),
		NoAplikasi:            utils.GetIntValue(data, "no_aplikasi"),
		MarketInterestRate:    utils.GetIntValue(data, "market_interest_rate"),
		RequestedInterestRate: utils.GetIntValue(data, "requested_interest_rate"),
		DocumentFile:          utils.GetStringValue(data, "document_file"),
		Status:                "L",
		DocumentId:            document.Id,
		GeneralInformationId:  generalIformation.Id,
	}

	// Buat data bisnis ke database
	if err := connection.DB.Create(&business).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	createdBy, err := TakeUsername(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
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
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
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
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
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
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Code:    fiber.StatusNotFound,
			Message: "Business not found",
		})
	}

	var updatedBusiness models.Business
	if err := c.BodyParser(&updatedBusiness); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request payload",
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
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to update the business data",
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    businesses,
	})
}

func BusinessDelete(c *fiber.Ctx) error {
	businessID := c.Params("id")

	var businesses models.Business
	if err := connection.DB.First(&businesses, businessID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Code:    fiber.StatusNotFound,
			Message: "Business not found",
		})
	}

	businesses.Status = "D"

	if err := connection.DB.Save(&businesses).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to delete the business data",
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    businesses,
	})
}

func BusinessApproveUpdate(c *fiber.Ctx) error {
	businessID := c.Params("id")

	var business models.Business
	if err := connection.DB.First(&business, businessID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Code:    fiber.StatusNotFound,
			Message: "Business not found",
		})
	}

	if err := connection.DB.Save(&business).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to update the business data",
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    business,
	})
}

func ShowCompanyFirstName(c *fiber.Ctx) error {
	var companyFirstNames []string

	if err := connection.DB.Model(&models.CompanyFirstName{}).Pluck("name", &companyFirstNames).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    companyFirstNames,
	})
}

func ShowCompanyType(c *fiber.Ctx) error {
	var companyType []string

	if err := connection.DB.Model(&models.CompanyType{}).Pluck("name", &companyType).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    companyType,
	})
}

func ShowBusinessAddressType(c *fiber.Ctx) error {
	var businessAddressType []string

	if err := connection.DB.Model(&models.BusinessAddressType{}).Pluck("name", &businessAddressType).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    businessAddressType,
	})
}

func ShowEternalRatingCompany(c *fiber.Ctx) error {
	var eternalRatingCompany []string

	if err := connection.DB.Model(&models.EternalRatingCompany{}).Pluck("name", &eternalRatingCompany).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    eternalRatingCompany,
	})
}

func ShowRatingClass(c *fiber.Ctx) error {
	var ratingClass []string

	if err := connection.DB.Model(&models.RatingClass{}).Pluck("name", &ratingClass).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    ratingClass,
	})
}

func ShowKodeBursa(c *fiber.Ctx) error {
	var kodeBursa []string

	if err := connection.DB.Model(&models.KodeBursa{}).Pluck("name", &kodeBursa).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    kodeBursa,
	})
}

func ShowBusinessType(c *fiber.Ctx) error {
	var businessType []string

	if err := connection.DB.Model(&models.BusinessType{}).Pluck("name", &businessType).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    businessType,
	})
}

// zipcode
func GetProvinces(c *fiber.Ctx) error {
	var provinces []string
	if err := connection.DB.Model(&models.ZipCode{}).Distinct("province").Pluck("province", &provinces).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    provinces,
	})
}

func GetCitiesByProvince(c *fiber.Ctx) error {
	province := c.Query("province")
	var cities []string
	if err := connection.DB.Model(&models.ZipCode{}).Distinct("city").Where("province = ?", province).Pluck("city", &cities).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    cities,
	})
}

func GetDistrictByCity(c *fiber.Ctx) error {
	city := c.Query("city")
	var district []string
	if err := connection.DB.Model(&models.ZipCode{}).Distinct("district").Where("city = ?", city).Pluck("district", &district).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    district,
	})
}

func GetSubdistrictByDistrict(c *fiber.Ctx) error {
	district := c.Query("district")
	var subdistrict []string
	if err := connection.DB.Model(&models.ZipCode{}).Distinct("subdistrict").Where("district = ?", district).Pluck("subdistrict", &subdistrict).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    subdistrict,
	})
}

func GetZipCodesBySubdistrict(c *fiber.Ctx) error {
	subdistrict := c.Query("subdistrict")
	var zipCodes []string
	if err := connection.DB.Model(&models.ZipCode{}).Where("subdistrict = ?", subdistrict).Pluck("zip_code", &zipCodes).First(&zipCodes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    zipCodes,
	})
}
