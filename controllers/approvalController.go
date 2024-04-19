package controllers

import (
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/oklog/ulid"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"gorm.io/gorm"
)

func TakeUsername(c *fiber.Ctx) (string, error) {
	cookie := c.Cookies("jwt")

	// Memverifikasi token dan mendapatkan klaim
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	// Mendapatkan data pengguna (user) dari database
	var user models.Users
	if err := connection.DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		return "", err
	}
	return user.Username, nil
}

func CreateApprovalSetting(c *fiber.Ctx) error {
	var data map[string]interface{}
	currentTime := time.Now()

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	createdBy, err := TakeUsername(c)
	if err != nil {
		return err
	}

	approvalSetting := models.ApprovalSetting{
		Module:      data["module"].(string),
		Type:        data["type"].(string),
		Description: data["description"].(string),
		Status:      "active",
		CreatedDate: currentTime,
		CreatedBy:   createdBy,
	}

	if err := connection.DB.Create(&approvalSetting).Error; err != nil {
		return err
	}

	approvalWorkflowsData := data["approval_workflows"].([]interface{})
	for _, wfData := range approvalWorkflowsData {
		wf := wfData.(map[string]interface{})
		approvalWorkflow := models.ApprovalWorkflow{
			ApprovalSettingID: approvalSetting.Id,
			Name:              wf["name"].(string),
			ProcessStatus:     wf["process_status"].(string),
			Order:             int(wf["order"].(float64)),
			CreatedDate:       currentTime,
			CreatedBy:         createdBy,
		}

		if err := connection.DB.Create(&approvalWorkflow).Error; err != nil {
			return err
		}

		// Get Role IDs from request
		roleIds := wf["role_id"].([]interface{})
		for _, roleId := range roleIds {
			// Create ApprovalWorkflowRole for each Role ID
			approvalWorkflowRole := models.ApprovalWorkflowRole{
				ApprovalWorkflowID: int(approvalWorkflow.Id),
				RoleID:             int(roleId.(float64)),
				Status:             "active",
				CreatedDate:        currentTime,
				CreatedBy:          createdBy,
			}

			if err := connection.DB.Create(&approvalWorkflowRole).Error; err != nil {
				return err
			}
		}
	}

	return c.JSON(fiber.Map{
		"message": "insert sukses",
	})
}

func UpdateApprovalStatus(c *fiber.Ctx) error {
	approvalID := c.Params("id")
	createdBy, err := TakeUsername(c)
	if err != nil {
		log.Println("Error taking username:", err)
		return err
	}

	// Get user role ID
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		log.Println("Error parsing JWT:", err)
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "status unauthorized",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.Users
	if err := connection.DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	// Get the Approval to be updated
	var approval models.Approval
	if err := connection.DB.Where("id = ?", approvalID).First(&approval).Error; err != nil {
		log.Println("Error retrieving approval:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Data Not Found",
		})
	}

	// Get the ApprovalWorkflow based on the current process of the Approval
	var currentWorkflow models.ApprovalWorkflow
	if err := connection.DB.Where("approval_setting_id = ? AND id = ?", approval.ApprovalSettingID, approval.CurrentProcess).First(&currentWorkflow).Error; err != nil {
		log.Println("Error retrieving current workflow:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "No current approval workflow found",
		})
	}

	// Check if the user has permission to access the current ApprovalWorkflow
	var workflowRole models.ApprovalWorkflowRole
	if err := connection.DB.Where("approval_workflow_id = ? AND role_id = ? AND status = ?", currentWorkflow.Id, user.RoleId, "active").First(&workflowRole).Error; err != nil {
		log.Println("User does not have permission:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User does not have permission to access the current ApprovalWorkflow",
		})
	}

	// Get the next ApprovalWorkflow
	var nextWorkflow models.ApprovalWorkflow
	if err := connection.DB.Where("approval_setting_id = ? AND \"order\" > ?", approval.ApprovalSettingID, currentWorkflow.Order).Order("\"order\" asc").First(&nextWorkflow).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("No next approval workflow found:", err)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "No next approval workflow found",
			})
		}
		log.Println("Error retrieving next workflow:", err)
		return err
	}

	// Update current_process on Approval with the id of the found ApprovalWorkflow
	approval.CurrentProcess = nextWorkflow.Id
	approval.UpdatedDate = time.Now()
	approval.UpdatedBy = createdBy

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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to update the approval data",
		})
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
		log.Println("Error creating approval history:", err)
		return err
	}

	log.Println("Approval updated successfully")
	return c.JSON(fiber.Map{
		"data":   approval,
		"status": "Updated!",
	})
}

