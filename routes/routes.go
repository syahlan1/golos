package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/syahlan1/golos/controllers"
)

func Setup(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	//business
	app.Post("/api/business/create", controllers.Authorize("create"), controllers.BusinessCreate)
	app.Get("/api/business/show", controllers.BusinessShow)
	app.Put("/api/business/update/:id", controllers.Authorize("update"), controllers.BusinessUpdate)
	app.Delete("/api/business/delete/:id", controllers.Authorize("delete"), controllers.BusinessDelete)
	app.Get("/api/business/showcompanyfirstname", controllers.ShowCompanyFirstName)
	app.Get("/api/business/showcompanytype", controllers.ShowCompanyType)
	app.Get("/api/business/showbusinessaddresstype", controllers.ShowBusinessAddressType)
	app.Get("/api/business/showratingclass", controllers.ShowRatingClass)
	app.Get("/api/business/showkodebursa", controllers.ShowKodeBursa)
	app.Get("/api/business/showbusinesstype", controllers.ShowBusinessType)

	//applicant
	app.Post("/api/applicant/create", controllers.Authorize("create"), controllers.ApplicantCreate)
	app.Get("/api/applicant/show", controllers.ApplicantShow)
	app.Put("/api/applicant/update/:id", controllers.Authorize("update"), controllers.ApplicantUpdate)
	app.Delete("/api/applicant/delete/:id", controllers.Authorize("delete"), controllers.ApplicantDelete)
	app.Get("/api/applicant/show-homestatus", controllers.ShowHomeStatus)
	app.Get("/api/applicant/show-applicant-addresstype", controllers.ShowApplicantAddressType)
	app.Get("/api/applicant/show-education", controllers.ShowEducation)
	app.Get("/api/applicant/show-job-position", controllers.ShowJobPosition)
	app.Get("/api/applicant/show-business-sector", controllers.ShowBusinessSector)
	app.Get("/api/applicant/show-kode-instansi", controllers.ShowKodeInstansi)
	app.Get("/api/applicant/show-negara", controllers.ShowNegara)
	app.Get("/api/applicant/show-sektor-ekonomi", controllers.ShowSektorEkonomi)
	app.Get("/api/applicant/show-hubungan-nasabah", controllers.ShowHubunganNasabah)
	app.Get("/api/applicant/show-hubungan-keluarga", controllers.ShowHubunganKeluarga)
	app.Get("/api/applicant/show-lokasi-pabrik", controllers.ShowLokasiPabrik)
	app.Get("/api/applicant/marital-status", controllers.ShowMaritalStatus)

	//zipcode
	app.Get("/api/provinces", controllers.GetProvinces)
	app.Get("/api/cities", controllers.GetCitiesByProvince)
	app.Get("/api/districts", controllers.GetDistrictByCity)
	app.Get("/api/subdistricts", controllers.GetSubdistrictByDistrict)
	app.Get("/api/zip-codes", controllers.GetZipCodesBySubdistrict)

	//approve
	app.Post("/api/approval/create", controllers.CreateApprovalSetting)
	app.Post("/api/approval/:id", controllers.UpdateApprovalStatus)
	app.Put("/api/approval/:id/reject", controllers.RejectApproval)
	app.Get("/api/approval/show", controllers.ShowAllData)

	//role
	app.Post("/api/role/create", controllers.Authorize("create_role"), controllers.CreateRole)
	app.Put("/api/role/update/:id", controllers.UpdateRole)

}
