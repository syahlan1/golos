package models

import "time"

type Approval struct {
	Id                   string    `gorm:"primaryKey"`
	ApprovalSettingID    int       `json:"approval_setting_id"`
	OwnerID              string    `json:"owner_id"`
	Data                 string    `json:"data" gorm:"type:text"`
	CurrentProcess       int       `json:"current_process"`
	CurrentNotes         string    `json:"current_notes" gorm:"type:text"`
	CreatedDate          time.Time `json:"created_date"`
	CreatedBy            string    `json:"created_by"`
	UpdatedDate          time.Time `json:"updated_date"`
	UpdatedBy            string    `json:"updated_by"`
	ApprovalStatus       string    `json:"approval_status"`
	FormNumber           string    `json:"form_number"`
	FormDocument         string    `json:"form_document"`
	FormDocumentOriginal string    `json:"form_document_original"`
	AgentBranchID        int       `json:"agent_branch_id"`
	DisplayData          string    `json:"display_data" gorm:"type:text"`
	Checklist            string    `json:"checklist" gorm:"type:text"`
	ProcessDate          time.Time `json:"process_date"`
}

type ApprovalSetting struct {
	Id                      int       `gorm:"primaryKey"`
	Module                  string    `json:"module"`
	Type                    string    `json:"type"`
	Description             string    `json:"description"`
	Status                  string    `json:"status"`
	CreatedDate             time.Time `json:"created_date"`
	CreatedBy               string    `json:"created_by"`
	UpdatedDate             time.Time `json:"updated_date"`
	UpdatedBy               string    `json:"updated_by"`
	ServiceCallback         string    `json:"service_callback"`
	ControllerClassName     string    `json:"controller_class_name"`
	FormEditFileName        string    `json:"form_edit_file_name"`
	FormApprovalFileName    string    `json:"form_approval_file_name"`
	RejectClose             bool      `json:"reject_close"`
	DisplayFields           string    `json:"display_fields" gorm:"type:text"`
	FormUploadRequired      bool      `json:"form_upload_required"`
	RecordUniqueKeys        string    `json:"record_unique_keys" gorm:"type:text"`
	RecordValidationService string    `json:"record_validation_service"`
	ServiceNotFound         string    `json:"service_not_found"`
	Checklist               string    `json:"checklist" gorm:"type:text"`
	BranchFilerRoles        string    `json:"branch_filter_roles" gorm:"type:text"`
	BranchFilterAdditionals string    `json:"branch_filter_additionals" gorm:"type:text"`
	ActionCallback          string    `json:"action_callback" gorm:"type:text"`
	OverrideFields          string    `json:"override_fields"`
	ShowOverrideInApproval  bool      `json:"show_override_in_approval"`
	ShowOverrideHistories   bool      `json:"show_override_histories"`
	UsingProcessDate        bool      `json:"using_process_date"`
}

type ShowApproval struct {
	Id             string `json:"id"`
	DisplayData    string `json:"display_data"`
	ApprovalStatus string `json:"approval_status"`
	CreatedBy      string `json:"created_by"`
	CreatedDate    string `json:"created_date"`
	Description    string `json:"description"`
	Module         string `json:"module"`
	Type           string `json:"type"`
}

type CreateApprovalSetting struct {
	Module            string                  `json:"module"`
	Type              string                  `json:"type"`
	Description       string                  `json:"description"`
	ApprovalWorkflows []CrateApprovalWorkflow `json:"approval_workflows"`
}

type CrateApprovalWorkflow struct {
	Name          string `json:"name"`
	ProcessStatus string `json:"process_status"`
	Order         int    `json:"order"`
	RoleID        []int  `json:"role_id"`
	IsFirstStep   bool   `json:"is_first_step"`
}

type ApprovalWorkflow struct {
	Id                     int       `gorm:"primary_key" json:"id"`
	ApprovalSettingID      int       `json:"approval_setting_id"`
	Name                   string    `json:"name"`
	Description            string    `json:"description"`
	ProcessStatus          string    `json:"process_status"`
	Order                  int       `json:"order"`
	Status                 string    `json:"status"`
	CreatedDate            time.Time `json:"created_date"`
	CreatedBy              string    `json:"created_by"`
	UpdatedDate            time.Time `json:"updated_date"`
	UpdatedBy              string    `json:"updated_by"`
	IsFirstStep            bool      `json:"is_first_step"`
	Conditions             string    `json:"conditions" gorm:"type:text"`
	CanSettingUser         bool      `json:"can_setting_user"`
	IsShowChecklists       bool      `json:"is_show_checklists"`
	GroupID                int       `json:"group_id"`
	OverrideFields         string    `json:"override_fields"`
	BindingFieldNames      string    `json:"binding_field_names" gorm:"type:text"`
	CanClearPending        bool      `json:"can_clear_pending"`
	CanOverrideProcessDate bool      `json:"can_override_process_date"`
	CanFilterProcessDate   bool      `json:"can_filter_process_date"`
	IsRequiredOverrides    bool      `json:"is_required_overrides"`
}

type ApprovalWorkflowRole struct {
	Id                 uint      `gorm:"primary_key" json:"id"`
	ApprovalWorkflowID int       `json:"approval_workflow_id"`
	RoleID             int       `json:"role_id"`
	Status             string    `json:"status"`
	CreatedDate        time.Time `json:"created_date"`
	CreatedBy          string    `json:"created_by"`
	UpdatedDate        time.Time `json:"updated_date"`
	UpdatedBy          string    `json:"updated_by"`
	Actions            string    `json:"actions" gorm:"type:text"`
	Conditions         string    `json:"conditions" gorm:"type:text"`
	UserApprover       string    `json:"user_approver"`
}

type UpdateApprovalWorkflowRoles struct {
	RoleIDs []int `json:"role_id"`
}

type ApprovalHistory struct {
	Id             string    `gorm:"primaryKey" json:"id"`
	ApprovalID     string    `json:"approval_id"`
	Date           time.Time `json:"date"`
	UserID         string    `json:"user_id"`
	Status         string    `json:"status"`
	Notes          string    `json:"notes" gorm:"type:text"`
	CurrentProcess int       `json:"current_process"`
	Data           string    `json:"data" gorm:"type:text"`
	Checklist      string    `json:"checklist" gorm:"type:text"`
}
