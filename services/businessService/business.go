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
	"github.com/syahlan1/golos/utils"
)

func ShowBusinessApplicant() (result models.BusinessApplicant) {
	connection.DB.Where("status != ?", "D").Find(&result.Business)
	connection.DB.Find(&result.Applicant)

	return result
}

func BusinessShow() (result []models.Business) {
	connection.DB.Where("status != ?", "D").Find(&result)

	return result
}

func BusinessShowDetail(id string) (result models.Business) {
	connection.DB.Where("id = ? AND status != ?", id, "D").Find(&result)

	return result
}

func BusinessCreate(username string, data models.CreateBusiness) (err error) {

	business := data.Business
	business.Status = "L"

	document := data.Document

	generalInformation := data.GeneralInformation
	generalInformation.BankName = business.BankName
	generalInformation.KCP = business.KCP
	generalInformation.SubProgram = business.SubProgram
	generalInformation.Analisis = business.Analisis
	generalInformation.CabangPencairan = business.CabangPencairan
	generalInformation.CabangAdmin = business.CabangAdmin
	generalInformation.TglAplikasi = business.TglAplikasi
	generalInformation.TglPenerusan = business.TglPenerusan
	generalInformation.Segmen = business.Segmen
	generalInformation.NoAplikasi = business.NoAplikasi
	generalInformation.MarketInterestRate = business.MarketInterestRate
	generalInformation.RequestedInterestRate = business.RequestedInterestRate
	// generalInformation.DocumentFile = business.DocumentFile
	generalInformation.Status = business.Status

	if err := connection.DB.Create(&generalInformation).Error; err != nil {
		return err
	}
	if err := connection.DB.Create(&document).Error; err != nil {
		return err
	}

	// business := data.Business
	business.CustomerName = business.CompanyFirstName + ". " + business.CompanyName
	// business.Status = "L"
	business.DocumentId = document.Id
	business.GeneralInformationId = generalInformation.Id

	// Buat data bisnis ke database
	if err := connection.DB.Create(&business).Error; err != nil {
		return err
	}

	//generate id
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	approval := models.Approval{
		Id:                id.String(),
		DisplayData:       "Data badan usaha " + business.CompanyName,
		Data:              BusinessToJson(business, document, generalInformation),
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
		Data:           BusinessToJson(business, document, generalInformation),
	}
	if err := connection.DB.Create(&history).Error; err != nil {
		return err
	}

	return nil
}

func BusinessUpdate(businessID string, updatedBusiness models.Business) (result models.Business, err error) {
	var businesses models.Business
	if err := connection.DB.First(&businesses, businessID).Error; err != nil {
		return result, errors.New("business not found")
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
		return result, errors.New("failed to update the business data")
	}

	return businesses, nil
}

func BusinessDelete(businessID string) (result models.Business, err error) {
	var businesses models.Business
	if err := connection.DB.First(&businesses, businessID).Error; err != nil {
		return result, errors.New("business not found")
	}

	businesses.Status = "D"

	if err := connection.DB.Save(&businesses).Error; err != nil {
		return result, errors.New("failed to delete the business data")
	}

	return businesses, nil
}

func BusinessUploadFile(file *multipart.FileHeader) (result models.Document, err error) {

	var paramPath models.MasterParameter
	connection.DB.Where("param_key = ?", "DOC_PATH_BSNS").First(&paramPath)

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

func ShowCompanyFirstName() (result []string, err error) {
	var companyFirstNames []string

	if err := connection.DB.Model(&models.CompanyFirstName{}).Pluck("name", &companyFirstNames).Error; err != nil {
		return result, err
	}

	return companyFirstNames, nil
}

func ShowCompanyType() (result []string, err error) {
	var companyType []string

	if err := connection.DB.Model(&models.CompanyType{}).Pluck("name", &companyType).Error; err != nil {
		return result, err
	}

	return companyType, nil
}

func ShowBusinessAddressType() (result []string, err error) {
	var businessAddressType []string

	if err := connection.DB.Model(&models.BusinessAddressType{}).Pluck("name", &businessAddressType).Error; err != nil {
		return result, err
	}

	return businessAddressType, nil
}

func ShowEternalRatingCompany() (result []string, err error) {
	var eternalRatingCompany []string

	if err := connection.DB.Model(&models.EternalRatingCompany{}).Pluck("name", &eternalRatingCompany).Error; err != nil {
		return result, err
	}

	return eternalRatingCompany, nil
}

func ShowRatingClass() (result []string, err error) {
	var ratingClass []string

	if err := connection.DB.Model(&models.RatingClass{}).Pluck("name", &ratingClass).Error; err != nil {
		return result, err
	}

	return ratingClass, nil
}

func ShowKodeBursa() (result []string, err error) {
	var kodeBursa []string

	if err := connection.DB.Model(&models.KodeBursa{}).Pluck("name", &kodeBursa).Error; err != nil {
		return result, err
	}

	return kodeBursa, nil
}

func ShowBusinessType() (result []string, err error) {
	var businessType []string

	if err := connection.DB.Model(&models.BusinessType{}).Pluck("name", &businessType).Error; err != nil {
		return result, err
	}

	return businessType, nil
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
