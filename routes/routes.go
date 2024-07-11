package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/syahlan1/golos/controllers/masterTableGroupController"
	"github.com/syahlan1/golos/controllers/masterTemplateController"

	"github.com/syahlan1/golos/controllers/aplicantController"
	"github.com/syahlan1/golos/controllers/approvalController"
	"github.com/syahlan1/golos/controllers/authController"
	"github.com/syahlan1/golos/controllers/businessController"
	"github.com/syahlan1/golos/controllers/generalInformationController"
	"github.com/syahlan1/golos/controllers/idCardController"
	"github.com/syahlan1/golos/controllers/masterCodeController"
	"github.com/syahlan1/golos/controllers/masterColumnController"
	"github.com/syahlan1/golos/controllers/masterModuleController"
	"github.com/syahlan1/golos/controllers/masterParameterController"
	"github.com/syahlan1/golos/controllers/masterTableController"
	"github.com/syahlan1/golos/controllers/masterWorkflowController"
	"github.com/syahlan1/golos/controllers/menuController"
	"github.com/syahlan1/golos/controllers/ownershipDataController"
	"github.com/syahlan1/golos/controllers/roleController"
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

	//route tes_table_1
	// api.Post("/tes_table_1/create", tesTable1Controller.CreateTesTable1)
	// api.Put("/tes_table_1/update/:id", tesTable1Controller.UpdateTesTable1)
	// api.Put("/tes_table_1/delete/:id", tesTable1Controller.DeleteTesTable1)
	// api.Get("/tes_table_1/show", tesTable1Controller.ShowTesTable1)
	// api.Get("/tes_table_1/show/:id", tesTable1Controller.ShowDetailTesTable1)

	api.Post("/register", authController.Register)
	api.Post("/login", authController.Login)
	api.Get("/user", authController.User)
	api.Get("/user/permission", authController.UserPermission)
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
	api.Post("/role/create", middleware.Authorize("create_role"), roleController.CreateRole)
	api.Put("/role/update/:id", roleController.UpdateRole)
	api.Get("/role/show", roleController.ShowRole)
	api.Get("/role/:id/permissions", roleController.ShowPermissions)
	api.Get("/role/permissions", roleController.ShowAllPermissions)
	api.Delete("/role/delete/:id", roleController.DeleteRole)
	api.Post("/role/modules", roleController.CreateRoleModules)
	api.Get("/role/:id/modules", roleController.ShowRoleModules)
	api.Post("/role/:id/module-tables", roleController.CreateRoleModuleTables)
	api.Get("/role/:id/menu", roleController.ShowAllRoleMenu)
	api.Post("/role/:id/menu", roleController.CreateRoleMenu)
	api.Get("/role/:id/workflow", roleController.ShowRoleWorkflows)
	api.Post("/role/:id/workflow", roleController.CreateRoleWorkflows)

	//menu
	api.Post("/menu/create", menuController.CreateMenu)
	api.Get("/menu/show", menuController.ShowMenu)

	//Validation
	api.Get("/validation/show", validationController.ShowAllValidations)
	// api.Get("/validation/show/:group_id", controllers.ShowDetailValidation)
	api.Delete("/validation/delete/:id", validationController.DeleteValidation)
	api.Put("/validation/update/:id", validationController.UpdateValidation)
	api.Post("/validation/create/:id", validationController.CreateValidation)

	//Master Code
	api.Get("/master-codes/show", masterCodeController.ShowMasterCode)
	api.Get("/master-codes/show/by-name/:code_group", masterCodeController.ShowDetailMasterCode)
	api.Get("/master-codes/show/by-id/:code_group_id", masterCodeController.ShowDetailMasterCode)
	api.Post("/master-codes/create", masterCodeController.CreateMasterCode)
	api.Put("/master-codes/update/:id", masterCodeController.UpdateMasterCode)
	api.Delete("/master-codes/delete/:id", masterCodeController.DeleteMasterCode)
	api.Get("/master-code-group/show", masterCodeController.ShowMasterCodeGroup)
	api.Post("/master-code-group/create", masterCodeController.CreateMasterCodeGroup)
	api.Put("/master-code-group/update/:id", masterCodeController.UpdateMasterCodeGroup)
	api.Delete("/master-code-group/delete/:id", masterCodeController.DeleteMasterCodeGroup)

	//Master Module
	api.Get("/master-module/show", masterModuleController.ShowMasterModule)
	api.Get("/master-module/show/:id", masterModuleController.ShowMasterModuleDetail)
	api.Post("/master-module/create", masterModuleController.CreateMasterModule)
	api.Put("/master-module/update/:id", masterModuleController.UpdateMasterModule)
	api.Delete("/master-module/delete/:id", masterModuleController.DeleteMasterModule)

	//Master Table
	api.Get("/master-table/show", masterTableController.ShowMasterTable)
	api.Get("/master-table/show/:id", masterTableController.ShowMasterTableDetail)
	api.Post("/master-table/create", masterTableController.CreateMasterTable)
	api.Delete("/master-table/delete/:id", masterTableController.DeleteMasterTable)
	api.Put("/master-table/update/:id", masterTableController.UpdateMasterTable)

	//Master Table Group
	api.Get("/master-table-group/show", masterTableGroupController.ShowMasterTableGroup)
	api.Get("/master-table-group/show/:id", masterTableGroupController.ShowMasterTableGroupDetail)
	api.Post("/master-table-group/create", masterTableGroupController.CreateMasterTableGroup)
	api.Put("/master-table-group/update/:id", masterTableGroupController.UpdateMasterTableGroup)
	api.Delete("/master-table-group/delete/:id", masterTableGroupController.DeleteMasterTableGroup)

	//Master Table Item
	api.Get("/master-table-item/show", masterTableGroupController.ShowMasterTableItem)
	api.Get("/master-table-item/show/:id", masterTableGroupController.ShowMasterTableItemDetail)
	api.Post("/master-table-item/create", masterTableGroupController.CreateMasterTableItem)
	api.Put("/master-table-item/update/:id", masterTableGroupController.UpdateMasterTableItem)
	api.Delete("/master-table-item/delete/:id", masterTableGroupController.DeleteMasterTableItem)

	api.Get("/data-master-table-group/:table_group/:table_item/show", masterTableGroupController.ShowDataMasterTableGroup)
	api.Get("/data-master-table-group/:table_group/:table_item/show/:id", masterTableGroupController.ShowDataMasterTableGroupById)
	api.Post("/data-master-table-group/:table_group/:table_item/create", masterTableGroupController.CreateDataMasterTableGroup)
	api.Put("/data-master-table-group/:table_group/:table_item/update", masterTableGroupController.UpdateDataMasterTableGroup)
	api.Delete("/data-master-table-group/:table_group/:table_item/delete", masterTableGroupController.DeleteDataMasterTableGroup)

	api.Get("/form-master-table-group/:group_name/show", masterTableGroupController.ShowFormMasterTableGroup)
	api.Get("/approval-master-table-group/:group_name/show", masterTableGroupController.ShowApprovalTableGroupItem)
	api.Post("/approval-master-table-group", masterTableGroupController.ApprovalTableGroupItem)
	api.Get("/approval-master-table-group/:group_name/show/:id", masterTableGroupController.ShowDetailApprovalTableGroupItem)
	api.Post("/submit-master-table-group", masterTableGroupController.SubmitTableGroupItem)

	// //Master Column
	api.Get("/master-column/field-type", masterColumnController.GetFieldType)
	api.Get("/master-column/ui-type", masterColumnController.GetUiType)
	api.Get("/master-column/show-column/:id", masterColumnController.ShowFormColumn)
	api.Get("/master-column/show", masterColumnController.ShowMasterColumn)
	api.Get("/master-column/show/:id", masterColumnController.ShowMasterColumnDetail)
	api.Get("/master-column/by-table/:id", masterColumnController.ShowMasterColumnByTable)
	api.Post("/master-column/create/:id", masterColumnController.CreateMasterColumn)
	api.Delete("/master-column/delete/:id", masterColumnController.DeleteMasterColumn)
	api.Put("/master-column/:id", masterColumnController.UpdateColumnTable)
	api.Post("/master-column/check-query", masterColumnController.CheckQuery)

	// //generate table
	api.Post("/master-table/generate/:id", masterTableController.GenerateTable)
	api.Post("/master-table-group/generate/:id", masterTableGroupController.GenerateTableGroup)

	// //master parameter
	api.Post("/master-parameter/create", masterParameterController.CreateParameter)
	api.Put("/master-parameter/update/:id", masterParameterController.UpdateMasterParameter)
	api.Delete("/master-parameter/delete/:id", masterParameterController.DeleteMasterParameter)
	api.Get("/master-parameter/show", masterParameterController.ShowAllParameter)
	api.Get("/master-parameter/show/:id", masterParameterController.ShowParameterDetail)

	// master workflow
	api.Post("/master-workflow/create", masterWorkflowController.CreateMasterWorkflow)
	api.Put("/master-workflow/update/:id", masterWorkflowController.UpdateMasterWorkflow)
	api.Delete("/master-workflow/delete/:id", masterWorkflowController.DeleteMasterWorkflow)
	api.Get("/master-workflow/show", masterWorkflowController.ShowMasterWorkflow)

	api.Get("/master-template/:module/:table/show", masterTemplateController.ShowMasterTemplate)
	api.Get("/master-template/:module/:table/show/:id", masterTemplateController.ShowMasterTemplateById)
	api.Post("/master-template/:module/:table/Create", masterTemplateController.CreateMasterTemplate)
	api.Put("/master-template/:module/:table/update/:id", masterTemplateController.UpdateMasterTemplate)
	api.Delete("/master-template/:module/:table/delete/:id", masterTemplateController.DeleteMasterTemplate)
}
