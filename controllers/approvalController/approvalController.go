package approvalController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/approvalService"
	"github.com/syahlan1/golos/utils"
)

func CreateApprovalSetting(c *fiber.Ctx) error {
	var data models.CreateApprovalSetting

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	username, err := utils.TakeUsername(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	err = approvalService.CreateApprovalSetting(data, username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusCreated,
		Message: "Approval Setting Created!",
	})
}

func UpdateApprovalStatus(c *fiber.Ctx) error {
	approvalID := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	result, err := approvalService.UpdateApprovalStatus(claims, approvalID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Approval updated successfully",
		Data:    result,
	})
}

func RejectApproval(c *fiber.Ctx) error {
	approvalId := c.Params("id")

	var updatedApproval models.Approval
	if err := c.BodyParser(&updatedApproval); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	result, err := approvalService.RejectApproval(approvalId, updatedApproval)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Approval rejected successfully",
		Data:    result,
	})
}

func ShowAllData(c *fiber.Ctx) error {

	result, err := approvalService.ShowAllData()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ApprovalDataDetail(c *fiber.Ctx) error {
	approvalID := c.Params("id")

	result, err := approvalService.ApprovalDataDetail(approvalID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func UpdateApprovalWorkflowRoles(c *fiber.Ctx) error {
	// Parse request body
	var data models.UpdateApprovalWorkflowRoles
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Get approval_workflow_id from URL parameter
	approvalWorkflowID := c.Params("id")

	err := approvalService.UpdateApprovalWorkflowRoles(approvalWorkflowID, data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Approval workflow roles updated successfully",
	})
}
