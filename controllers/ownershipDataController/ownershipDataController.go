package ownershipDataController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/ownershipDataService"
)

func CreateOwnershipData(c *fiber.Ctx) error {
	var data models.CreateOwnershipData

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err := ownershipDataService.CreateOwnershipData(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusCreated,
		Message: "insert sukses",
	})
}

func ShowOwnershipData(c *fiber.Ctx) error {
	result := ownershipDataService.ShowOwnershipData()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func EditOwnershipData(c *fiber.Ctx) error {
	id := c.Params("id")

	var data models.OwnershipData
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid Data",
		})
	}

	result, err := ownershipDataService.EditOwnershipData(id, data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func DeleteOwnershipData(c *fiber.Ctx) error {
	ownershipId := c.Params("id")

	result, err := ownershipDataService.DeleteOwnershipData(ownershipId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Deleted",
		Data:    result,
	})

}

func CreateRelationWithBank(c *fiber.Ctx) error {
	ownershipId := c.Params("id")

	ownershipIdInt, err := strconv.Atoi(ownershipId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ownership ID"})
	}

	var data models.CreateRelationWithBank

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err = ownershipDataService.CreateRelationWithBank(ownershipIdInt, data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusCreated,
		Message: "insert sukses",
	})
}

func UpdateRekeningDebitur(c *fiber.Ctx) error {
	id := c.Params("id")

	var data models.DataRekeningDebitur
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid Data",
		})
	}

	result, err := ownershipDataService.UpdateRekeningDebitur(id, data)
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

func UpdateRelationWithBank(c *fiber.Ctx) error {
	id := c.Params("id")

	var data models.RelationWithBank
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid Data",
		})
	}

	result, err := ownershipDataService.UpdateRelationWithBank(id, data)
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

func DeleteRelationWithBank(c *fiber.Ctx) error {
	Id := c.Params("id")

	result, err := ownershipDataService.DeleteRelationWithBank(Id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Deleted!",
		Data:    result,
	})

}

func DeleteRekeningDebitur(c *fiber.Ctx) error {
	Id := c.Params("id")

	result, err := ownershipDataService.DeleteRekeningDebitur(Id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Deleted!",
		Data:    result,
	})

}

func ShowRelationWithBank(c *fiber.Ctx) error {
	result := ownershipDataService.ShowRelationWithBank()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowRekeningDebitur(c *fiber.Ctx) error {
	result := ownershipDataService.ShowRekeningDebitur()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}
