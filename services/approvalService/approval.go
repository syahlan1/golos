package approvalService

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"gorm.io/gorm"
)

func CreateApprovalSetting(data models.CreateApprovalSetting, username string) (err error) {
	currentTime := time.Now()

	approvalSetting := models.ApprovalSetting{
		Module:      data.Module,
		Type:        data.Type,
		Description: data.Description,
		Status:      "active",
		CreatedDate: currentTime,
		CreatedBy:   username,
	}

	if err := connection.DB.Create(&approvalSetting).Error; err != nil {
		return err
	}

	approvalWorkflowsData := data.ApprovalWorkflows
	for _, wfData := range approvalWorkflowsData {
		approvalWorkflow := models.ApprovalWorkflow{
			ApprovalSettingID: approvalSetting.Id,
			Name:              wfData.Name,
			ProcessStatus:     wfData.ProcessStatus,
			Order:             wfData.Order,
			IsFirstStep:       wfData.IsFirstStep,
			CreatedDate:       currentTime,
			CreatedBy:         username,
		}

		if err := connection.DB.Create(&approvalWorkflow).Error; err != nil {
			return err
		}

		// Get Role IDs from request
		roleIds := wfData.RoleID
		for _, roleId := range roleIds {
			// Create ApprovalWorkflowRole for each Role ID
			approvalWorkflowRole := models.ApprovalWorkflowRole{
				ApprovalWorkflowID: approvalWorkflow.Id,
				RoleID:             roleId,
				Status:             "active",
				CreatedDate:        currentTime,
				CreatedBy:          username,
			}

			if err := connection.DB.Create(&approvalWorkflowRole).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func UpdateApprovalStatus(userID, approvalID string) (result models.Approval, err error) {
	var user models.Users
	if err := connection.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		log.Println("error retrieving user:", err)
		return result, err
	}

	// Get the Approval to be updated
	var approval models.Approval
	if err := connection.DB.Where("id = ?", approvalID).First(&approval).Error; err != nil {
		log.Println("Error retrieving approval:", err)
		return result, errors.New("data Not Found")
	}

	// Get the ApprovalWorkflow based on the current process of the Approval
	var currentWorkflow models.ApprovalWorkflow
	if err := connection.DB.Where("approval_setting_id = ? AND id = ?", approval.ApprovalSettingID, approval.CurrentProcess).First(&currentWorkflow).Error; err != nil {
		log.Println("Error retrieving current workflow:", err)
		return result, errors.New("no current approval workflow found")
	}

	// Check if the user has permission to access the current ApprovalWorkflow
	var workflowRole models.ApprovalWorkflowRole
	if err := connection.DB.Where("approval_workflow_id = ? AND role_id = ? AND status = ?", currentWorkflow.Id, user.RoleId, "active").First(&workflowRole).Error; err != nil {
		log.Println("User does not have permission:", err)
		return result, errors.New("user does not have permission to access the current ApprovalWorkflow")
	}

	// Get the next ApprovalWorkflow
	var nextWorkflow models.ApprovalWorkflow
	if err := connection.DB.Where("approval_setting_id = ? AND \"order\" > ?", approval.ApprovalSettingID, currentWorkflow.Order).Order("\"order\" asc").First(&nextWorkflow).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("No next approval workflow found:", err)
			return result, errors.New("no next approval workflow found")
		}
		log.Println("Error retrieving next workflow:", err)
		return result, err
	}

	// Update current_process on Approval with the id of the found ApprovalWorkflow
	approval.CurrentProcess = nextWorkflow.Id
	approval.UpdatedDate = time.Now()
	approval.UpdatedBy = user.Username

	// Determine approval_status based on process_status
	switch nextWorkflow.ProcessStatus {
	case "draft":
		approval.ApprovalStatus = "draft"
	case "open":
		approval.ApprovalStatus = "pending"
	case "closed":
		approval.ApprovalStatus = "approved"
	default:
		// If process_status does not match the expected ones, set approval_status to a default value or adjust it according to your application's needs
		approval.ApprovalStatus = "unknown"
	}

	if err := connection.DB.Save(&approval).Error; err != nil {
		log.Println("Error saving approval:", err)
		return result, errors.New("failed to update approval data")
	}

	// Approval history
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	history := models.ApprovalHistory{
		Id:             id.String(),
		ApprovalID:     approval.Id,
		Date:           time.Now(),
		Data:           approval.Data,
		UserID:         approval.CreatedBy,
		CurrentProcess: approval.CurrentProcess,
		Notes:          approval.CurrentNotes,
		Status:         approval.ApprovalStatus,
	}

	if err := connection.DB.Create(&history).Error; err != nil {
		log.Println("error creating approval history:", err)
		return result, err
	}

	log.Println("approval updated successfully")
	return approval, nil
}

