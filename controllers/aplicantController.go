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

func ApplicantShow(c *fiber.Ctx) error {
	var applicants []models.Applicant

	// Find all applicants with status "L"
	if err := connection.DB.Where("status = ?", "L").Find(&applicants).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Error fetching applicants",
		})
	}

	// Struct to hold detailed applicant information
	type ApplicantDetail struct {
		models.Applicant
		SpouseData         models.SpouseData         `json:"spouse"`
		IdCard             models.IdCard             `json:"id_card"`
		Document           models.Document           `json:"document"`
		GeneralInformation models.GeneralInformation `json:"general_information"`
	}

	var detailedApplicants []ApplicantDetail

	// Load related data for each applicant
	for _, applicant := range applicants {
		applicantDetail := ApplicantDetail{
			Applicant: applicant,
		}

		// Load related IdCard
		if applicant.IdCard != 0 {
			connection.DB.First(&applicantDetail.IdCard, "id = ?", applicant.IdCard)
		}

		// Load related Document
		if applicant.DocumentId != 0 {
			connection.DB.First(&applicantDetail.Document, "id = ?", applicant.DocumentId)
		}

		// Load related Spouse
		if applicant.SpouseId != 0 {
			connection.DB.First(&applicantDetail.SpouseData, "id = ?", applicant.SpouseId)
		}

		// Load related GeneralInformation
		if applicant.GeneralInformationId != 0 {
			connection.DB.First(&applicantDetail.GeneralInformation, "id = ?", applicant.GeneralInformationId)
		}

		detailedApplicants = append(detailedApplicants, applicantDetail)
	}

	return c.JSON(detailedApplicants)
}

func ApplicantShowDetail(c *fiber.Ctx) error {
	applicantID := c.Params("id")

	// Buat struct untuk menampung data yang akan ditampilkan
	type ApplicantDetail struct {
		models.Applicant
		SpouseData         models.SpouseData         `json:"spouse"`
		IdCard             models.IdCard             `json:"id_card"`
		Document           models.Document           `json:"document"`
		GeneralInformation models.GeneralInformation `json:"general_information"`
	}

	var applicant models.Applicant
	// Find the applicant
	if err := connection.DB.First(&applicant, "id = ? AND status = ?", applicantID, "L").Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Applicant Not Found",
		})
	}

	// Initialize the detail struct with applicant data
	applicantDetail := ApplicantDetail{
		Applicant: applicant,
	}

	// Load related IdCard
	if applicant.IdCard != 0 {
		if err := connection.DB.First(&applicantDetail.IdCard, "id = ?", applicant.IdCard).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "IdCard Not Found",
			})
		}
	}

	// Load related Document
	if applicant.DocumentId != 0 {
		if err := connection.DB.First(&applicantDetail.Document, "id = ?", applicant.DocumentId).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "Document Not Found",
			})
		}
	}

	// Load related Spouse
	if applicant.SpouseId != 0 {
		if err := connection.DB.First(&applicantDetail.SpouseData, "id = ?", applicant.SpouseId).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "Spouse Not Found",
			})
		}
	}

	// Load related GeneralInformation
	if applicant.GeneralInformationId != 0 {
		if err := connection.DB.First(&applicantDetail.GeneralInformation, "id = ?", applicant.GeneralInformationId).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "General Information Not Found",
			})
		}
	}

	return c.JSON(applicantDetail)
}

