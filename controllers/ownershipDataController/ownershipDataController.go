package ownershipDataController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/ownershipDataService"
)

func CreateOwnershipData(c *fiber.Ctx) error {
	var data models.CreateOwnershipData
	generalInformationId := c.Query("general_information_id")

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err := ownershipDataService.CreateOwnershipData(generalInformationId,data)
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
	generalInformationId := c.Query("general_information_id")

	result := ownershipDataService.ShowOwnershipData(generalInformationId)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowOwnershipName(c *fiber.Ctx) error {
	generalInformationId := c.Query("general_information_id")

	result := ownershipDataService.ShowOwnershipName(generalInformationId)

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
	generalInformationId := c.Query("general_information_id")

	var data models.RelationWithBank

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err := ownershipDataService.CreateRelationWithBank(generalInformationId, &data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusCreated,
		Message: "insert sukses",
		Data:    data,
	})
}

func CreateRekeningDebitur(c *fiber.Ctx) error {
	generalInformationId := c.Query("general_information_id")
	var data models.DataRekeningDebitur
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err := ownershipDataService.CreateRekeningDebitur(generalInformationId, &data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "insert sukses",
		Data:    data,
	})
}

func CreateCustomerLoanInfo(c *fiber.Ctx) error {
	generalInformationId := c.Query("general_information_id")
	var data models.CustomerLoanInfo
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err := ownershipDataService.CreateCustomerLoanInfo(generalInformationId, &data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "insert sukses",
		Data:    data,
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

func UpdateCustomerLoanInfo(c *fiber.Ctx) error {
	id := c.Params("id")

	var data models.CustomerLoanInfo
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Invalid Data",
		})
	}

	result, err := ownershipDataService.UpdateCustomerLoanInfo(id, data)
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

func DeleteCustomerLoanInfo(c *fiber.Ctx) error {
	Id := c.Params("id")

	result, err := ownershipDataService.DeleteCustomerLoanInfo(Id)
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
	generalInformationId := c.Query("general_information_id")
	result := ownershipDataService.ShowRekeningDebitur(generalInformationId)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowCustomerLoanInfo(c *fiber.Ctx) error  {
	generalInformationId := c.Query("general_information_id")
	result := ownershipDataService.ShowCustomerLoanInfo(generalInformationId)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowFacilityNo(c *fiber.Ctx) error {
	result := ownershipDataService.ShowFacilityNo()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowProduct(c *fiber.Ctx) error {
	result := ownershipDataService.ShowProduct()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowCustomerAA(c *fiber.Ctx) error {
	generalInformationId := c.Query("general_information_id")
	result := ownershipDataService.ShowCustomerAA(generalInformationId)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}