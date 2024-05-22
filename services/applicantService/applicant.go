package applicantService

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func ApplicantShow() (result []models.ApplicantDetail, err error) {
	var applicants []models.Applicant

	// Find all applicants with status "L"
	if err = connection.DB.Where("status = ?", "L").Find(&applicants).Error; err != nil {
		return nil, err
	}

	// Load related data for each applicant
	for _, applicant := range applicants {
		applicantDetail := models.ApplicantDetail{
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

		result = append(result, applicantDetail)
	}

	return result, nil
}

func ApplicantShowDetail(applicantID string) (result models.ApplicantDetail, err error) {

	var applicant models.Applicant
	// Find the applicant
	if err = connection.DB.First(&applicant, "id = ? AND status = ?", applicantID, "L").Error; err != nil {
		return result, errors.New("applicant Not Found")
	}

	// Initialize the detail struct with applicant data
	applicantDetail := models.ApplicantDetail{
		Applicant: applicant,
	}

	// Load related IdCard
	if applicant.IdCard != 0 {
		if err := connection.DB.First(&applicantDetail.IdCard, "id = ?", applicant.IdCard).Error; err != nil {
			return result, errors.New("id Card Not Found")
		}
	}

	// Load related Document
	if applicant.DocumentId != 0 {
		if err := connection.DB.First(&applicantDetail.Document, "id = ?", applicant.DocumentId).Error; err != nil {
			return result, errors.New("document Not Found")
		}
	}

	// Load related Spouse
	if applicant.SpouseId != 0 {
		if err := connection.DB.First(&applicantDetail.SpouseData, "id = ?", applicant.SpouseId).Error; err != nil {
			return result, errors.New("spouse Not Found")
		}
	}

	// Load related GeneralInformation
	if applicant.GeneralInformationId != 0 {
		if err := connection.DB.First(&applicantDetail.GeneralInformation, "id = ?", applicant.GeneralInformationId).Error; err != nil {
			return result, errors.New("general Information Not Found")
		}
	}

	return applicantDetail, nil
}

func ApplicantCreate(data models.CreateApplicant, username string) (err error) {

	spouse := data.Spouse
	idCard := data.IdCard
	document := data.Document
	generalInformation := data.GeneralInformation
	generalInformation.Status = "L"

	if err = connection.DB.Create(&spouse).Error; err != nil {
		return err
	}
	if err = connection.DB.Create(&idCard).Error; err != nil {
		return err
	}
	if err = connection.DB.Create(&generalInformation).Error; err != nil {
		return err
	}
	if err = connection.DB.Create(&document).Error; err != nil {
		return err
	}

	applicant := data.Applicant
	applicant.IdCard = idCard.Id
	applicant.DocumentId = document.Id
	applicant.SpouseId = spouse.Id
	applicant.GeneralInformationId = generalInformation.Id
	applicant.Status = "L"

	// Buat data appplicant ke database
	if err := connection.DB.Create(&applicant).Error; err != nil {
		return err
	}

	//generate id
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	approval := models.Approval{
		Id:                id.String(),
		DisplayData:       "Data " + applicant.CustomerName,
		Data:              ApplicantToJson(applicant, spouse, idCard, document, generalInformation),
		ApprovalSettingID: 1,
		CurrentProcess:    7,
		ApprovalStatus:    "draft",
		CreatedDate:       time.Now(),
		CreatedBy:         username,
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
		Data:           ApplicantToJson(applicant, spouse, idCard, document, generalInformation),
	}
	if err := connection.DB.Create(&history).Error; err != nil {
		return err
	}

	return nil
}

func ApplicantUpdate(applicantID string, data models.CreateApplicant) (applicant models.Applicant, err error) {

	if err := connection.DB.First(&applicant, applicantID).Error; err != nil {
		return applicant, errors.New("applicant not found")
	}
	updatedApplicant := data.Applicant

	idCardId := applicant.IdCard
	var idCard models.IdCard
	if err := connection.DB.First(&idCard, idCardId).Error; err != nil {
		return applicant, errors.New("id card not found")
	}
	updatedIdCard := data.IdCard

	spouseId := applicant.SpouseId
	var spouse models.SpouseData
	if err := connection.DB.First(&spouse, spouseId).Error; err != nil {
		return applicant, errors.New("spouse not found")
	}
	updatedSpouse := data.Spouse

	documentId := applicant.DocumentId
	var document models.Document
	if err := connection.DB.First(&document, documentId).Error; err != nil {
		return applicant, errors.New("document not found")
	}
	updatedDocument := data.Document

	generalInformationId := applicant.GeneralInformationId
	var generalInformation models.GeneralInformation
	if err := connection.DB.First(&generalInformation, generalInformationId).Error; err != nil {
		return applicant, errors.New("general Information not found")
	}
	updatedGeneralInformation := data.GeneralInformation

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
		return applicant, errors.New("failed to update the applicant data")
	}

	if err := connection.DB.Save(&spouse).Error; err != nil {
		return applicant, errors.New("failed to update the spouse data")
	}

	if err := connection.DB.Save(&document).Error; err != nil {
		return applicant, errors.New("failed to update the document data")
	}

	if err := connection.DB.Save(&generalInformation).Error; err != nil {
		return applicant, errors.New("failed to update the general information data")
	}

	if err := connection.DB.Save(&idCard).Error; err != nil {
		return applicant, errors.New("failed to update the id card data")
	}

	return applicant, nil
}

