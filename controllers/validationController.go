package controllers

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/utils"
)

func ShowAllValidations(c *fiber.Ctx) error {
	var validations []models.MasterValidation
	if err := connection.DB.Where("status = ? AND is_active = ?", "L", 1).Find(&validations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to fetch validations",
		})
	}
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    validations,
	})
}

func ShowDetailValidation(c *fiber.Ctx) error {
	var validations models.MasterValidation
	validationId := c.Params("id")

	if err := connection.DB.Where("id = ?", validationId).Find(&validations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to fetch validations",
		})
	}
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    validations,
	})
}

func ShowValidationByColumn(c *fiber.Ctx) error {
	var validations []models.MasterValidation
	columnId := c.Params("id")

	if err := connection.DB.Where("status = ? AND is_active = ? AND column_id = ?", "L", 1, columnId).Find(&validations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to fetch validations"})
	}
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    validations,
	})
}

func CreateValidation(c *fiber.Ctx) error {
	columnIdStr := c.Params("id")
	columnId, err := strconv.Atoi(columnIdStr)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var data map[string]interface{}

	timeNow := time.Now()

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	newMasterValidation := models.MasterValidation{
		CreatedBy:          user.Username,
		CreatedDate:        timeNow,
		Status:             "L",
		Description:        data["description"].(string),
		EnglishDescription: data["english_description"].(string),
		MessageType:        data["message_type"].(string),
		ValidationFunction: data["validation_function"].(string),
		IsActive:           int(data["is_active"].(float64)),
		MasterCodeId:       int(data["master_code_id"].(float64)),
		ColumnId:           columnId,
	}

	if err := connection.DB.Create(&newMasterValidation).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to create Master Validation",
		})
	}

	// Return success response
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Master Validation Created!",
	})
}

func UpdateValidation(c *fiber.Ctx) error {
	masterValidationId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	//

	var masterValidation models.MasterValidation
	if err := connection.DB.First(&masterValidation, masterValidationId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Code:    fiber.StatusNotFound,
			Message: "Data Not Found",
		})
	}

	var updatedMasterValidation models.MasterValidation
	if err := c.BodyParser(&updatedMasterValidation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid Master Validation Data",
		})
	}

	masterValidation.UpdatedBy = user.Username
	masterValidation.UpdatedDate = time.Now()
	masterValidation.Description = updatedMasterValidation.Description
	masterValidation.EnglishDescription = updatedMasterValidation.EnglishDescription
	masterValidation.MessageType = updatedMasterValidation.MessageType
	masterValidation.ValidationFunction = updatedMasterValidation.ValidationFunction
	masterValidation.IsActive = updatedMasterValidation.IsActive

	if err := connection.DB.Save(&masterValidation).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to update Master Validation",
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Updated!",
		Data:   masterValidation,
	})
}

func DeleteValidation(c *fiber.Ctx) error {
	masterValidateId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	//

	var masterValidate models.MasterCode
	if err := connection.DB.First(&masterValidate, masterValidateId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Code:    fiber.StatusNotFound,
			Message: "Data Not Found",
		})
	}

	masterValidate.UpdatedBy = user.Username
	masterValidate.UpdatedDate = time.Now()
	masterValidate.Status = "D"

	if err := connection.DB.Save(&masterValidate).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to delete Master Validation",
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Deleted!",
		Data:   masterValidate,
	})
}
