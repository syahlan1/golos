package businessService

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
	"github.com/syahlan1/golos/services/sectorEconomyService"
	"github.com/syahlan1/golos/utils"
)

func ShowBusinessApplicant() (result models.BusinessApplicant) {
	connection.DB.Where("status != ?", "D").Find(&result.Business)
	connection.DB.Find(&result.Applicant)

	return result
}

func BusinessShow() (result []models.BusinessDetail, err error) {
	var businesses []models.ShowBusiness
	if err = connection.DB.Select("businesses.*, cfn.name AS company_name, ct.name AS company_type, at.name AS address_type",
		"erc.name AS external_rating_company, rc.name AS rating_class, kb.name AS listing_bursa_code, bt.name AS business_type").
		Joins("JOIN company_first_names cfn ON cfn.id = businesses.company_first_name_id").
		Joins("JOIN company_types ct ON ct.id = businesses.company_type_id").
		Joins("JOIN address_types at ON at.id = businesses.address_type_id").
		Joins("JOIN external_rating_companies erc ON erc.id = businesses.external_rating_company_id").
		Joins("JOIN rating_classes rc ON rc.id = businesses.rating_class_id").
		Joins("JOIN kode_bursas kb ON kb.id = businesses.listing_bursa_code_id").
		Joins("JOIN business_types bt ON bt.id = businesses.business_type_id").
		Model(models.Business{}).
		Find(&businesses).Error; err != nil {
		return nil, err
	}

	for _, business := range businesses {
		businessDetail := models.BusinessDetail{
			ShowBusiness: business,
		}

		if business.DocumentId != 0 {
			businessDetail.Document, _ = documentService.ShowDocumentById(business.DocumentId)
		}

		if business.GeneralInformationId != 0 {
			businessDetail.GeneralInformation, _ = generalInformationService.ShowGeneralInformationById(business.GeneralInformationId)
		}

		if business.SectorEconomyId != 0 {
			businessDetail.SectorEconomy, _ = sectorEconomyService.ShowSectorEconomyById(business.SectorEconomyId)
		}

		result = append(result, businessDetail)
	}

	return result, nil
}

func BusinessShowDetail(id any) (result models.BusinessDetail, err error) {
	var business models.ShowBusiness
	if err = connection.DB.Select("businesses.*, cfn.name AS company_name, ct.name AS company_type, at.name AS address_type",
		"erc.name AS external_rating_company, rc.name AS rating_class, kb.name AS listing_bursa_code, bt.name AS business_type").
		Joins("JOIN company_first_names cfn ON cfn.id = businesses.company_first_name_id").
		Joins("JOIN company_types ct ON ct.id = businesses.company_type_id").
		Joins("JOIN address_types at ON at.id = businesses.address_type_id").
		Joins("JOIN external_rating_companies erc ON erc.id = businesses.external_rating_company_id").
		Joins("JOIN rating_classes rc ON rc.id = businesses.rating_class_id").
		Joins("JOIN kode_bursas kb ON kb.id = businesses.listing_bursa_code_id").
		Joins("JOIN business_types bt ON bt.id = businesses.business_type_id").
		Model(models.Business{}).
		First(&business, "businesses.id = ?", id).Error; err != nil {
		return result, err
	}

	businessDetail := models.BusinessDetail{
		ShowBusiness: business,
	}

	if business.DocumentId != 0 {
		businessDetail.Document, _ = documentService.ShowDocumentById(business.DocumentId)
	}

	if business.GeneralInformationId != 0 {
		businessDetail.GeneralInformation, _ = generalInformationService.ShowGeneralInformationById(business.GeneralInformationId)
	}

	if business.SectorEconomyId != 0 {
		businessDetail.SectorEconomy, _ = sectorEconomyService.ShowSectorEconomyById(business.SectorEconomyId)
	}

	return businessDetail, nil
}

