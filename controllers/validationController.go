package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
)

func ShowAllValidation(c *fiber.Ctx) error {
	db := connection.DB
	var validations []models.Validation
	db.Find(&validations)
	return c.JSON(validations)
}

func CreateValidation(c *fiber.Ctx) error {
	db := connection.DB

	data := new(models.Validation)
	if err := c.BodyParser(data); err != nil {
		return err
	}

	if err := db.Where("name = ?", data.Name).First(&data).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Validation Name already exists"})
	}

	db.Create(&data)
	return c.JSON("create validation success!")
}

func UpdateValidation(c *fiber.Ctx) error {
	db := connection.DB
	id := c.Params("id")
	data := new(models.Validation)
	if err := c.BodyParser(data); err != nil {
		return err
	}
	var validation models.Validation
	if err := db.First(&validation, id).Error; err != nil {
		return err
	}
	db.Model(&validation).Updates(data)
	return c.JSON(validation)
}

func DeleteValidation(c *fiber.Ctx) error {
	db := connection.DB
	id := c.Params("id")
	var validation models.Validation
	if err := db.First(&validation, id).Error; err != nil {
		return err
	}
	db.Delete(&validation)
	return c.JSON("Deleted!")
}
