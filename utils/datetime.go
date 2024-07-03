package utils

import (
	"time"

	"github.com/syahlan1/golos/models"
	"gorm.io/gorm"
)

func GetDateNow() string {
	return time.Now().Format("2006-01-02")
}

func GetDateNowFormat() string {
	return time.Now().Format("02012006")
}

func GetDateTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func SoftDelete(username string) models.ModelMasterForm {
	return models.ModelMasterForm{
		DeletedAt: gorm.DeletedAt{
			Time: time.Now(),
			Valid: true,},
		UpdatedBy: username,
	}
}
