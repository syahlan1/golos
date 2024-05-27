package models

type SpouseData struct {
	Id                  int    `json:"id"`
	SpouseName          string `json:"spouse_name"`
	SpouseIdCard        string `json:"spouse_id_card"`
	SpouseAddress       string `json:"spouse_address"`
	GroupNasabah        string `json:"group_nasabah"`
	SektorEkonomi1      string `json:"sektor_ekonomi_1"`
	SektorEkonomi2      string `json:"sektor_ekonomi_2"`
	SektorEkonomi3      string `json:"sektor_ekonomi_3"`
	SektorEkonomiOjk    string `json:"sektor_ekonomi_ojk"`
	NetIncome           int    `json:"net_income"`
	LokasiPabrik        string `json:"lokasi_pabrik"`
	KeyPerson           string `json:"key_person"`
	LokasiDati2         string `json:"lokasi_dati_2"`
	HubunganNasabahBank string `json:"hubungan_nasabah_bank"`
	HubunganKeluarga    string `json:"hubungan_keluarga"`
}
