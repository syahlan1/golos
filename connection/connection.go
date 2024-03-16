package connection

import (
	"github.com/syahlan1/golos/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "sqlserver://sa:12345@127.0.0.1:1434?database=LOS&connection+timeout=30"

	connection, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.Users{})
	// connection.AutoMigrate(&models.Roles{})
	// connection.AutoMigrate(&models.UserRole{})
	connection.AutoMigrate(&models.Business{})
	connection.AutoMigrate(&models.Applicant{})

	connection.AutoMigrate(&models.CompanyFirstName{})
	connection.AutoMigrate(&models.CompanyType{})
	connection.AutoMigrate(&models.BusinessAddressType{})
	connection.AutoMigrate(&models.EternalRatingCompany{})
	connection.AutoMigrate(&models.RatingClass{})
	connection.AutoMigrate(&models.KodeBursa{})
	connection.AutoMigrate(&models.BusinessType{})

	connection.AutoMigrate(&models.HomeStatus{})
	connection.AutoMigrate(&models.ApplicantAddressType{})
	connection.AutoMigrate(&models.Education{})
	connection.AutoMigrate(&models.JobPosition{})
	connection.AutoMigrate(&models.BusinessSector{})
	connection.AutoMigrate(&models.KodeInstansi{})
	connection.AutoMigrate(&models.Negara{})

	connection.AutoMigrate(&models.ZipCode{})

}
