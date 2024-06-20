package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/syahlan1/golos/controllers/aplicantController"
	"github.com/syahlan1/golos/controllers/approvalController"
	"github.com/syahlan1/golos/controllers/authController"
	"github.com/syahlan1/golos/controllers/businessController"
	"github.com/syahlan1/golos/controllers/generalInformationController"
	"github.com/syahlan1/golos/controllers/idCardController"
	"github.com/syahlan1/golos/controllers/masterCodeController"
	"github.com/syahlan1/golos/controllers/masterColumnController"
	"github.com/syahlan1/golos/controllers/masterParameterController"
	"github.com/syahlan1/golos/controllers/masterTableController"
	"github.com/syahlan1/golos/controllers/ownershipDataController"
	"github.com/syahlan1/golos/controllers/sectorEconomyController"
	"github.com/syahlan1/golos/controllers/validationController"
	"github.com/syahlan1/golos/middleware"
)

func Setup(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://192.168.100.32:5173, http://localhost:5173",
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
	api.Put("/user/update/:id", authController.UpdateUser)
	api.Put("/user/delete/:id", authController.DeleteUser)
	api.Put("/user/change-password/:id", authController.ChangePassword)

	//business
	api.Post("/business/create", middleware.Authorize("create"), businessController.BusinessCreate)
	api.Get("/business/show", businessController.BusinessShow)
	api.Put("/business/update/:id", middleware.Authorize("update"), businessController.BusinessUpdate)
	api.Delete("/business/delete/:id", middleware.Authorize("delete"), businessController.BusinessDelete)
	api.Post("/business/upload", businessController.BusinessUploadFile)
	api.Get("/business/file/:id", businessController.BusinessShowFile)
	api.Get("/business/show/:id", businessController.BusinessShowDetail)
	api.Get("/business/show-company-first-name", businessController.ShowCompanyFirstName)
	api.Get("/business/show-company-type", businessController.ShowCompanyType)
	api.Get("/business/showbusinessaddresstype", businessController.ShowBusinessAddressType)
	api.Get("/business/show-external-rating-company", businessController.ShowExternalRatingCompany)
	api.Get("/business/show-rating-class", businessController.ShowRatingClass)
	api.Get("/business/show-kode-bursa", businessController.ShowKodeBursa)
	api.Get("/business/show-business-type", businessController.ShowBusinessType)

	//ownership
	api.Get("/ownership/show", ownershipDataController.ShowOwnershipData)
	api.Get("/ownership/name", ownershipDataController.ShowOwnershipName)
	api.Post("/ownership/create", ownershipDataController.CreateOwnershipData)
	api.Put("/ownership/update/:id", ownershipDataController.EditOwnershipData)
	api.Put("/ownership/delete/:id", ownershipDataController.DeleteOwnershipData)

	//relation with bank
	api.Post("/relation-with-bank/create", ownershipDataController.CreateRelationWithBank)
	api.Get("/relation-with-bank/show", ownershipDataController.ShowRelationWithBank)
	api.Put("/relation-with-bank/update/:id", ownershipDataController.UpdateRelationWithBank)
	api.Delete("/relation-with-bank/delete/:id", ownershipDataController.DeleteRelationWithBank)
	
	api.Post("/customer-loan-info/create", ownershipDataController.CreateCustomerLoanInfo)
	api.Get("/customer-loan-info/show", ownershipDataController.ShowCustomerLoanInfo)
	api.Put("/customer-loan-info/update/:id", ownershipDataController.UpdateCustomerLoanInfo)
	api.Delete("/customer-loan-info/delete/:id", ownershipDataController.DeleteCustomerLoanInfo)
	api.Get("/show-facility-no", ownershipDataController.ShowFacilityNo)
	api.Get("/show-product", ownershipDataController.ShowProduct)
	api.Get("/show-customer-aa", ownershipDataController.ShowCustomerAA)

	//Rekening Debitur
	api.Post("/rekening-debitur/create", ownershipDataController.CreateRekeningDebitur)
	api.Get("/rekening-debitur/show", ownershipDataController.ShowRekeningDebitur)
	api.Put("/rekening-debitur/update/:id", ownershipDataController.UpdateRekeningDebitur)
	api.Delete("/rekening-debitur/delete/:id", ownershipDataController.DeleteRekeningDebitur)

	api.Get("/show-business-applicant", businessController.ShowBusinessApplicant)

	//applicant
	api.Post("/applicant/create", middleware.Authorize("create"), aplicantController.ApplicantCreate)
	api.Get("/applicant/show", aplicantController.ApplicantShow)
	api.Get("/applicant/show/:id", aplicantController.ApplicantShowDetail)
	api.Put("/applicant/update/:id", middleware.Authorize("update"), aplicantController.ApplicantUpdate)
	api.Delete("/applicant/delete/:id", middleware.Authorize("delete"), aplicantController.ApplicantDelete)
	api.Post("/applicant/upload", aplicantController.ApplicantUploadFile)
	api.Get("/applicant/file/:id", aplicantController.ApplicantShowFile)
	api.Get("/applicant/show-homestatus", aplicantController.ShowHomeStatus)
	api.Get("/applicant/show-applicant-addresstype", aplicantController.ShowApplicantAddressType)
	api.Get("/applicant/show-education", aplicantController.ShowEducation)
	api.Get("/applicant/show-job-position", aplicantController.ShowJobPosition)
	api.Get("/applicant/show-business-sector", aplicantController.ShowBusinessSector)
	api.Get("/applicant/show-kode-instansi", aplicantController.ShowKodeInstansi)
	api.Get("/applicant/show-negara", aplicantController.ShowNegara)
	api.Get("/applicant/show-nationality", aplicantController.ShowNationality)
	api.Get("/applicant/show-gender", aplicantController.ShowGender)
	api.Get("/applicant/show-sektor-ekonomi", aplicantController.ShowSektorEkonomi)
	api.Get("/applicant/marital-status", aplicantController.ShowMaritalStatus)

	api.Get("/address-type", idCardController.ShowAddressType)

	api.Get("/sektor-ekonomi-1", sectorEconomyController.ShowSektorEkonomi1)
	api.Get("/sektor-ekonomi-2", sectorEconomyController.ShowSektorEkonomi2)
	api.Get("/sektor-ekonomi-3", sectorEconomyController.ShowSektorEkonomi3)
	api.Get("/sektor-ekonomi-ojk", sectorEconomyController.ShowSektorEkonomiOjk)
	api.Get("/lokasi-pabrik", sectorEconomyController.ShowLokasiPabrik)
	api.Get("/lokasi-dati2", sectorEconomyController.ShowLokasiDati2)
	api.Get("/hubungan-nasabah", sectorEconomyController.ShowHubunganNasabahBank)
	api.Get("/hubungan-keluarga", sectorEconomyController.ShowHubunganKeluarga)

	api.Get("/cabang-pencairan", generalInformationController.ShowCabangPencairan)
	api.Get("/cabang-admin", generalInformationController.ShowCabangAdmin)
	api.Get("/segment", generalInformationController.ShowSegment)
	api.Get("/program", generalInformationController.ShowProgram)
	api.Get("/application-number/:id", generalInformationController.GenerateApplicationNumber)

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

	// //master parameter
	api.Post("/master-parameter/create", masterParameterController.CreateParameter)
	api.Put("/master-parameter/update/:id", masterParameterController.UpdateMasterParameter)
	api.Put("/master-parameter/delete/:id", masterParameterController.DeleteMasterParameter)
	api.Get("/master-parameter/show", masterParameterController.ShowAllParameter)
	api.Get("/master-parameter/show/:id", masterParameterController.ShowParameterDetail)
}
