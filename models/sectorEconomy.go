package models

type SectorEconomy struct {
	Id                    int    `json:"id" gorm:"primaryKey"`
	GroupNasabah          string `json:"group_nasabah"`
	SektorEkonomi1Id      int    `json:"sektor_ekonomi_1_id"`
	SektorEkonomi2Id      int    `json:"sektor_ekonomi_2_id"`
	SektorEkonomi3Id      int    `json:"sektor_ekonomi_3_id"`
	SektorEkonomiOjkId    int    `json:"sektor_ekonomi_ojk_id"`
	NetIncome             int    `json:"net_income"`
	LokasiPabrikId        int    `json:"lokasi_pabrik_id"`
	KeyPerson             string `json:"key_person"`
	LokasiDati2Id         int    `json:"lokasi_dati_2_id"`
	HubunganNasabahBankId int    `json:"hubungan_nasabah_bank_id"`
	HubunganKeluargaId    int    `json:"hubungan_keluarga_id"`
	BICheck               bool   `json:"bi_check"`
	PelakasanaId          int    `json:"pelakasana_id"`
	TglTerakhirCek        string `json:"tgl_terakhir_cek"`
	Model
}

type ShowSectorEconomy struct {
	Id                    int    `json:"id"`
	GroupNasabah          string `json:"group_nasabah"`
	SektorEkonomi1Id      int    `json:"sektor_ekonomi_1_id"`
	SektorEkonomi1        string `json:"sektor_ekonomi_1"`
	SektorEkonomi2Id      int    `json:"sektor_ekonomi_2_id"`
	SektorEkonomi2        string `json:"sektor_ekonomi_2"`
	SektorEkonomi3Id      int    `json:"sektor_ekonomi_3_id"`
	SektorEkonomi3        string `json:"sektor_ekonomi_3"`
	SektorEkonomiOjkId    int    `json:"sektor_ekonomi_ojk_id"`
	SektorEkonomiOjk      string `json:"sektor_ekonomi_ojk"`
	NetIncome             int    `json:"net_income"`
	LokasiPabrikId        int    `json:"lokasi_pabrik_id"`
	LokasiPabrik          string `json:"lokasi_pabrik"`
	KeyPerson             string `json:"key_person"`
	LokasiDati2Id         int    `json:"lokasi_dati_2_id"`
	LokasiDati2           string `json:"lokasi_dati_2"`
	HubunganNasabahBankId int    `json:"hubungan_nasabah_bank_id"`
	HubunganNasabahBank   string `json:"hubungan_nasabah_bank"`
	HubunganKeluargaId    int    `json:"hubungan_keluarga_id"`
	HubunganKeluarga      string `json:"hubungan_keluarga"`
	BICheck               bool   `json:"bi_check"`
	PelakasanaId          int    `json:"pelakasana_id"`
	Pelakasana            int    `json:"pelakasana"`
	TglTerakhirCek        string `json:"tgl_terakhir_cek"`
}

type SectorEconomy1 struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type SectorEconomy2 struct {
	Id               int    `json:"id" gorm:"primaryKey"`
	SectorEconomy1Id int    `json:"sector_economy_1_id"`
	Code             string `json:"code" gorm:"type:varchar(10)"`
	Name             string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type SectorEconomy3 struct {
	Id               int    `json:"id" gorm:"primaryKey"`
	SectorEconomy2Id int    `json:"sector_economy_2_id"`
	Code             string `json:"code" gorm:"type:varchar(10)"`
	Name             string `json:"name" gorm:"type:varchar(250)"`
	Seq              string `json:"seq" gorm:"type:varchar(10)"`
	Model
}

type SectorEconomyOjk struct {
	Id               int    `json:"id" gorm:"primaryKey"`
	SectorEconomy3Id int    `json:"sector_economy_2_id"`
	Code             string `json:"code" gorm:"type:varchar(10)"`
	Name             string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type LokasiPabrik struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type LokasiDati2 struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type HubunganNasabahBank struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}

type HubunganKeluarga struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"type:varchar(10)"`
	Name string `json:"name" gorm:"type:varchar(250)"`
	Model
}
