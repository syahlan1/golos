package models

import "github.com/syahlan1/golos/utils/formatTime"

type SpouseData struct {
	Id            int                 `json:"id"`
	SpouseName    string              `json:"spouse_name"`
	SpouseIdCard  string              `json:"spouse_id_card"`
	SpouseAddress string              `json:"spouse_address"`
	SpouseIdDate  formatTime.WrapDate `json:"spouse_id_date"  gorm:"type:date"`
}