func ApplicantCreate(c *fiber.Ctx) error {
	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Buat objek bisnis dengan nilai-nilai yang diberikan

	spouse := models.SpouseData{
		SpouseName:          utils.GetStringValue(data, "spouse_name"),
		SpouseIdCard:        utils.GetStringValue(data, "spouse_id_card"),
		SpouseAddress:       utils.GetStringValue(data, "spouse_address"),
		GroupNasabah:        utils.GetStringValue(data, "group_nasabah"),
		SektorEkonomi1:      utils.GetStringValue(data, "sektor_ekonomi_1"),
		SektorEkonomi2:      utils.GetStringValue(data, "sektor_ekonomi_2"),
		SektorEkonomi3:      utils.GetStringValue(data, "sektor_ekonomi_3"),
		SektorEkonomiOjk:    utils.GetStringValue(data, "sektor_ekonomi_ojk"),
		NetIncome:           utils.GetIntValue(data, "net_income"),
		LokasiPabrik:        utils.GetStringValue(data, "lokasi_pabrik"),
		KeyPerson:           utils.GetStringValue(data, "key_person"),
		LokasiDati2:         utils.GetStringValue(data, "lokasi_dati_2"),
		HubunganNasabahBank: utils.GetStringValue(data, "hubungan_nasabah_bank"),
		HubunganKeluarga:    utils.GetStringValue(data, "hubungan_keluarga"),
	}

	idCard := models.IdCard{
		IdCardIssuedDate: utils.GetStringValue(data, "id_card_issued_date"),
		IdCard:           utils.GetStringValue(data, "id_card"),
		IdCardExpireDate: utils.GetStringValue(data, "id_card_expire_date"),
		IdCardAddress:    utils.GetStringValue(data, "id_card_address"),
		IdCardDistrict:   utils.GetStringValue(data, "id_card_district"),
		IdCardCity:       utils.GetStringValue(data, "id_card_city"),
		IdCardZipCode:    utils.GetStringValue(data, "id_card_zip_code"),
		AddressType:      utils.GetStringValue(data, "address_type"),
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

	if err := connection.DB.Create(&spouse).Error; err != nil {
		return err
	}
	if err := connection.DB.Create(&idCard).Error; err != nil {
		return err
	}
	if err := connection.DB.Create(&generalIformation).Error; err != nil {
		return err
	}
	if err := connection.DB.Create(&document).Error; err != nil {
		return err
	}

	applicant := models.Applicant{
		TitleBeforeName:      utils.GetStringValue(data, "title_before_name"),
		CustomerName:         utils.GetStringValue(data, "customer_name"),
		TitleAfterName:       utils.GetStringValue(data, "title_after_name"),
		NickName:             utils.GetStringValue(data, "nickname"),
		HomeAddress:          utils.GetStringValue(data, "home_address"),
		District:             utils.GetStringValue(data, "district"),
		City:                 utils.GetStringValue(data, "city"),
		ZipCode:              utils.GetStringValue(data, "zip_code"),
		HomeStatus:           utils.GetStringValue(data, "home_status"),
		StaySince:            utils.GetStringValue(data, "stay_since"),
		NoTelp:               utils.GetStringValue(data, "no_telp"),
		NoFax:                utils.GetStringValue(data, "no_fax"),
		BirthPlace:           utils.GetStringValue(data, "birth_place"),
		BirthDate:            utils.GetStringValue(data, "birth_date"),
		MaritalStatus:        utils.GetStringValue(data, "marital_status"),
		Gender:               utils.GetStringValue(data, "gender"),
		Nationality:          utils.GetStringValue(data, "nationality"),
		NumberOfChildren:     utils.GetIntValue(data, "number_of_children"),
		NoKartuKeluarga:      utils.GetStringValue(data, "no_kartu_keluarga"),
		Education:            utils.GetStringValue(data, "education"),
		JobPosition:          utils.GetStringValue(data, "job_position"),
		BusinessSector:       utils.GetStringValue(data, "business_sector"),
		EstablishDate:        utils.GetStringValue(data, "establish_date"),
		NPWP:                 utils.GetStringValue(data, "npwp"),
		GrossIncomePerMonth:  utils.GetIntValue(data, "gross_income_per_month"),
		NumberOfEmployees:    utils.GetIntValue(data, "number_of_employees"),
		MotherName:           utils.GetStringValue(data, "mother_name"),
		NamaPelaporan:        utils.GetStringValue(data, "nama_pelaporan"),
		NegaraDomisili:       utils.GetStringValue(data, "negara_domisili"),
		NamaInstansi:         utils.GetStringValue(data, "nama_instansi"),
		KodeInstansi:         utils.GetStringValue(data, "kode_instansi"),
		NoPegawai:            utils.GetStringValue(data, "no_pegawai"),
		IdCard:               idCard.Id,
		DocumentId:           document.Id,
		SpouseId:             spouse.Id,
		GeneralInformationId: generalIformation.Id,
		Status:               "L",
	}

	// Buat data appplicant ke database
	if err := connection.DB.Create(&applicant).Error; err != nil {
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
		DisplayData:       "Data " + applicant.CustomerName,
		Data:              ApplicantToJson(applicant, spouse, idCard, document, generalIformation),
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
		Data:           ApplicantToJson(applicant, spouse, idCard, document, generalIformation),
	}
	if err := connection.DB.Create(&history).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "insert sukses",
	})
}