func ApplicantDelete(applicantID string) (err error) {

	var applicant models.Applicant
	if err := connection.DB.First(&applicant, applicantID).Error; err != nil {
		return errors.New("applicant not found")
	}

	applicant.Status = "D"

	if err := connection.DB.Save(&applicant).Error; err != nil {
		return errors.New("failed to delete the applicant data")
	}

	return nil
}

func ShowHomeStatus() (result []string, err error) {
	if err := connection.DB.Model(&models.HomeStatus{}).Pluck("name", &result).Error; err != nil {
		return nil, errors.New("failed to get Home Status")
	}

	return result, nil
}

func ShowApplicantAddressType() (result []string, err error) {
	if err := connection.DB.Model(&models.ApplicantAddressType{}).Pluck("name", &result).Error; err != nil {
		return nil, errors.New("failed to get Applicant Address Type")
	}
	
	return result, nil
}

func ShowEducation() (result []string, err error) {
	if err := connection.DB.Model(&models.Education{}).Pluck("name", &result).Error; err != nil {
		return nil, errors.New("failed to get Education Data")
	}

	return result, nil
}

func ShowJobPosition() (result []string, err error) {
if err := connection.DB.Model(&models.JobPosition{}).Pluck("name", &result).Error; err != nil {
		return nil, errors.New("failed to get Job Position Data")
	}

	return result, nil
}

func ShowBusinessSector() (result []string, err error) {
	if err := connection.DB.Model(&models.BusinessSector{}).Pluck("name", &result).Error; err != nil {
		return nil, errors.New("failed to get Business Sector Data")
	}

	return result, nil
}

func ShowKodeInstansi() (result []string, err error) {
	if err := connection.DB.Model(&models.KodeInstansi{}).Pluck("name", &result).Error; err != nil {
		return nil, errors.New("failed to get Kode Instansi Data")
	}

	return result, nil
}

func ShowNegara() (result []string, err error) {
	if err := connection.DB.Model(&models.Negara{}).Pluck("name", &result).Error; err != nil {
		return nil, errors.New("failed to get Negara Data")
	}

	return result, nil
}

func ShowSektorEkonomi() (result []string, err error) {
	if err := connection.DB.Model(&models.SektorEkonomi{}).Pluck("name", &result).Error; err != nil {
		return nil, errors.New("failed to get Sektor Ekonomi Data")
	}

	return result, nil
}

func ShowHubunganNasabah() (result []string, err error) {
	if err := connection.DB.Model(&models.HubunganNasabah{}).Pluck("name", &result).Error; err != nil {
		return nil, errors.New("failed to get Hubungan Nasabah Data")
	}

	return result, nil
}

func ShowHubunganKeluarga() (result []string, err error) {
	if err := connection.DB.Model(&models.HubunganKeluarga{}).Pluck("name", &result).Error; err != nil {
		return nil, errors.New("failed to get Hubungan Keluarga Data")
	}

	return result, nil
}

func ShowLokasiPabrik() (result []string, err error) {
if err := connection.DB.Model(&models.LokasiPabrik{}).Pluck("name", &result).Error; err != nil {
		return nil, errors.New("failed to get Lokasi Pabrik Data")
	}

	return result, nil
}

func ShowMaritalStatus() (result []string, err error) {
	if err := connection.DB.Model(&models.MaritalStatus{}).Pluck("name", &result).Error; err != nil {
		return nil, errors.New("failed to get Marital Status Data")
	}
	
	return result, nil
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
