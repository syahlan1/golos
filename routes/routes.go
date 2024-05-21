package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/syahlan1/golos/controllers"
	"github.com/syahlan1/golos/middleware"
)

func Setup(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	api := app.Group("/api")

	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)
	api.Get("/user", controllers.User)
	api.Post("/logout", controllers.Logout)

	//business
	api.Post("/business/create", middleware.Authorize("create"), controllers.BusinessCreate)
	api.Get("/business/show", controllers.BusinessShow)
	api.Put("/business/update/:id", middleware.Authorize("update"), controllers.BusinessUpdate)
	api.Put("/business/delete/:id", middleware.Authorize("delete"), controllers.BusinessDelete)
	api.Get("/business/show/:id", controllers.BusinessShowDetail)
	api.Get("/business/showcompanyfirstname", controllers.ShowCompanyFirstName)
	api.Get("/business/showcompanytype", controllers.ShowCompanyType)
	api.Get("/business/showbusinessaddresstype", controllers.ShowBusinessAddressType)
	api.Get("/business/showratingclass", controllers.ShowRatingClass)
	api.Get("/business/showkodebursa", controllers.ShowKodeBursa)
	api.Get("/business/showbusinesstype", controllers.ShowBusinessType)

	api.Get("/show-business-applicant", controllers.ShowBusinessApplicant)

	//applicant
	api.Post("/applicant/create", middleware.Authorize("create"), controllers.ApplicantCreate)
	api.Get("/applicant/show", controllers.ApplicantShow)
	api.Get("/applicant/show/:id", controllers.ApplicantShowDetail)
	api.Put("/applicant/update/:id", middleware.Authorize("update"), controllers.ApplicantUpdate)
	api.Put("/applicant/delete/:id", middleware.Authorize("delete"), controllers.ApplicantDelete)
	api.Get("/applicant/show-homestatus", controllers.ShowHomeStatus)
	api.Get("/applicant/show-applicant-addresstype", controllers.ShowApplicantAddressType)
	api.Get("/applicant/show-education", controllers.ShowEducation)
	api.Get("/applicant/show-job-position", controllers.ShowJobPosition)
	api.Get("/applicant/show-business-sector", controllers.ShowBusinessSector)
	api.Get("/applicant/show-kode-instansi", controllers.ShowKodeInstansi)
	api.Get("/applicant/show-negara", controllers.ShowNegara)
	api.Get("/applicant/show-sektor-ekonomi", controllers.ShowSektorEkonomi)
	api.Get("/applicant/show-hubungan-nasabah", controllers.ShowHubunganNasabah)
	api.Get("/applicant/show-hubungan-keluarga", controllers.ShowHubunganKeluarga)
	api.Get("/applicant/show-lokasi-pabrik", controllers.ShowLokasiPabrik)
	api.Get("/applicant/marital-status", controllers.ShowMaritalStatus)

	//zipcode
	api.Get("/provinces", controllers.GetProvinces)
	api.Get("/cities", controllers.GetCitiesByProvince)
	api.Get("/districts", controllers.GetDistrictByCity)
	api.Get("/subdistricts", controllers.GetSubdistrictByDistrict)
	api.Get("/zip-codes", controllers.GetZipCodesBySubdistrict)

	//approve
	api.Post("/approval/create", controllers.CreateApprovalSetting)
	api.Put("/approval/:id", controllers.UpdateApprovalStatus)
	api.Put("/approval/:id/reject", controllers.RejectApproval)
	api.Get("/approval/show", controllers.ShowAllData)
	api.Put("/approval/set-role/:id", controllers.UpdateApprovalWorkflowRoles)
	api.Get("/approval/data/:id", controllers.ApprovalDataDetail)

	//role
	api.Post("/role/create", middleware.Authorize("create_role"), controllers.CreateRole)
	api.Put("/role/update/:id", controllers.UpdateRole)
	api.Get("/role/show", controllers.ShowRole)
	api.Get("/role/:id/permissions", controllers.ShowPermissions)
	api.Get("/role/permissions", controllers.ShowAllPermissions)
	api.Delete("/role/delete/:id", controllers.DeleteRole)

	//Validation
	api.Get("/validation/show", controllers.ShowAllValidations)
	// api.Get("/validation/show/:group_id", controllers.ShowDetailValidation)
	api.Put("/validation/delete/:id", controllers.DeleteValidation)
	api.Put("/validation/update/:id", controllers.UpdateValidation)
	api.Post("api/validation/create", controllers.CreateValidation)

	//Master Code
	api.Get("/master-codes/show", controllers.ShowMasterCode)
	api.Get("/master-codes/show/by-name/:code_group", controllers.ShowDetailMasterCode)
	api.Get("/master-codes/show/by-id/:code_group_id", controllers.ShowDetailMasterCode)
	api.Post("/master-codes/create", controllers.CreateMasterCode)
	api.Put("/master-codes/edit/:id", controllers.UpdateMasterCode)
	api.Put("/master-codes/delete/:id", controllers.DeleteMasterCode)
	api.Get("/master-code-group/show", controllers.ShowMasterCodeGroup)
	api.Post("/master-code-group/create", controllers.CreateMasterCodeGroup)
	api.Put("/master-code-group/edit/:id", controllers.UpdateMasterCodeGroup)
	api.Put("/master-code-group/delete/:id", controllers.DeleteMasterCodeGroup)

	//Master Table
	api.Get("/master-table/show", controllers.ShowMasterTable)
	api.Get("/master-table/show/:id", controllers.ShowMasterTableDetail)
	api.Post("/master-table/create", controllers.CreateMasterTable)
	api.Put("/master-table/delete/:id", controllers.DeleteMasterTable)
	api.Put("/master-table/update/:id", controllers.UpdateMasterTable)

	// //Master Column
	api.Get("/master-column/show", controllers.ShowMasterColumn)
	api.Get("/master-column/show/:id", controllers.ShowMasterColumnDetail)
	api.Get("/master-column/by-table/:id", controllers.ShowMasterColumnByTable)
	api.Post("/master-column/create/:id", controllers.CreateMasterColumn)
	api.Put("/master-column/delete/:id", controllers.DeleteMasterColumn)
	api.Put("/master-column/:id", controllers.UpdateColumnTable)

	// //generate table
	api.Post("/master-table/generate/:id", controllers.GenerateTable)
}
