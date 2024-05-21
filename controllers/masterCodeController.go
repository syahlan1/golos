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

func ShowDetailMasterCode(c *fiber.Ctx) error {
	var masterCode []models.MasterCode
	status := "L"

	groupId := c.Params("code_group_id")
	groupName := c.Params("code_group")

	// Periksa apakah groupId tidak kosong
	if groupId != "" {
		// Konversi groupId ke tipe data integer
		groupIdInt, err := strconv.Atoi(groupId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Response{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid group ID",
			})
		}

		// Gunakan groupIdInt dalam kondisi Where
		connection.DB.Where("status = ? AND code_group_id = ?", status, groupIdInt).Find(&masterCode)
	} else if groupName != "" {
		// Jika groupId kosong, gunakan groupName
		connection.DB.Where("status = ? AND code_group = ?", status, groupName).Find(&masterCode)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Missing parameter",
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    masterCode,
	})
}

func ShowMasterCode(c *fiber.Ctx) error {
	var masterCode []models.MasterCode

	connection.DB.Where("status = ?", "L").Find(&masterCode)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    masterCode,
	})
}

func ShowMasterCodeGroup(c *fiber.Ctx) error {
	var masterCodeGroup []models.MasterCodeGroup

	connection.DB.Where("status = ?", "L").Find(&masterCodeGroup)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    masterCodeGroup,
	})
}

func CreateMasterCode(c *fiber.Ctx) error {
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
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Code:    fiber.StatusNotFound,
			Message: "User not found",
		})
	}

	newMasterCode := models.MasterCode{
		Authoriser:         user.Username,
		AuthorizeDate:      timeNow,
		CreatedBy:          user.Username,
		CreatedDate:        timeNow,
		Code:               data["code"].(string),
		CodeGroupId:        int(data["code_group_id"].(float64)),
		Status:             "L",
		Description:        data["description"].(string),
		EnglishDescription: data["english_description"].(string),
		Sequence:           int(data["sequence"].(float64)),
		CodeGroup:          data["code_group"].(string),
	}

	if err := connection.DB.Create(&newMasterCode).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to create Master Code",
		})
	}

	// Return success response
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Master Code Created!",
	})
}

func CreateMasterCodeGroup(c *fiber.Ctx) error {
	var data map[string]string
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
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	//check existing name code group
	var existingCodeGroup models.MasterCodeGroup
	if err := connection.DB.Where("code_group = ?", data["code_group"]).First(&existingCodeGroup).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(models.Response{
			Code:    fiber.StatusConflict,
			Message: "Code Group Name Already Exists",
		})
	}

	newMasterCodeGroup := models.MasterCodeGroup{
		Authoriser:         user.Username,
		AuthorizeDate:      timeNow,
		CreatedBy:          user.Username,
		CreatedDate:        timeNow,
		Status:             "L",
		CodeGroup:          data["code_group"],
		Description:        data["description"],
		EnglishDescription: data["english_description"],
	}

	if err := connection.DB.Create(&newMasterCodeGroup).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to Master Code Group",
		})
	}

	// Return success response
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Master Code Group Created!",
	})
}

func UpdateMasterCode(c *fiber.Ctx) error {
	masterCodeId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	//

	var masterCode models.MasterCode
	if err := connection.DB.First(&masterCode, masterCodeId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Code:    fiber.StatusNotFound,
			Message: "Data Not Found",
		})
	}

	var updatedMasterCode models.MasterCode
	if err := c.BodyParser(&updatedMasterCode); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid Master Code Data",
		})
	}

	masterCode.UpdatedBy = user.Username
	masterCode.UpdatedDate = time.Now()
	masterCode.Code = updatedMasterCode.Code
	masterCode.CodeGroupId = updatedMasterCode.CodeGroupId
	masterCode.Description = updatedMasterCode.Description
	masterCode.EnglishDescription = updatedMasterCode.EnglishDescription
	masterCode.Sequence = updatedMasterCode.Sequence
	masterCode.CodeGroup = updatedMasterCode.CodeGroup

	if err := connection.DB.Save(&masterCode).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to update Master Code",
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Updated!",
		Data:    masterCode,
	})
}

func UpdateMasterCodeGroup(c *fiber.Ctx) error {
	masterCodeGroupId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	//

	var masterCodeGroup models.MasterCodeGroup
	if err := connection.DB.First(&masterCodeGroup, masterCodeGroupId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Code:    fiber.StatusNotFound,
			Message: "Data Not Found",
		})
	}

	var updatedMasterCodeGroup models.MasterCodeGroup
	if err := c.BodyParser(&updatedMasterCodeGroup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid Master Code Data",
		})
	}

	masterCodeGroup.UpdatedBy = user.Username
	masterCodeGroup.UpdatedDate = time.Now()
	masterCodeGroup.CodeGroup = updatedMasterCodeGroup.CodeGroup
	masterCodeGroup.Description = updatedMasterCodeGroup.Description
	masterCodeGroup.EnglishDescription = updatedMasterCodeGroup.EnglishDescription

	if err := connection.DB.Save(&masterCodeGroup).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to update Master Code Group",
		})
	}

	return c.JSON(models.Response{
		Message: "Updated!",
		Data:    masterCodeGroup,
	})
}

func DeleteMasterCode(c *fiber.Ctx) error {
	masterCodeId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	//

	var masterCode models.MasterCode
	if err := connection.DB.First(&masterCode, masterCodeId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Code:    fiber.StatusNotFound,
			Message: "Data Not Found",
		})
	}

	masterCode.UpdatedBy = user.Username
	masterCode.UpdatedDate = time.Now()
	masterCode.Status = "D"

	if err := connection.DB.Save(&masterCode).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to delete Master Code",
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Deleted!",
		Data:   masterCode,
	})
}

func DeleteMasterCodeGroup(c *fiber.Ctx) error {
	masterCodeGroupId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var user models.Users
	if err := connection.DB.Where("id = ?", claims).First(&user).Error; err != nil {
		log.Println("Error retrieving user:", err)
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Message: err.Error(),
			Code:    fiber.StatusNotFound,
		})
	}

	//

	var masterCodeGroup models.MasterCodeGroup
	if err := connection.DB.First(&masterCodeGroup, masterCodeGroupId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Code:    fiber.StatusNotFound,
			Message: "Data Not Found",
		})
	}

	masterCodeGroup.UpdatedBy = user.Username
	masterCodeGroup.UpdatedDate = time.Now()
	masterCodeGroup.Status = "D"

	if err := connection.DB.Save(&masterCodeGroup).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to delete Master Code Group",
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Deleted!",
		Data:   masterCodeGroup,
	})
}
