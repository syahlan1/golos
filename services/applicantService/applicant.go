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
	var applicants []models.ShowApplicant

	// Find all applicants with status "L"
	if err = connection.DB.
		Select("applicants.*, hs.name AS home_status, ms.name AS marital_status, n.name AS nationality",
			"e.name AS education, j.name AS job_position, bs.name AS business_sector, nd.name AS negara_domisili, g.name AS gender").
		Joins("JOIN home_statuses hs ON hs.id = applicants.home_status_id").
		Joins("JOIN marital_statuses ms ON ms.id = applicants.marital_status_id").
		Joins("JOIN nationalities n ON n.id = applicants.nationality_id").
		Joins("JOIN educations e ON e.id = applicants.education_id").
		Joins("JOIN job_positions j ON j.id = applicants.job_position_id").
		Joins("JOIN business_sectors bs ON bs.id = applicants.business_sector_id").
		Joins("JOIN negaras nd ON nd.id = applicants.negara_domisili_id").
		Joins("JOIN genders g ON g.id = applicants.gender_id").
		Model(models.Applicant{}).
		Find(&applicants).Error; err != nil {
		return nil, err
	}

	// Load related data for each applicant
	for _, applicant := range applicants {
		applicantDetail := models.ApplicantDetail{
			ShowApplicant: applicant,
		}

		// Load related IdCard
		if applicant.IdCard != 0 {
			applicantDetail.IdCard, _ = idCardService.ShowIdCardById(applicant.IdCard)
		}

		// Load related Document
		if applicant.DocumentId != 0 {
			applicantDetail.Document, _ = documentService.ShowDocumentById(applicant.DocumentId)
		}

		// Load related Spouse
		if applicant.SpouseId != 0 {
			applicantDetail.SpouseData, _ = spouseService.ShowSpouseById(applicant.SpouseId)
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

func ApplicantShowDetail(applicantID any) (result models.ApplicantDetail, err error) {

	var applicant models.ShowApplicant
	// Find the applicant
	if err = connection.DB.
		Select("applicants.*, hs.name AS home_status, ms.name AS marital_status, n.name AS nationality",
			"e.name AS education, j.name AS job_position, bs.name AS business_sector, nd.name AS negara_domisili, g.name AS gender").
		Joins("JOIN home_statuses hs ON hs.id = applicants.home_status_id").
		Joins("JOIN marital_statuses ms ON ms.id = applicants.marital_status_id").
		Joins("JOIN nationalities n ON n.id = applicants.nationality_id").
		Joins("JOIN educations e ON e.id = applicants.education_id").
		Joins("JOIN job_positions j ON j.id = applicants.job_position_id").
		Joins("JOIN business_sectors bs ON bs.id = applicants.business_sector_id").
		Joins("JOIN negaras nd ON nd.id = applicants.negara_domisili_id").
		Joins("JOIN genders g ON g.id = applicants.gender_id").
		Model(models.Applicant{}).
		First(&applicant, "applicant.id = ?", applicantID).Error; err != nil {
		return result, errors.New("applicant Not Found")
	}

	// Initialize the detail struct with applicant data
	applicantDetail := models.ApplicantDetail{
		ShowApplicant: applicant,
	}

	// Load related IdCard
	if applicant.IdCard != 0 {
		applicantDetail.IdCard, _ = idCardService.ShowIdCardById(applicant.IdCard)
	}

	// Load related Document
	if applicant.DocumentId != 0 {
		applicantDetail.Document, _ = documentService.ShowDocumentById(applicant.DocumentId)
	}

	// Load related Spouse
	if applicant.SpouseId != 0 {
		applicantDetail.SpouseData, _ = spouseService.ShowSpouseById(applicant.SpouseId)
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
	// applicant.Status = "L"

	// Buat data appplicant ke database
	if err := connection.DB.Create(&applicant).Error; err != nil {
		return err
	}

	showGeneralInformation, _ := generalInformationService.ShowGeneralInformationById(generalInformation.Id)
	showSectorEconomy, _ := sectorEconomyService.ShowSectorEconomyById(sectorEconomy.Id)
	showIdCard, _ := idCardService.ShowIdCardById(idCard.Id)
	showApplicant, _ := ApplicantShowDetail(applicant.Id)

	//generate id
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	approval := models.Approval{
		Id:                id.String(),
		DisplayData:       "Data " + applicant.CustomerName,
		Data:              ApplicantToJson(showApplicant, spouse, showIdCard, document, showGeneralInformation, showSectorEconomy),
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
		Data:           ApplicantToJson(showApplicant, spouse, showIdCard, document, showGeneralInformation, showSectorEconomy),
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

	if err = sectorEconomyService.UpdateSectorEconomy(applicant.SectorEconomyId, data.SectorEconomy); err != nil {
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
	applicant.HomeStatusId = updatedApplicant.HomeStatusId
	applicant.StaySince = updatedApplicant.StaySince
	applicant.NoTelp = updatedApplicant.NoTelp
	applicant.NoFax = updatedApplicant.NoFax
	applicant.BirthPlace = updatedApplicant.BirthPlace
	applicant.BirthDate = updatedApplicant.BirthDate
	applicant.MaritalStatusId = updatedApplicant.MaritalStatusId
	applicant.GenderId = updatedApplicant.GenderId
	applicant.NationalityId = updatedApplicant.NationalityId
	applicant.NumberOfChildren = updatedApplicant.NumberOfChildren
	applicant.NoKartuKeluarga = updatedApplicant.NoKartuKeluarga
	applicant.EducationId = updatedApplicant.EducationId
	applicant.JobPositionId = updatedApplicant.JobPositionId
	applicant.BusinessSectorId = updatedApplicant.BusinessSectorId
	applicant.EstablishDate = updatedApplicant.EstablishDate
	applicant.NPWP = updatedApplicant.NPWP
	applicant.GrossIncomePerMonth = updatedApplicant.GrossIncomePerMonth
	applicant.NumberOfEmployees = updatedApplicant.NumberOfEmployees
	applicant.MotherName = updatedApplicant.MotherName
	applicant.NamaPelaporan = updatedApplicant.NamaPelaporan
	applicant.NegaraDomisiliId = updatedApplicant.NegaraDomisiliId
	applicant.NamaInstansi = updatedApplicant.NamaInstansi
	applicant.KodeInstansi = updatedApplicant.KodeInstansi
	applicant.NoPegawai = updatedApplicant.NoPegawai

	if err := connection.DB.Save(&applicant).Error; err != nil {
		return applicant, errors.New("failed to update the applicant data")
	}

	return applicant, nil
}

func ApplicantDelete(applicantID string) (err error) {

	// var applicant models.Applicant
	// if err := connection.DB.First(&applicant, applicantID).Error; err != nil {
	// 	return errors.New("applicant not found")
	// }

	// applicant.Status = "D"

	// if err := connection.DB.Save(&applicant).Error; err != nil {
	// 	return errors.New("failed to delete the applicant data")
	// }

	return connection.DB.Delete(&models.Applicant{}, applicantID).Error
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

func ShowHomeStatus() (result []models.Dropdown) {
	connection.DB.Select("id, CONCAT(code, ' - ', name) AS name").
		Model(&models.HomeStatus{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowApplicantAddressType() (result []string, err error) {
	if err := connection.DB.Model(&models.ApplicantAddressType{}).Pluck("name", &result).Error; err != nil {
		return nil, errors.New("failed to get Applicant Address Type")
	}

	return result, nil
}

func ShowEducation() (result []models.Dropdown) {
	connection.DB.Select("id, name").
		Model(&models.Education{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowJobPosition() (result []models.Dropdown) {
	connection.DB.Select("id, name").
		Model(&models.JobPosition{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowBusinessSector() (result []models.Dropdown) {
	connection.DB.Select("id, name").
		Model(&models.BusinessSector{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowKodeInstansi() (result []string, err error) {
	if err := connection.DB.Model(&models.KodeInstansi{}).Pluck("name", &result).Error; err != nil {
		return nil, errors.New("failed to get Kode Instansi Data")
	}

	return result, nil
}

func ShowNegara() (result []models.Dropdown) {
	connection.DB.Select("id, name").
		Model(&models.Negara{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowSektorEkonomi() (result []string, err error) {
	if err := connection.DB.Model(&models.SektorEkonomi{}).Pluck("name", &result).Error; err != nil {
		return nil, errors.New("failed to get Sektor Ekonomi Data")
	}

	return result, nil
}

func ShowMaritalStatus() (result []models.Dropdown) {
	connection.DB.Select("id, name").
		Model(&models.MaritalStatus{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowNationality() (result []models.Dropdown) {
	connection.DB.Select("id, name").
		Model(&models.Nationality{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowGender() (result []models.Dropdown) {
	connection.DB.Select("id, name").
		Model(&models.Gender{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ApplicantToJson(applicant any, spouse models.SpouseData, idCard any, document models.Document, generalInformation any, sectorEconomy any) string {
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
