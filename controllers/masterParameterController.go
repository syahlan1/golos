package controllers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func CreateParameter(c *fiber.Ctx) error {
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

	masterParameter := models.MasterParameter{
		CreatedBy:          createdBy,
		CreatedDate:        timeNow,
		Status:             "L",
		Description:        data["description"].(string),
		EnglishDescription: data["english_description"].(string),
		ParamKey:           data["param_key"].(string),
		ParamValue:         data["param_value"].(string),
	}

	if err := connection.DB.Create(&masterParameter).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create Parameter"})
	}

	// Return success response
	return c.JSON(fiber.Map{"message": "Master Parameter Created!"})
}

func ShowAllParameter(c *fiber.Ctx) error {
	var masterParameters []models.MasterParameter

	if err := connection.DB.Where("status = ?", "L").Find(&masterParameters).Error; err != nil {
		return err
	}

	return c.JSON(masterParameters)
}

func ShowParameterDetail(c *fiber.Ctx) error {
	parameterId := c.Params("id")
	var masterParameter models.MasterParameter

	if err := connection.DB.Where("id = ?", parameterId).First(&masterParameter).Error; err != nil {
		return err
	}

	return c.JSON(masterParameter)
}

func UpdateMasterParameter(c *fiber.Ctx) error {
	parameterId := c.Params("id")

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

	var masterParameter models.MasterParameter
	if err := connection.DB.First(&masterParameter, parameterId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Data Not Found",
		})
	}

	var updatedMasterParameter models.MasterParameter
	if err := c.BodyParser(&updatedMasterParameter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid Data Parameter",
		})
	}

	masterParameter.UpdatedBy = updatedBy
	masterParameter.UpdatedDate = time.Now()
	masterParameter.Description = updatedMasterParameter.Description
	masterParameter.EnglishDescription = updatedMasterParameter.EnglishDescription
	masterParameter.ParamKey = updatedMasterParameter.ParamKey
	masterParameter.ParamValue = updatedMasterParameter.ParamValue

	if err := connection.DB.Save(&masterParameter).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to update Master Parameter",
		})
	}

	return c.JSON(fiber.Map{
		"status": "Updated!",
		"data":   masterParameter,
	})
}

func DeleteMasterParameter(c *fiber.Ctx) error {
	parameterId := c.Params("id")

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

	var masterParameter models.MasterParameter
	if err := connection.DB.First(&masterParameter, parameterId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Data Not Found",
		})
	}

	var updatedMasterParameter models.MasterParameter
	if err := c.BodyParser(&updatedMasterParameter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid Data Parameter",
		})
	}

	masterParameter.UpdatedBy = updatedBy
	masterParameter.UpdatedDate = time.Now()
	masterParameter.Status = "D"

	if err := connection.DB.Save(&masterParameter).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Failed to Delete Master Parameter",
		})
	}

	return c.JSON(fiber.Map{
		"status": "Deleted!",
		"data":   masterParameter,
	})
}
