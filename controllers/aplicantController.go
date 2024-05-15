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

func ApplicantShow(c *fiber.Ctx) error {
	var applicant []models.Applicant

	connection.DB.Where("status = ?", "L").Find(&applicant)

	return c.JSON(applicant)
}

func ApplicantShowDetail(c *fiber.Ctx) error {
	var applicant models.Applicant

	connection.DB.Find(&applicant, c.Params("id"))

	return c.JSON(applicant)
}

func ApplicantCreate(c *fiber.Ctx) error {
	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Buat objek bisnis dengan nilai-nilai yang diberikan
	applicant := models.Applicant{
		TitleBeforeName:       getStringValue(data, "title_before_name"),
		CustomerName:          getStringValue(data, "customer_name"),
		TitleAfterName:        getStringValue(data, "title_after_name"),
		NickName:              getStringValue(data, "nickname"),
		HomeAddress:           getStringValue(data, "home_address"),
		District:              getStringValue(data, "district"),
		City:                  getStringValue(data, "city"),
		ZipCode:               getStringValue(data, "zip_code"),
		HomeStatus:            getStringValue(data, "home_status"),
		StaySince:             getStringValue(data, "stay_since"),
		NoTelp:                getStringValue(data, "no_telp"),
		NoFax:                 getStringValue(data, "no_fax"),
		BirthPlace:            getStringValue(data, "birth_place"),
		BirthDate:             getStringValue(data, "birth_date"),
		MaritalStatus:         getStringValue(data, "marital_status"),
		Gender:                getStringValue(data, "gender"),
		Nationality:           getStringValue(data, "nationality"),
		NumberOfChildren:      getIntValue(data, "number_of_children"),
		NoKartuKeluarga:       getStringValue(data, "no_kartu_keluarga"),
		SpouseName:            getStringValue(data, "spouse_name"),
		SpouseIdCard:          getStringValue(data, "spouse_id_card"),
		SpouseAddress:         getStringValue(data, "spouse_address"),
		GroupNasabah:          getStringValue(data, "group_nasabah"),
		SektorEkonomi1:        getStringValue(data, "sektor_ekonomi_1"),
		SektorEkonomi2:        getStringValue(data, "sektor_ekonomi_2"),
		SektorEkonomi3:        getStringValue(data, "sektor_ekonomi_3"),
		SektorEkonomiOjk:      getStringValue(data, "sektor_ekonomi_ojk"),
		NetIncome:             getIntValue(data, "net_income"),
		LokasiPabrik:          getStringValue(data, "lokasi_pabrik"),
		KeyPerson:             getStringValue(data, "key_person"),
		LokasiDati2:           getStringValue(data, "lokasi_dati_2"),
		HubunganNasabahBank:   getStringValue(data, "hubungan_nasabah_bank"),
		HubunganKeluarga:      getStringValue(data, "hubungan_keluarga"),
		IdCardIssuedDate:      getStringValue(data, "id_card_issued_date"),
		IdCard:                getStringValue(data, "id_card"),
		IdCardExpireDate:      getStringValue(data, "id_card_expire_date"),
		IdCardAddress:         getStringValue(data, "id_card_address"),
		IdCardDistrict:        getStringValue(data, "id_card_district"),
		IdCardCity:            getStringValue(data, "id_card_city"),
		IdCardZipCode:         getStringValue(data, "id_card_zip_code"),
		AddressType:           getStringValue(data, "address_type"),
		Education:             getStringValue(data, "education"),
		JobPosition:           getStringValue(data, "job_position"),
		BusinessSector:        getStringValue(data, "business_sector"),
		EstablishDate:         getStringValue(data, "establish_date"),
		NPWP:                  getStringValue(data, "npwp"),
		GrossIncomePerMonth:   getIntValue(data, "gross_income_per_month"),
		NumberOfEmployees:     getIntValue(data, "number_of_employees"),
		MotherName:            getStringValue(data, "mother_name"),
		NamaPelaporan:         getStringValue(data, "nama_pelaporan"),
		NegaraDomisili:        getStringValue(data, "negara_domisili"),
		NamaInstansi:          getStringValue(data, "nama_instansi"),
		KodeInstansi:          getStringValue(data, "kode_instansi"),
		NoPegawai:             getStringValue(data, "no_pegawai"),
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
	}

	// Buat data bisnis ke database
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
		Data:              ApplicantToJson(applicant),
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
		Data:           ApplicantToJson(applicant),
	}
	if err := connection.DB.Create(&history).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "insert sukses",
	})
}

func ApplicantToJson(applicant models.Applicant) string {
	jsonData, err := json.Marshal(applicant)
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
	applicant.BankName = updatedApplicant.BankName
	applicant.KCP = updatedApplicant.KCP
	applicant.SubProgram = updatedApplicant.SubProgram
	applicant.Analisis = updatedApplicant.Analisis
	applicant.CabangPencairan = updatedApplicant.CabangPencairan
	applicant.CabangAdmin = updatedApplicant.CabangAdmin
	applicant.TglAplikasi = updatedApplicant.TglAplikasi
	applicant.TglPenerusan = updatedApplicant.TglPenerusan
	applicant.Segmen = updatedApplicant.Segmen
	applicant.NoAplikasi = updatedApplicant.NoAplikasi
	applicant.MarketInterestRate = updatedApplicant.MarketInterestRate
	applicant.RequestedInterestRate = updatedApplicant.RequestedInterestRate
	applicant.DocumentFile = updatedApplicant.DocumentFile
	applicant.Status = updatedApplicant.Status

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
