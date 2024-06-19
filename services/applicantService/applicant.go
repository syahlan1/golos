package applicantService

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"mime/multipart"
	"time"

	"github.com/oklog/ulid"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/documentService"
	"github.com/syahlan1/golos/services/generalInformationService"
	"github.com/syahlan1/golos/services/idCardService"
	"github.com/syahlan1/golos/services/sectorEconomyService"
	"github.com/syahlan1/golos/services/spouseService"
	"github.com/syahlan1/golos/utils"
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
			applicantDetail.IdCard, _ = idCardService.ShowIdCardById(applicant.IdCard)
		}

		// Load related Document
		if applicant.DocumentId != 0 {
			applicantDetail.Document,_ = documentService.ShowDocumentById(applicant.DocumentId)
		}

		// Load related Spouse
		if applicant.SpouseId != 0 {
			applicantDetail.SpouseData,_ = spouseService.ShowSpouseById(applicant.SpouseId)
		}

		// Load related GeneralInformation
		if applicant.GeneralInformationId != 0 {
			applicantDetail.GeneralInformation, _ = generalInformationService.ShowGeneralInformationById(applicant.GeneralInformationId)
		}

		// Load related SectorEconomy
		if applicant.SectorEconomyId != 0 {
			applicantDetail.SectorEconomy, _ = sectorEconomyService.ShowSectorEconomyById(applicant.SectorEconomyId)
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
		applicantDetail.IdCard, _ = idCardService.ShowIdCardById(applicant.IdCard)
	}

	// Load related Document
	if applicant.DocumentId != 0 {
		applicantDetail.Document,_ = documentService.ShowDocumentById(applicant.DocumentId)
	}

	// Load related Spouse
	if applicant.SpouseId != 0 {
		applicantDetail.SpouseData,_ = spouseService.ShowSpouseById(applicant.SpouseId)
	}

	// Load related GeneralInformation
	if applicant.GeneralInformationId != 0 {
		applicantDetail.GeneralInformation, _ = generalInformationService.ShowGeneralInformationById(applicant.GeneralInformationId)
	}

	// Load related SectorEconomy
	if applicant.SectorEconomyId != 0 {
		applicantDetail.SectorEconomy, _ = sectorEconomyService.ShowSectorEconomyById(applicant.SectorEconomyId)
	}

	return applicantDetail, nil
}

func ApplicantCreate(data models.CreateApplicant, username string) (err error) {

	spouse := data.Spouse
	idCard := data.IdCard
	document := data.Document
	generalInformation := data.GeneralInformation
	sectorEconomy := data.SectorEconomy

	if err = spouseService.CreateSpouse(&spouse); err != nil {
		return err
	}

	if err = idCardService.CreateIdCard(&idCard); err != nil {
		return err
	}

	if err = generalInformationService.CreateGeneralInformation(&generalInformation); err != nil {
		return err
	}

	if err = documentService.CreateDocument(&document); err != nil {
		return err
	}

	if err = sectorEconomyService.CreateSectorEconomy(&sectorEconomy); err != nil {
		return err
	}

	applicant := data.Applicant
	applicant.IdCard = idCard.Id
	applicant.DocumentId = document.Id
	applicant.SpouseId = spouse.Id
	applicant.GeneralInformationId = generalInformation.Id
	applicant.SectorEconomyId = sectorEconomy.Id
	applicant.Status = "L"

	// Buat data appplicant ke database
	if err := connection.DB.Create(&applicant).Error; err != nil {
		return err
	}

	showGeneralInformation, _ := generalInformationService.ShowGeneralInformationById(generalInformation.Id)
	showSectorEconomy, _ := sectorEconomyService.ShowSectorEconomyById(sectorEconomy.Id)

	//generate id
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	approval := models.Approval{
		Id:                id.String(),
		DisplayData:       "Data " + applicant.CustomerName,
		Data:              ApplicantToJson(applicant, spouse, idCard, document, showGeneralInformation, showSectorEconomy),
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
		Data:           ApplicantToJson(applicant, spouse, idCard, document, showGeneralInformation, showSectorEconomy),
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

	if err = idCardService.UpdateIdCard(applicant.IdCard, data.IdCard); err != nil {
		return applicant, err
	}

	if err = spouseService.UpdateSpouse(applicant.SpouseId, data.Spouse); err != nil {
		return applicant, err
	}

	if err = documentService.UpdateDocument(applicant.DocumentId, data.Document); err != nil {
		return applicant, err
	}

	if err = generalInformationService.UpdateGeneralInformation(applicant.GeneralInformationId, data.GeneralInformation); err != nil {
		return applicant, err
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

	if err := connection.DB.Save(&applicant).Error; err != nil {
		return applicant, errors.New("failed to update the applicant data")
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

func ApplicantUploadFile(file *multipart.FileHeader) (result models.Document, err error) {

	var paramPath models.MasterParameter
	connection.DB.Where("param_key = ?", "DOC_PATH_APNT").First(&paramPath)

	if err != nil {
		return result, errors.New("invalid DOC_PATH_APNT value")
	}
	filename, filepath, err := utils.UploadFile(file, paramPath.ParamValue)
	if err != nil {
		return result, err
	}

	result.DocumentFile = filename
	result.DocumentPath = filepath

	return
}

func ApplicantShowFile(id string) (result models.Document, err error) {

	// if err := connection.DB.Model(&models.Document{}).Where("id = ?", id).Pluck("document_path", &result).Error; err != nil {
	if err := connection.DB.Select("documents.*").
		Joins("JOIN applicants ON documents.id = applicants.document_id").
		Where("applicants.id = ?", id).
		Find(&result).Error; err != nil {
		return result, errors.New("failed to get Document Data")
	}

	result.DocumentPath = "." + result.DocumentPath

	return result, nil
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
	// if err := connection.DB.Model(&models.HubunganNasabah{}).Pluck("name", &result).Error; err != nil {
	// 	return nil, errors.New("failed to get Hubungan Nasabah Data")
	// }

	return result, nil
}

func ShowHubunganKeluarga() (result []string, err error) {
	// if err := connection.DB.Model(&models.HubunganKeluarga{}).Pluck("name", &result).Error; err != nil {
	// 	return nil, errors.New("failed to get Hubungan Keluarga Data")
	// }

	return result, nil
}

func ShowLokasiPabrik() (result []string, err error) {
	// if err := connection.DB.Model(&models.LokasiPabrik{}).Pluck("name", &result).Error; err != nil {
	// 	return nil, errors.New("failed to get Lokasi Pabrik Data")
	// }

	return result, nil
}

func ShowMaritalStatus() (result []string, err error) {
	if err := connection.DB.Model(&models.MaritalStatus{}).Pluck("name", &result).Error; err != nil {
		return nil, errors.New("failed to get Marital Status Data")
	}

	return result, nil
}

func ApplicantToJson(applicant models.Applicant, spouse models.SpouseData, idCard models.IdCard, document models.Document, generalInformation any, sectorEconomy any) string {
	data := map[string]interface{}{
		"applicant":          applicant,
		"spouse":             spouse,
		"idCard":             idCard,
		"document":           document,
		"generalInformation": generalInformation,
		"sectorEconomy":      sectorEconomy,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error converting business data to JSON:", err)
		return "{}"
	}
	return string(jsonData)
}
