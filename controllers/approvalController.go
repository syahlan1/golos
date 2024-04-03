package controllers

import (
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

		// Create ApprovalWorkflowRole
		approvalWorkflowRole := models.ApprovalWorkflowRole{
			ApprovalWorkflowID: int(approvalWorkflow.Id),
			RoleID:             1,
			Status:             "active",
			CreatedDate:        currentTime,
			CreatedBy:          createdBy,
		}

		if err := connection.DB.Create(&approvalWorkflowRole).Error; err != nil {
			return err
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
		return err
	}

	// Get user role ID
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "status unauthorized",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.Users
	if err := connection.DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		return err
	}

	// Get the Approval to be updated
	var approval models.Approval
	if err := connection.DB.Where("id = ?", approvalID).First(&approval).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Data Not Found",
		})
	}

	// Get the ApprovalWorkflow based on the current process of the Approval
	var currentWorkflow models.ApprovalWorkflow
	if err := connection.DB.Where("approval_setting_id = ? AND \"order\" = ?", approval.ApprovalSettingID, approval.CurrentProcess).First(&currentWorkflow).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "No current approval workflow found",
		})
	}

	// Check if the user has permission to access the current ApprovalWorkflow
	var workflowRole models.ApprovalWorkflowRole
	if err := connection.DB.Where("approval_workflow_id = ? AND role_id = ? AND status = ?", currentWorkflow.Id, user.RoleId, "active").First(&workflowRole).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User does not have permission to access the current ApprovalWorkflow",
		})
	}

	// Get the next ApprovalWorkflow
	var nextWorkflow models.ApprovalWorkflow
	if err := connection.DB.Where("approval_setting_id = ? AND \"order\" > ?", approval.ApprovalSettingID, approval.CurrentProcess).Order("\"order\" asc").First(&nextWorkflow).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "No next approval workflow found",
			})
		}
		return err
	}

	// Update current_process on Approval with the order of the found ApprovalWorkflow
	approval.CurrentProcess = nextWorkflow.Order
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
		return err
	}

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
	approval.CurrentProcess = firstWorkflow.Order
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
		DisplayData    string `json:"display_data"`
		ApprovalStatus string `json:"approval_status"`
		CreatedBy      string `json:"created_by"`
		Description    string `json:"description"`
	}

	// Query menggunakan raw SQL untuk melakukan join antara tabel Approval dan ApprovalWorkflow
	query := `SELECT a.display_data, a.approval_status, a.created_by, aw.description
              FROM approvals a
              JOIN approval_workflows aw ON a.current_process = aw.id`

	if err := connection.DB.Raw(query).Scan(&approvals).Error; err != nil {
		return err
	}

	return c.JSON(approvals)
}
