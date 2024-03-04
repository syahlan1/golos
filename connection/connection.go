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
	connection.AutoMigrate(&models.Business{})
	connection.AutoMigrate(&models.Applicant{})

}