func ApplicantToJson(applicant models.Applicant, spouse models.SpouseData, idCard models.IdCard, document models.Document, generalInformation models.GeneralInformation) string {
	data := map[string]interface{}{
		"applicant":          applicant,
		"spouse":             spouse,
		"idCard":             idCard,
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

func ApplicantUpdate(c *fiber.Ctx) error {
	applicantID := c.Params("id")

	var applicant models.Applicant
	if err := connection.DB.First(&applicant, applicantID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Applicant Not Found",
		})
	}

	var updatedApplicant models.Applicant
	if err := c.BodyParser(&updatedApplicant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid applicant Data",
		})
	}

	idCardId := applicant.IdCard

	var idCard models.IdCard
	if err := connection.DB.First(&idCard, idCardId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "ID Card Not Found",
		})
	}

	var updatedIdCard models.IdCard
	if err := c.BodyParser(&updatedIdCard); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid id card Data",
		})
	}

	spouseId := applicant.SpouseId

	var spouse models.SpouseData
	if err := connection.DB.First(&spouse, spouseId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Spouse Not Found",
		})
	}

	var updatedSpouse models.SpouseData
	if err := c.BodyParser(&updatedSpouse); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid Spouse Data",
		})
	}

	documentId := applicant.DocumentId

	var document models.Document
	if err := connection.DB.First(&document, documentId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Document Not Found",
		})
	}

	var updatedDocument models.Document
	if err := c.BodyParser(&updatedDocument); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid Spouse Data",
		})
	}

	generalInformationId := applicant.GeneralInformationId

	var generalInformation models.GeneralInformation
	if err := connection.DB.First(&generalInformation, generalInformationId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "General Information Not Found",
		})
	}

	var updatedGeneralInformation models.GeneralInformation
	if err := c.BodyParser(&updatedGeneralInformation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid General Information",
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

	spouse.SpouseName = updatedSpouse.SpouseName
	spouse.SpouseIdCard = updatedSpouse.SpouseIdCard
	spouse.SpouseAddress = updatedSpouse.SpouseAddress
	spouse.GroupNasabah = updatedSpouse.GroupNasabah
	spouse.SektorEkonomi1 = updatedSpouse.SektorEkonomi1
	spouse.SektorEkonomi2 = updatedSpouse.SektorEkonomi2
	spouse.SektorEkonomi3 = updatedSpouse.SektorEkonomi3
	spouse.SektorEkonomiOjk = updatedSpouse.SektorEkonomiOjk
	spouse.NetIncome = updatedSpouse.NetIncome
	spouse.LokasiPabrik = updatedSpouse.LokasiPabrik
	spouse.KeyPerson = updatedSpouse.KeyPerson
	spouse.LokasiDati2 = updatedSpouse.LokasiDati2
	spouse.HubunganNasabahBank = updatedSpouse.HubunganNasabahBank
	spouse.HubunganKeluarga = updatedSpouse.HubunganKeluarga

	idCard.IdCardIssuedDate = updatedIdCard.IdCardIssuedDate
	idCard.IdCard = updatedIdCard.IdCard
	idCard.IdCardExpireDate = updatedIdCard.IdCardExpireDate
	idCard.IdCardAddress = updatedIdCard.IdCardAddress
	idCard.IdCardDistrict = updatedIdCard.IdCardDistrict
	idCard.IdCardCity = updatedIdCard.IdCardCity
	idCard.IdCardZipCode = updatedIdCard.IdCardZipCode
	idCard.AddressType = updatedIdCard.AddressType

	document.DocumentFile = updatedDocument.DocumentFile
	document.DocumentPath = updatedDocument.DocumentPath
	document.Status = updatedDocument.Status
	document.NoCreditSalesForm = updatedDocument.NoCreditSalesForm
	document.DateOfLetter = updatedDocument.DateOfLetter
	document.DateOfReceipt = updatedDocument.DateOfReceipt

	generalInformation.BankName = updatedGeneralInformation.BankName
	generalInformation.KCP = updatedGeneralInformation.KCP
	generalInformation.SubProgram = updatedGeneralInformation.SubProgram
	generalInformation.Analisis = updatedGeneralInformation.Analisis
	generalInformation.CabangPencairan = updatedGeneralInformation.CabangPencairan
	generalInformation.CabangAdmin = updatedGeneralInformation.CabangAdmin
	generalInformation.TglAplikasi = updatedGeneralInformation.TglAplikasi
	generalInformation.TglPenerusan = updatedGeneralInformation.TglPenerusan
	generalInformation.Segmen = updatedGeneralInformation.Segmen
	generalInformation.NoAplikasi = updatedGeneralInformation.NoAplikasi
	generalInformation.MarketInterestRate = updatedGeneralInformation.MarketInterestRate
	generalInformation.RequestedInterestRate = updatedGeneralInformation.RequestedInterestRate

	if err := connection.DB.Save(&applicant).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to update the applicant data",
		})
	}

	if err := connection.DB.Save(&spouse).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to update the applicant data",
		})
	}

	if err := connection.DB.Save(&document).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to update the applicant data",
		})
	}

	if err := connection.DB.Save(&generalInformation).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to update the applicant data",
		})
	}

	if err := connection.DB.Save(&idCard).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to update the applicant data",
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
	if err := connection.DB.First(&applicant, applicantID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "User Not Found",
		})
	}

	applicant.Status = "D"

	if err := connection.DB.Save(&applicant).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to delete",
		})
	}

	return c.JSON(fiber.Map{
		"data":   applicant,
		"status": "Updated!",
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