func RejectApproval(c *fiber.Ctx) error {
	approvalId := c.Params("id")

	var approval models.Approval
	if err := connection.DB.Where("id = ?", approvalId).First(&approval).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Data Not Found",
		})
	}

	var updatedApproval models.Approval
	if err := c.BodyParser(&updatedApproval); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid Data",
		})
	}

	// Dapatkan ApprovalSettingID dari Approval yang akan diperbarui
	approvalSettingID := approval.ApprovalSettingID

	// Dapatkan ApprovalWorkflow dengan urutan pertama
	var firstWorkflow models.ApprovalWorkflow
	if err := connection.DB.Where("approval_setting_id = ?", approvalSettingID).Order("\"order\" asc").First(&firstWorkflow).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "No first approval workflow found",
			})
		}
		return err
	}

	// Perbarui current_process pada Approval dengan urutan ApprovalWorkflow pertama yang ditemukan
	approval.CurrentProcess = firstWorkflow.Id
	approval.CurrentNotes = updatedApproval.CurrentNotes

	// Set approval_status menjadi "rejected"
	approval.ApprovalStatus = "rejected"

	if err := connection.DB.Save(&approval).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to reject the approval",
		})
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
		return err
	}

	return c.JSON(fiber.Map{
		"data":   approval,
		"status": "Approval rejected!",
	})
}

func ShowAllData(c *fiber.Ctx) error {
	var approvals []struct {
		Id             string `json:"id"`
		DisplayData    string `json:"display_data"`
		ApprovalStatus string `json:"approval_status"`
		CreatedBy      string `json:"created_by"`
		CreatedDate    string `json:"created_date"`
		Description    string `json:"description"`
		Module         string `json:"module"`
		Type           string `json:"type"`
	}

	// Query menggunakan raw SQL untuk melakukan join antara tabel Approval dan ApprovalWorkflow
	query := `SELECT a.id, a.display_data, a.approval_status, a.created_by, a.created_date, aw.description, ast.module, ast.type, ast.type
              FROM approvals a
              JOIN approval_workflows aw ON a.current_process = aw.id
			  JOIN approval_settings ast ON a.approval_setting_id = ast.id`

	if err := connection.DB.Raw(query).Scan(&approvals).Error; err != nil {
		return err
	}

	return c.JSON(approvals)
}

func ApprovalDataDetail(c *fiber.Ctx) error {
	approvalID := c.Params("id")

	var data string
	if err := connection.DB.Raw("SELECT data FROM approvals WHERE id = ?", approvalID).Scan(&data).Error; err != nil {
		return err
	}

	return c.SendString(data)
}

func UpdateApprovalWorkflowRoles(c *fiber.Ctx) error {
	// Parse request body
	var req struct {
		RoleIDs []int `json:"role_id"`
	}
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	// Get approval_workflow_id from URL parameter
	approvalWorkflowID := c.Params("id")

	// Find existing approval workflow by ID
	var existingWorkflow models.ApprovalWorkflow
	if err := connection.DB.Where("id = ?", approvalWorkflowID).First(&existingWorkflow).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Approval workflow not found"})
	}

	// Update approval workflow roles
	if err := updateApprovalWorkflowRoles(existingWorkflow.Id, req.RoleIDs); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Approval workflow roles updated successfully"})
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
