package utils

import (
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func GenerateApplicationNumber(CabangAdmin int) (result string, err error) {

	date := GetDateNowFormat()
	var seqNumber, initCode string

	if err = connection.DB.Select("branch_code").
		Model(models.Cabang{}).
		First(&initCode, CabangAdmin).Error; err != nil {
		return
	}

	if err = connection.DB.Select("LPAD((COUNT(*)+1)::TEXT,6,'0')").
		Table("general_informations AS gi").
		Joins("JOIN cabangs c ON c.id = gi.cabang_admin_id").
		Where("gi.created_at::date = ?", GetDateNow()).
		Scan(&seqNumber).Error; err != nil {
		return
	}

	result = date + initCode + seqNumber

	return
}