func BusinessCreate(username string, data models.CreateBusiness) (err error) {

	business := data.Business
	document := data.Document
	generalInformation := data.GeneralInformation
	sectorEcomony := data.SectorEconomy

	if err = generalInformationService.CreateGeneralInformation(&generalInformation); err != nil {
		return err
	}

	if err = documentService.CreateDocument(&document); err != nil {
		return err
	}

	if err = sectorEconomyService.CreateSectorEconomy(&sectorEcomony); err != nil {
		return err
	}

	// get company first name
	var companyFirstName string
	if err = connection.DB.Select("name").Model(models.CompanyFirstName{}).First(&companyFirstName, "id = ?", business.CompanyFirstNameId).
		Error; err != nil {
		return err
	}

	business.CustomerName = companyFirstName + ". " + business.CompanyName
	business.DocumentId = document.Id
	business.GeneralInformationId = generalInformation.Id
	business.SectorEconomyId = sectorEcomony.Id

	// Buat data bisnis ke database
	if err := connection.DB.Create(&business).Error; err != nil {
		return err
	}

	showGeneralinformation, _ := generalInformationService.ShowGeneralInformationById(generalInformation.Id)
	showSectorEcomony, _ := sectorEconomyService.ShowSectorEconomyById(sectorEcomony.Id)
	showBusiness, _ := BusinessShowDetail(business.Id)

	//generate id
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	approval := models.Approval{
		Id:                id.String(),
		DisplayData:       "Data badan usaha " + business.CompanyName,
		Data:              BusinessToJson(showBusiness, document, showGeneralinformation, showSectorEcomony),
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
		Data:           BusinessToJson(showBusiness, document, showGeneralinformation, showSectorEcomony),
	}
	if err := connection.DB.Create(&history).Error; err != nil {
		return err
	}

	return nil
}

func BusinessUpdate(businessID string, data models.CreateBusiness) (result models.Business, err error) {
	var businesses models.Business
	if err := connection.DB.First(&businesses, businessID).Error; err != nil {
		return result, errors.New("business not found")
	}

	updatedBusiness := data.Business

	if err = documentService.UpdateDocument(businesses.DocumentId, data.Document); err != nil {
		return businesses, err
	}

	if err = generalInformationService.UpdateGeneralInformation(businesses.GeneralInformationId, data.GeneralInformation); err != nil {
		return businesses, err
	}

	if err = sectorEconomyService.UpdateSectorEconomy(businesses.SectorEconomyId, data.SectorEconomy); err != nil {
		return businesses, err
	}

	businesses.Cif = updatedBusiness.Cif
	businesses.CompanyFirstNameId = updatedBusiness.CompanyFirstNameId
	businesses.CompanyName = updatedBusiness.CompanyName
	businesses.CompanyTypeId = updatedBusiness.CompanyTypeId
	businesses.EstablishDate = updatedBusiness.EstablishDate
	businesses.EstablishPlace = updatedBusiness.EstablishPlace
	businesses.CompanyAddress = updatedBusiness.CompanyAddress
	businesses.District = updatedBusiness.District
	businesses.City = updatedBusiness.City
	businesses.ZipCode = updatedBusiness.ZipCode
	businesses.AddressTypeId = updatedBusiness.AddressTypeId
	businesses.ExternalRatingCompanyId = updatedBusiness.ExternalRatingCompanyId
	businesses.RatingClassId = updatedBusiness.RatingClassId
	businesses.RatingDate = updatedBusiness.RatingDate
	businesses.ListingBursaCodeId = updatedBusiness.ListingBursaCodeId
	businesses.ListingBursaDate = updatedBusiness.ListingBursaDate
	businesses.BusinessTypeId = updatedBusiness.BusinessTypeId
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
		return result, errors.New("failed to update the business data")
	}

	return businesses, nil
}

func BusinessDelete(businessID string) (err error) {
	// var businesses models.Business
	// if err := connection.DB.First(&businesses, businessID).Error; err != nil {
	// 	return result, errors.New("business not found")
	// }

	// businesses.Status = "D"

	// if err := connection.DB.Save(&businesses).Error; err != nil {
	// 	return result, errors.New("failed to delete the business data")
	// }

	return connection.DB.Delete(&models.Business{}, businessID).Error
}

