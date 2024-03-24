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
	app.Get("/api/applicant/showhomestatus", controllers.ShowHomeStatus)
	app.Get("/api/applicant/showapplicantaddresstype", controllers.ShowApplicantAddressType)
	app.Get("/api/applicant/showeducation", controllers.ShowEducation)
	app.Get("/api/applicant/showjobposition", controllers.ShowJobPosition)
	app.Get("/api/applicant/showbusinesssector", controllers.ShowBusinessSector)
	app.Get("/api/applicant/showkodeinstansi", controllers.ShowKodeInstansi)
	app.Get("/api/applicant/shownegara", controllers.ShowNegara)

	//zipcode
	app.Get("/api/provinces", controllers.GetProvinces)
	app.Get("/api/cities", controllers.GetCitiesByProvince)
	app.Get("/api/districts", controllers.GetDistrictByCity)
	app.Get("/api/subdistricts", controllers.GetSubdistrictByDistrict)
	app.Get("/api/zip-codes", controllers.GetZipCodesBySubdistrict)

	//approve
	app.Put("/api/business/approve/:id", controllers.Authorize("approve"), controllers.BusinessApproveUpdate)
	app.Put("/api/applicant/approve/:id", controllers.Authorize("approve"), controllers.ApplicantApproveUpdate)

	//role
	app.Post("/api/create_role", controllers.Authorize("create_role"), controllers.CreateRole)

}
