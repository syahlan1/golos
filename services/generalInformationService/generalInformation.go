package generalInformationService

import (
	"errors"
	"strconv"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
	"gorm.io/gorm"
)

func ShowCabangPencairan() (result []models.Dropdown) {
	connection.DB.Select("id, CONCAT(branch_code, ' - ', name) AS name").
		Model(&models.Cabang{}).
		Where("cabang_pencairan = ?", true).Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowCabangAdmin() (result []models.Dropdown) {
	connection.DB.Select("id, CONCAT(branch_code, ' - ', name) AS name").
		Model(&models.Cabang{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowSegment() (result []models.Dropdown) {
	connection.DB.Select("id, name").
		Model(&models.Segment{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowProgram() (result []models.Dropdown) {
	connection.DB.Select("id, name").
		Model(&models.Program{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func GenerateApplicationNumber(CabangAdmin string) (result string, err error) {

	cabangAdminInt, err := strconv.Atoi(CabangAdmin)
	if err != nil {
		return
	}

	result, err = utils.GenerateApplicationNumber(cabangAdminInt)
	if err != nil {
		return
	}

	return result, nil
}

func CreateGeneralInformation(data *models.GeneralInformation) (err error) {

	data.NoReferensi, err = utils.GenerateApplicationNumber(data.CabangAdminId)
	if err != nil {
		return
	}
	data.NoReferensi += "C"

	return connection.DB.Create(&data).Error
}

func ShowGeneralInformationById(id int) (result models.ShowGeneralInformation, err error) {

	if err = connection.DB.Select("general_informations.*, p.name AS sub_program, c1.name AS cabang_pencairan, c2.name AS cabang_admin, s.name AS segmen").
		Model(&models.GeneralInformation{}).
		Joins("JOIN programs p ON p.id = general_informations.sub_program_id").
		Joins("JOIN cabangs c1 ON c1.id = general_informations.cabang_pencairan_id").
		Joins("JOIN cabangs c2 ON c2.id = general_informations.cabang_admin_id").
		Joins("JOIN segments s ON s.id = general_informations.segmen_id").
		First(&result, "general_informations.id = ?", id).Error; !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		return
	}

	return result, nil
}

func UpdateGeneralInformation(id any, data models.GeneralInformation) (err error) {

	var generalInformation models.GeneralInformation
	if err := connection.DB.First(&generalInformation, id).Error; err != nil {
		return errors.New("general Information not found : " + err.Error())
	}

	generalInformation.BankName = data.BankName
	generalInformation.KCP = data.KCP
	generalInformation.SubProgramId = data.SubProgramId
	generalInformation.Analisis = data.Analisis
	generalInformation.CabangPencairanId = data.CabangPencairanId
	generalInformation.CabangAdminId = data.CabangAdminId
	generalInformation.TglAplikasi = data.TglAplikasi
	generalInformation.TglPenerusan = data.TglPenerusan
	generalInformation.SegmenId = data.SegmenId
	generalInformation.NoAplikasi = data.NoAplikasi
	generalInformation.MarketInterestRate = data.MarketInterestRate
	generalInformation.RequestedInterestRate = data.RequestedInterestRate

	if err := connection.DB.Save(&generalInformation).Error; err != nil {
		return errors.New("failed to update the general information data : " + err.Error())
	}

	return
}

func DeleteGeneralInformation(id any) error {
	return connection.DB.Delete(&models.GeneralInformation{}, id).Error
}
