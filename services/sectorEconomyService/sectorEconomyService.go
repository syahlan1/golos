package sectorEconomyService

import (
	"errors"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
)

func ShowSektorEkonomi1() (result []models.Dropdown) {
	connection.DB.Select("id, CONCAT(code, ' - ', name) AS name").
		Model(&models.SectorEconomy1{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowSektorEkonomi2(idSectorEconomy1 string) (result []models.Dropdown) {
	connection.DB.Select("id, CONCAT(code, ' - ', name) AS name").
		Model(&models.SectorEconomy2{}).
		Where("sector_economy1_id = ?", idSectorEconomy1).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowSektorEkonomi3(idSectorEconomy2 string) (result []models.Dropdown) {
	connection.DB.Select("id, CONCAT(code, ' - ', name) AS name").
		Model(&models.SectorEconomy3{}).
		Where("sector_economy2_id = ?", idSectorEconomy2).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowSektorEkonomiOjk(idSectorEconomy3 string) (result []models.Dropdown) {

	connection.DB.Select("id, CONCAT(code, ' - ', name) AS name").
		Model(&models.SectorEconomyOjk{}).
		Where("sector_economy3_id = ?", idSectorEconomy3).
		Find(&result)

	if result == nil {
		result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	}
	return
}

func ShowLokasiPabrik() (result []models.Dropdown) {
	connection.DB.Select("id, CONCAT(code, ' - ', name) AS name").
		Model(&models.LokasiPabrik{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowLokasiDati2() (result []models.Dropdown) {
	connection.DB.Select("id, CONCAT(code, ' - ', name) AS name").
		Model(&models.LokasiDati2{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowHubunganNasabahBank() (result []models.Dropdown) {
	connection.DB.Select("id, name").
		Model(&models.HubunganNasabahBank{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func ShowHubunganKeluarga() (result []models.Dropdown) {
	connection.DB.Select("id, name").
		Model(&models.HubunganKeluarga{}).
		Find(&result)

	result = utils.Prepend(result, models.Dropdown{Name: "- SELECT -"})
	return
}

func CreateSectorEconomy(data *models.SectorEconomy) (err error) {

	if !data.BICheck {
		data.PelakasanaId = 0
	}

	return connection.DB.Create(&data).Error
}

func ShowSectorEconomyById(id int) (result models.ShowSectorEconomy, err error) {

	if err = connection.DB.Select("sector_economies.*, se1.Name AS sektor_ekonomi1, se2.Name AS sektor_ekonomi2",
		"se3.Name AS sektor_ekonomi3, seo.Name AS sektor_ekonomi_ojk, lp.Name AS lokasi_pabrik, ld.Name AS lokasi_dati2",
		"hnb.Name AS hubungan_nasabah_bank, hk.Name AS hubungan_keluarga").
		Joins("JOIN sector_economy1 se1 ON se1.id = sector_economies.sektor_ekonomi1_id").
		Joins("JOIN sector_economy2 se2 ON se2.id = sector_economies.sektor_ekonomi2_id").
		Joins("JOIN sector_economy3 se3 ON se3.id = sector_economies.sektor_ekonomi3_id").
		Joins("JOIN sector_economy_ojks seo ON seo.id = sector_economies.sektor_ekonomi_ojk_id").
		Joins("JOIN lokasi_pabriks lp ON lp.id = sector_economies.lokasi_pabrik_id").
		Joins("JOIN lokasi_dati2 ld ON ld.id = sector_economies.lokasi_dati2_id").
		Joins("JOIN hubungan_nasabah_banks hnb ON hnb.id = sector_economies.hubungan_nasabah_bank_id").
		Joins("JOIN hubungan_keluargas hk ON hk.id = sector_economies.hubungan_keluarga_id").
		Model(&models.SectorEconomy{}).
		First(&result, "sector_economies.id = ?", id).Error; err != nil {
		return
	}
	return
}

func UpdateSectorEconomy(id int, data models.SectorEconomy) (err error) {

	var sectorEconomy models.SectorEconomy
	if err := connection.DB.First(&sectorEconomy, id).Error; err != nil {
		return errors.New("sectorEconomy not found")
	}
	updatedSectorEconomy := data

	sectorEconomy.GroupNasabah = updatedSectorEconomy.GroupNasabah
	sectorEconomy.SektorEkonomi1Id = updatedSectorEconomy.SektorEkonomi1Id
	sectorEconomy.SektorEkonomi2Id = updatedSectorEconomy.SektorEkonomi2Id
	sectorEconomy.SektorEkonomi3Id = updatedSectorEconomy.SektorEkonomi3Id
	sectorEconomy.SektorEkonomiOjkId = updatedSectorEconomy.SektorEkonomiOjkId
	sectorEconomy.NetIncome = updatedSectorEconomy.NetIncome
	sectorEconomy.LokasiPabrikId = updatedSectorEconomy.LokasiPabrikId
	sectorEconomy.KeyPerson = updatedSectorEconomy.KeyPerson
	sectorEconomy.LokasiDati2Id = updatedSectorEconomy.LokasiDati2Id
	sectorEconomy.HubunganNasabahBankId = updatedSectorEconomy.HubunganNasabahBankId
	sectorEconomy.HubunganKeluargaId = updatedSectorEconomy.HubunganKeluargaId

	if err := connection.DB.Save(&sectorEconomy).Error; err != nil {
		return errors.New("failed to update the spouse data")
	}

	return
}