func BusinessUploadFile(file *multipart.FileHeader) (result models.Document, err error) {

	var paramPath models.MasterParameter
	err = connection.DB.Where("param_key = ?", "DOC_PATH_BSNS").First(&paramPath).Error

	if err != nil {
		return result, errors.New("invalid DOC_PATH_BSNS value")
	}
	filename, filepath, err := utils.UploadFile(file, paramPath.ParamValue)
	if err != nil {
		return result, err
	}

	result.DocumentFile = filename
	result.DocumentPath = filepath

	return
}

func BusinessShowFile(id string) (result models.Document, err error) {

	// if err := connection.DB.Model(&models.Document{}).Where("id = ?", id).Pluck("document_path", &result).Error; err != nil {
	if err := connection.DB.Select("documents.*").
		Joins("JOIN businesses ON documents.id = businesses.document_id").
		Where("businesses.id = ?", id).
		Find(&result).Error; err != nil {
		return result, errors.New("failed to get Document Data")
	}

	result.DocumentPath = "." + result.DocumentPath

	return result, nil
}

func BusinessApproveUpdate(businessID string) (result models.Business, err error) {
	var business models.Business
	if err := connection.DB.First(&business, businessID).Error; err != nil {
		return result, errors.New("business not found")
	}

	if err := connection.DB.Save(&business).Error; err != nil {
		return result, errors.New("failed to update the business data")
	}

	return business, nil
}

func ShowCompanyFirstName() (result []models.Dropdown) {
	connection.DB.Select("id, name").
		Model(&models.CompanyFirstName{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowCompanyType() (result []models.Dropdown) {
	connection.DB.Select("id, CONCAT(code, ' - ', name) AS name").
		Model(&models.CompanyType{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowBusinessAddressType() (result []string, err error) {
	var businessAddressType []string

	if err := connection.DB.Model(&models.BusinessAddressType{}).Pluck("name", &businessAddressType).Error; err != nil {
		return result, err
	}

	return businessAddressType, nil
}

func ShowExternalRatingCompany() (result []models.Dropdown) {
	connection.DB.Select("id, name").
		Model(&models.ExternalRatingCompany{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowRatingClass(ExternalRatingId string) (result []models.Dropdown) {
	connection.DB.Select("id, name").
		Model(&models.RatingClass{}).
		Where("external_rating_id = ?", ExternalRatingId).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowKodeBursa() (result []models.Dropdown) {
	connection.DB.Select("id, name").
		Model(&models.KodeBursa{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowBusinessType() (result []models.Dropdown) {
	connection.DB.Select("id, CONCAT(code, ' - ', name) AS name").
		Model(&models.BusinessType{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func GetProvinces() (result []string, err error) {
	var provinces []string
	if err := connection.DB.Model(&models.ZipCode{}).Distinct("province").Pluck("province", &provinces).Error; err != nil {
		return result, err
	}

	return provinces, nil
}

func GetCitiesByProvince(province string) (result []string, err error) {
	var cities []string
	if err := connection.DB.Model(&models.ZipCode{}).Distinct("city").Where("province = ?", province).Pluck("city", &cities).Error; err != nil {
		return result, err
	}

	return cities, nil
}

func GetDistrictByCity(city string) (result []string, err error) {
	var district []string
	if err := connection.DB.Model(&models.ZipCode{}).Distinct("district").Where("city = ?", city).Pluck("district", &district).Error; err != nil {
		return result, err
	}

	return district, nil
}

func GetSubdistrictByDistrict(district string) (result []string, err error) {
	var subdistrict []string
	if err := connection.DB.Model(&models.ZipCode{}).Distinct("subdistrict").Where("district = ?", district).Pluck("subdistrict", &subdistrict).Error; err != nil {
		return result, err
	}

	return subdistrict, nil
}

func GetZipCodesBySubdistrict(subdistrict string) (result []string, err error) {
	var zipCodes []string
	if err := connection.DB.Model(&models.ZipCode{}).Where("subdistrict = ?", subdistrict).Pluck("zip_code", &zipCodes).First(&zipCodes).Error; err != nil {
		return result, err
	}

	return zipCodes, nil
}

func BusinessToJson(business any, document models.Document, generalInformation any, sectorEconomy any) string {
	data := map[string]interface{}{
		"business":           business,
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
