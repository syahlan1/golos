package controllers

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func ShowAllValidations(c *fiber.Ctx) error {
	var validations []models.MasterValidation
	if err := connection.DB.Where("status = ? AND is_active = ?", "L", 1).Find(&validations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch validations"})
	}
	return c.JSON(validations)
}

func ShowDetailValidation(c *fiber.Ctx) error {
	var validations models.MasterValidation
	validationId := c.Params("id")

	if err := connection.DB.Where("id = ?", validationId).Find(&validations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch validations"})
	}
	return c.JSON(validations)
}

func ShowValidationByColumn(c *fiber.Ctx) error {
	var validations []models.MasterValidation
	columnId := c.Params("id")

	if err := connection.DB.Where("status = ? AND is_active = ? AND column_id = ?", "L", 1, columnId).Find(&validations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch validations"})
	}
	return c.JSON(validations)
}

func CreateValidation(c *fiber.Ctx) error {
	columnIdStr := c.Params("id")
	columnId, err := strconv.Atoi(columnIdStr)

	if err != nil {
		return err
	}

	var data map[string]interface{}

	timeNow := time.Now()

	if err := c.BodyParser(&data); err != nil {
		return err
	}

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

	newMasterValidation := models.MasterValidation{
		CreatedBy:          createdBy,
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create Master Validation"})
	}

	// Return success response
	return c.JSON(fiber.Map{"message": "Master Validation Created!"})
}

func UpdateValidation(c *fiber.Ctx) error {
	masterValidationId := c.Params("id")

	updatedBy, err := TakeUsername(c)
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

	//

	var masterValidation models.MasterValidation
	if err := connection.DB.First(&masterValidation, masterValidationId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Data Not Found",
		})
	}

	var updatedMasterValidation models.MasterValidation
	if err := c.BodyParser(&updatedMasterValidation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid Master Validation Data",
		})
	}

	masterValidation.UpdatedBy = updatedBy
	masterValidation.UpdatedDate = time.Now()
	masterValidation.Description = updatedMasterValidation.Description
	masterValidation.EnglishDescription = updatedMasterValidation.EnglishDescription
	masterValidation.MessageType = updatedMasterValidation.MessageType
	masterValidation.ValidationFunction = updatedMasterValidation.ValidationFunction
	masterValidation.IsActive = updatedMasterValidation.IsActive

	if err := connection.DB.Save(&masterValidation).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to update Master Validation",
		})
	}

	return c.JSON(fiber.Map{
		"status": "Updated!",
		"data":   masterValidation,
	})
}

func DeleteValidation(c *fiber.Ctx) error {
	masterValidateId := c.Params("id")

	deletedBy, err := TakeUsername(c)
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

	//

	var masterValidate models.MasterCode
	if err := connection.DB.First(&masterValidate, masterValidateId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Data Not Found",
		})
	}

	masterValidate.UpdatedBy = deletedBy
	masterValidate.UpdatedDate = time.Now()
	masterValidate.Status = "D"

	if err := connection.DB.Save(&masterValidate).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to delete Master Validation",
		})
	}

	return c.JSON(fiber.Map{
		"status": "Deleted!",
		"data":   masterValidate,
	})
}
