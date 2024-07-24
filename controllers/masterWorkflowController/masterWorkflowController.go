package masterWorkflowController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/masterWorkflowService"
	"github.com/syahlan1/golos/utils"
)

func CreateMasterWorkflow(c *fiber.Ctx) error {
	var data models.MasterWorkflow

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err = masterWorkflowService.CreateMasterWorkflow(claims, data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Return success response
	return c.JSON(models.Response{
		Code:    fiber.StatusCreated,
		Message: "Master Workflow Created!",
	})
}

func ShowMasterWorkflow(c *fiber.Ctx) error {
	result := masterWorkflowService.ShowMasterWorkflow()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func UpdateMasterWorkflow(c *fiber.Ctx) error {
	masterWorkflowId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var data models.MasterWorkflow
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid Master Workflow Data",
		})
	}

	result, err := masterWorkflowService.UpdateMasterWorkflow(claims, masterWorkflowId, data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Updated!",
		Data:    result,
	})
}

func DeleteMasterWorkflow(c *fiber.Ctx) error {
	masterWorkflowId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err = masterWorkflowService.DeleteMasterWorkflow(claims, masterWorkflowId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Deleted!",
	})
}

func CreateMasterWorkflowStep(c *fiber.Ctx) error {
	var data models.MasterWorkflowStep

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err = masterWorkflowService.CreateMasterWorkflowStep(claims, data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Return success response
	return c.JSON(models.Response{
		Code:    fiber.StatusCreated,
		Message: "Master Workflow Step Created!",
	})
}

func ShowMasterWorkflowStep(c *fiber.Ctx) error {
	result := masterWorkflowService.ShowMasterWorkflowStep("")

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowMasterWorkflowStepById(c *fiber.Ctx) error {
	masterWorkflowStepId := c.Params("id")
	result := masterWorkflowService.ShowMasterWorkflowStep(masterWorkflowStepId)
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowMasterWorkflowStepByGroupId(c *fiber.Ctx) error {
	groupId := c.Params("group_id")
	result := masterWorkflowService.ShowMasterWorkflowStepByGroupId(groupId)
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func UpdateMasterWorkflowStep(c *fiber.Ctx) error {
	masterWorkflowStepId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var data models.MasterWorkflowStep
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid Master Workflow Step Data",
		})
	}

	result, err := masterWorkflowService.UpdateMasterWorkflowStep(claims, masterWorkflowStepId, data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Updated!",
		Data:    result,
	})
}

func DeleteMasterWorkflowStep(c *fiber.Ctx) error {
	masterWorkflowStepId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err = masterWorkflowService.DeleteMasterWorkflowStep(claims, masterWorkflowStepId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Deleted!",
	})
}
