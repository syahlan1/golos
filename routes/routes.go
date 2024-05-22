package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/syahlan1/golos/controllers/aplicantController"
	"github.com/syahlan1/golos/controllers/approvalController"
	"github.com/syahlan1/golos/controllers/authController"
	"github.com/syahlan1/golos/controllers/businessController"
	"github.com/syahlan1/golos/controllers/masterCodeController"
	"github.com/syahlan1/golos/controllers/masterColumnController"
	"github.com/syahlan1/golos/controllers/masterTableController"
	"github.com/syahlan1/golos/controllers/ownershipDataController"
	"github.com/syahlan1/golos/controllers/validationController"
	"github.com/syahlan1/golos/middleware"
)

func Setup(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	app.Use(logger.New())

	api := app.Group("/api")

	api.Post("/register", authController.Register)
	api.Post("/login", authController.Login)
	api.Get("/user", authController.User)
	api.Post("/logout", authController.Logout)

	//business
	api.Post("/business/create", middleware.Authorize("create"), businessController.BusinessCreate)
	api.Get("/business/show", businessController.BusinessShow)
	api.Put("/business/update/:id", middleware.Authorize("update"), businessController.BusinessUpdate)
	api.Put("/business/delete/:id", middleware.Authorize("delete"), businessController.BusinessDelete)
	api.Get("/business/show/:id", businessController.BusinessShowDetail)
	api.Get("/business/showcompanyfirstname", businessController.ShowCompanyFirstName)
	api.Get("/business/showcompanytype", businessController.ShowCompanyType)
	api.Get("/business/showbusinessaddresstype", businessController.ShowBusinessAddressType)
	api.Get("/business/showratingclass", businessController.ShowRatingClass)
	api.Get("/business/showkodebursa", businessController.ShowKodeBursa)
	api.Get("/business/showbusinesstype", businessController.ShowBusinessType)

	//ownership
	api.Get("/ownership/show", ownershipDataController.ShowOwnershipData)
	api.Post("/ownership/create", ownershipDataController.CreateOwnershipData)
	api.Put("/ownership/update/:id", ownershipDataController.EditOwnershipData)
	api.Put("/ownership/delete/:id", ownershipDataController.DeleteOwnershipData)

	//relation with bank
	api.Post("/relation-with-bank/create/:id", ownershipDataController.CreateRelationWithBank)
	api.Get("/relation-with-bank/show", ownershipDataController.ShowRelationWithBank)
	api.Put("/relation-with-bank/update/:id", ownershipDataController.UpdateRelationWithBank)
	api.Put("/relation-with-bank/delete/:id", ownershipDataController.DeleteRelationWithBank)

	//Rekening Debitur
	api.Get("/rekening-debitur/show", ownershipDataController.ShowRekeningDebitur)
	api.Put("/rekening-debitur/update/:id", ownershipDataController.UpdateRekeningDebitur)
	api.Put("/rekening-debitur/delete/:id", ownershipDataController.DeleteRekeningDebitur)

	api.Get("/show-business-applicant", businessController.ShowBusinessApplicant)

	//applicant
	api.Post("/applicant/create", middleware.Authorize("create"), aplicantController.ApplicantCreate)
	api.Get("/applicant/show", aplicantController.ApplicantShow)
	api.Get("/applicant/show/:id", aplicantController.ApplicantShowDetail)
	api.Put("/applicant/update/:id", middleware.Authorize("update"), aplicantController.ApplicantUpdate)
	api.Delete("/applicant/delete/:id", middleware.Authorize("delete"), aplicantController.ApplicantDelete)
	api.Get("/applicant/show-homestatus", aplicantController.ShowHomeStatus)
	api.Get("/applicant/show-applicant-addresstype", aplicantController.ShowApplicantAddressType)
	api.Get("/applicant/show-education", aplicantController.ShowEducation)
	api.Get("/applicant/show-job-position", aplicantController.ShowJobPosition)
	api.Get("/applicant/show-business-sector", aplicantController.ShowBusinessSector)
	api.Get("/applicant/show-kode-instansi", aplicantController.ShowKodeInstansi)
	api.Get("/applicant/show-negara", aplicantController.ShowNegara)
	api.Get("/applicant/show-sektor-ekonomi", aplicantController.ShowSektorEkonomi)
	api.Get("/applicant/show-hubungan-nasabah", aplicantController.ShowHubunganNasabah)
	api.Get("/applicant/show-hubungan-keluarga", aplicantController.ShowHubunganKeluarga)
	api.Get("/applicant/show-lokasi-pabrik", aplicantController.ShowLokasiPabrik)
	api.Get("/applicant/marital-status", aplicantController.ShowMaritalStatus)

	//zipcode
	api.Get("/provinces", businessController.GetProvinces)
	api.Get("/cities", businessController.GetCitiesByProvince)
	api.Get("/districts", businessController.GetDistrictByCity)
	api.Get("/subdistricts", businessController.GetSubdistrictByDistrict)
	api.Get("/zip-codes", businessController.GetZipCodesBySubdistrict)

	//approve
	api.Post("/approval/create", approvalController.CreateApprovalSetting)
	api.Put("/approval/:id", approvalController.UpdateApprovalStatus)
	api.Put("/approval/:id/reject", approvalController.RejectApproval)
	api.Get("/approval/show", approvalController.ShowAllData)
	api.Put("/approval/set-role/:id", approvalController.UpdateApprovalWorkflowRoles)
	api.Get("/approval/data/:id", approvalController.ApprovalDataDetail)

	//role
	api.Post("/role/create", middleware.Authorize("create_role"), authController.CreateRole)
	api.Put("/role/update/:id", authController.UpdateRole)
	api.Get("/role/show", authController.ShowRole)
	api.Get("/role/:id/permissions", authController.ShowPermissions)
	api.Get("/role/permissions", authController.ShowAllPermissions)
	api.Delete("/role/delete/:id", authController.DeleteRole)

	//Validation
	api.Get("/validation/show", validationController.ShowAllValidations)
	// api.Get("/validation/show/:group_id", controllers.ShowDetailValidation)
	api.Put("/validation/delete/:id", validationController.DeleteValidation)
	api.Put("/validation/update/:id", validationController.UpdateValidation)
	api.Post("/validation/create", validationController.CreateValidation)

	//Master Code
	api.Get("/master-codes/show", masterCodeController.ShowMasterCode)
	api.Get("/master-codes/show/by-name/:code_group", masterCodeController.ShowDetailMasterCode)
	api.Get("/master-codes/show/by-id/:code_group_id", masterCodeController.ShowDetailMasterCode)
	api.Post("/master-codes/create", masterCodeController.CreateMasterCode)
	api.Put("/master-codes/edit/:id", masterCodeController.UpdateMasterCode)
	api.Put("/master-codes/delete/:id", masterCodeController.DeleteMasterCode)
	api.Get("/master-code-group/show", masterCodeController.ShowMasterCodeGroup)
	api.Post("/master-code-group/create", masterCodeController.CreateMasterCodeGroup)
	api.Put("/master-code-group/edit/:id", masterCodeController.UpdateMasterCodeGroup)
	api.Put("/master-code-group/delete/:id", masterCodeController.DeleteMasterCodeGroup)

	//Master Table
	api.Get("/master-table/show", masterTableController.ShowMasterTable)
	api.Get("/master-table/show/:id", masterTableController.ShowMasterTableDetail)
	api.Post("/master-table/create", masterTableController.CreateMasterTable)
	api.Put("/master-table/delete/:id", masterTableController.DeleteMasterTable)
	api.Put("/master-table/update/:id", masterTableController.UpdateMasterTable)

	// //Master Column
	api.Get("/master-column/show", masterColumnController.ShowMasterColumn)
	api.Get("/master-column/show/:id", masterColumnController.ShowMasterColumnDetail)
	api.Get("/master-column/by-table/:id", masterColumnController.ShowMasterColumnByTable)
	api.Post("/master-column/create/:id", masterColumnController.CreateMasterColumn)
	api.Put("/master-column/delete/:id", masterColumnController.DeleteMasterColumn)
	api.Put("/master-column/:id", masterColumnController.UpdateColumnTable)

	// //generate table
	api.Post("/master-table/generate/:id", masterTableController.GenerateTable)
}
