package connection

import (
	"fmt"
	"os"

	"github.com/syahlan1/golos/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, dbname, port)

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	DB = connection

	connection.AutoMigrate(
		&models.RolePermission{},
		&models.Permission{},
		&models.Roles{},
		&models.Users{},
		&models.RoleModules{},
		&models.RoleTables{},
		&models.RoleWorkflow{},
		&models.Menu{},
		&models.RoleMenu{},

		&models.Business{},
		&models.Applicant{},
		&models.IdCard{},
		&models.SpouseData{},
		&models.Document{},
		&models.GeneralInformation{},
		&models.SectorEconomy{},

		&models.AddressType{},

		&models.HomeStatus{},
		&models.Education{},
		&models.JobPosition{},
		&models.BusinessSector{},
		&models.Negara{},
		&models.MaritalStatus{},
		&models.Nationality{},
		&models.Gender{},

		&models.SectorEconomy1{},
		&models.SectorEconomy2{},
		&models.SectorEconomy3{},
		&models.SectorEconomyOjk{},
		&models.LokasiPabrik{},
		&models.LokasiDati2{},
		&models.HubunganNasabahBank{},
		&models.HubunganKeluarga{},

		&models.Cabang{},
		&models.Program{},
		&models.Segment{},

		&models.CompanyFirstName{},
		&models.CompanyType{},
		&models.ExternalRatingCompany{},
		&models.RatingClass{},
		&models.KodeBursa{},
		&models.BusinessType{},

		&models.ZipCode{},

		&models.Approval{},
		&models.ApprovalSetting{},
		&models.ApprovalWorkflow{},
		&models.ApprovalHistory{},
		&models.ApprovalWorkflowRole{},

		&models.MasterValidation{},

		&models.MasterFile{},
		&models.MasterCode{},
		&models.MasterCodeGroup{},

		&models.MasterModule{},
		&models.MasterTable{},
		&models.MasterTableGroup{},
		&models.MasterTableItem{},
		&models.TableGroupItemStatus{},
		&models.UiType{},
		&models.MasterColumn{},
		&models.MasterSourceColumn{},
		&models.MasterMapperColumn{},
		&models.MasterMapperTable{},
		&models.MasterParameter{},
		&models.MasterWorkflow{},

		&models.OwnershipData{},
		&models.RelationWithBank{},
		&models.DataRekeningDebitur{},
		&models.CustomerLoanInfo{},
	)

	// connection.AutoMigrate(&models.RolePermission{})
	// connection.AutoMigrate(&models.Permission{})
	// connection.AutoMigrate(&models.Roles{})
	// connection.AutoMigrate(&models.Users{})
	// connection.AutoMigrate(&models.RoleModules{})
	// connection.AutoMigrate(&models.RoleTables{})
	// connection.AutoMigrate(&models.RoleWorkflow{})
	// connection.AutoMigrate(&models.Menu{})
	// connection.AutoMigrate(&models.RoleMenu{})

	// connection.AutoMigrate(&models.Business{})
	// connection.AutoMigrate(&models.Applicant{})
	// connection.AutoMigrate(&models.IdCard{})
	// connection.AutoMigrate(&models.SpouseData{})
	// connection.AutoMigrate(&models.Document{})
	// connection.AutoMigrate(&models.GeneralInformation{})
	// connection.AutoMigrate(&models.SectorEconomy{})

	// connection.AutoMigrate(&models.AddressType{})

	// connection.AutoMigrate(&models.HomeStatus{})
	// connection.AutoMigrate(&models.Education{})
	// connection.AutoMigrate(&models.JobPosition{})
	// connection.AutoMigrate(&models.BusinessSector{})
	// connection.AutoMigrate(&models.Negara{})
	// connection.AutoMigrate(&models.MaritalStatus{})
	// connection.AutoMigrate(&models.Nationality{})
	// connection.AutoMigrate(&models.Gender{})

	// connection.AutoMigrate(&models.SectorEconomy1{})
	// connection.AutoMigrate(&models.SectorEconomy2{})
	// connection.AutoMigrate(&models.SectorEconomy3{})
	// connection.AutoMigrate(&models.SectorEconomyOjk{})
	// connection.AutoMigrate(&models.LokasiPabrik{})
	// connection.AutoMigrate(&models.LokasiDati2{})
	// connection.AutoMigrate(&models.HubunganNasabahBank{})
	// connection.AutoMigrate(&models.HubunganKeluarga{})

	// connection.AutoMigrate(&models.Cabang{})
	// connection.AutoMigrate(&models.Program{})
	// connection.AutoMigrate(&models.Segment{})

	// connection.AutoMigrate(&models.CompanyFirstName{})
	// connection.AutoMigrate(&models.CompanyType{})
	// connection.AutoMigrate(&models.ExternalRatingCompany{})
	// connection.AutoMigrate(&models.RatingClass{})
	// connection.AutoMigrate(&models.KodeBursa{})
	// connection.AutoMigrate(&models.BusinessType{})

	// // connection.AutoMigrate(&models.BusinessAddressType{})

	// // connection.AutoMigrate(&models.HomeStatus{})
	// // connection.AutoMigrate(&models.ApplicantAddressType{})
	// // connection.AutoMigrate(&models.Education{})
	// // connection.AutoMigrate(&models.JobPosition{})
	// // connection.AutoMigrate(&models.BusinessSector{})
	// // connection.AutoMigrate(&models.KodeInstansi{})
	// // connection.AutoMigrate(&models.Negara{})
	// // connection.AutoMigrate(&models.SektorEkonomi{})
	// // connection.AutoMigrate(&models.HubunganNasabah{})
	// // connection.AutoMigrate(&models.HubunganKeluarga{})
	// // connection.AutoMigrate(&models.LokasiPabrik{})
	// // connection.AutoMigrate(&models.MaritalStatus{})

	// connection.AutoMigrate(&models.ZipCode{})

	// connection.AutoMigrate(&models.Approval{})
	// connection.AutoMigrate(&models.ApprovalSetting{})
	// connection.AutoMigrate(&models.ApprovalWorkflow{})
	// connection.AutoMigrate(&models.ApprovalHistory{})
	// connection.AutoMigrate(&models.ApprovalWorkflowRole{})

	// connection.AutoMigrate(&models.MasterValidation{})

	// connection.AutoMigrate(&models.MasterCode{})
	// connection.AutoMigrate(&models.MasterCodeGroup{})

	// connection.AutoMigrate(&models.MasterModule{})
	// connection.AutoMigrate(&models.MasterTable{})
	// connection.AutoMigrate(&models.UiType{})
	// connection.AutoMigrate(&models.MasterColumn{})
	// connection.AutoMigrate(&models.MasterSourceColumn{})
	// connection.AutoMigrate(&models.MasterMapperColumn{})
	// connection.AutoMigrate(&models.MasterMapperTable{})
	// connection.AutoMigrate(&models.MasterParameter{})
	// connection.AutoMigrate(&models.MasterWorkflow{})

	// connection.AutoMigrate(&models.OwnershipData{})
	// connection.AutoMigrate(&models.RelationWithBank{})
	// connection.AutoMigrate(&models.DataRekeningDebitur{})
	// connection.AutoMigrate(&models.CustomerLoanInfo{})
}