func RejectApproval(approvalId string, data models.Approval) (result models.Approval, err error) {
	var approval models.Approval
	if err := connection.DB.Where("id = ?", approvalId).First(&approval).Error; err != nil {
		return result, errors.New("data not found")
	}

	// Dapatkan ApprovalSettingID dari Approval yang akan diperbarui
	approvalSettingID := approval.ApprovalSettingID

	// Dapatkan ApprovalWorkflow dengan urutan pertama
	var firstWorkflow models.ApprovalWorkflow
	if err := connection.DB.Where("approval_setting_id = ?", approvalSettingID).Order("\"order\" asc").First(&firstWorkflow).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return result, errors.New("no first approval workflow found")
		}
		return result, err
	}

	// Perbarui current_process pada Approval dengan urutan ApprovalWorkflow pertama yang ditemukan
	approval.CurrentProcess = firstWorkflow.Id
	approval.CurrentNotes = data.CurrentNotes

	// Set approval_status menjadi "rejected"
	approval.ApprovalStatus = "rejected"

	if err := connection.DB.Save(&approval).Error; err != nil {
		return result, errors.New("failed to reject approval data")
	}

	//history approval
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	history := models.ApprovalHistory{
		Id:             id.String(),
		ApprovalID:     approval.Id,
		Date:           time.Now(),
		Data:           approval.Data,
		UserID:         approval.CreatedBy,
		CurrentProcess: approval.CurrentProcess,
		Notes:          approval.CurrentNotes,
		Status:         approval.ApprovalStatus,
	}

	if err := connection.DB.Create(&history).Error; err != nil {
		return result, err
	}

	return approval, nil
}

func ShowAllData() (result []models.ShowApproval, err error) {

	// Query menggunakan raw SQL untuk melakukan join antara tabel Approval dan ApprovalWorkflow
	query := `SELECT a.id, a.display_data, a.approval_status, a.created_by, a.created_date, aw.description, ast.module, ast.type, ast.type
              FROM approvals a
              JOIN approval_workflows aw ON a.current_process = aw.id
			  JOIN approval_settings ast ON a.approval_setting_id = ast.id`

	if err := connection.DB.Raw(query).Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func ApprovalDataDetail(approvalID string) (result string, err error) {
	if err := connection.DB.Raw("SELECT data FROM approvals WHERE id = ?", approvalID).Scan(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func UpdateApprovalWorkflowRoles(approvalWorkflowID string, data models.UpdateApprovalWorkflowRoles) (err error) {
	var existingWorkflow models.ApprovalWorkflow
	if err := connection.DB.Where("id = ?", approvalWorkflowID).First(&existingWorkflow).Error; err != nil {
		return errors.New("approval workflow not found")
	}

	// Update approval workflow roles
	if err := updateApprovalWorkflowRoles(existingWorkflow.Id, data.RoleIDs); err != nil {
		return err
	}

	return nil
}

func updateApprovalWorkflowRoles(approvalWorkflowID int, roleIDs []int) error {
	// Delete existing ApprovalWorkflowRoles entries for the approval_workflow_id
	if err := connection.DB.Where("approval_workflow_id = ?", approvalWorkflowID).Delete(&models.ApprovalWorkflowRole{}).Error; err != nil {
		return err
	}

	// Create new ApprovalWorkflowRoles entries for the approval_workflow_id and role_ids
	var workflowRoles []models.ApprovalWorkflowRole
	for _, roleID := range roleIDs {
		workflowRoles = append(workflowRoles, models.ApprovalWorkflowRole{
			ApprovalWorkflowID: approvalWorkflowID,
			RoleID:             roleID,
		})
	}
	if err := connection.DB.Create(&workflowRoles).Error; err != nil {
		return err
	}

	return nil
}