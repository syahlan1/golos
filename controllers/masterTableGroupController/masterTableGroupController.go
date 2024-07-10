package masterTableGroupController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/masterTableGroupService"
	"github.com/syahlan1/golos/utils"
)

func CreateMasterTableGroup(c *fiber.Ctx) error {
	var data models.MasterTableGroup

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

	err = masterTableGroupService.CreateMasterTableGroup(claims, data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Return success response
	return c.JSON(models.Response{
		Code:    fiber.StatusCreated,
		Message: "Master TableGroup Created!",
	})
}

func ShowMasterTableGroup(c *fiber.Ctx) error {
	result := masterTableGroupService.ShowMasterTableGroup()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowMasterTableGroupDetail(c *fiber.Ctx) error {
	masterTableGroupId := c.Params("id")
	result, err := masterTableGroupService.ShowMasterTableGroupDetail(masterTableGroupId)
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

func UpdateMasterTableGroup(c *fiber.Ctx) error {
	masterTableGroupId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var data models.MasterTableGroup
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid Master TableGroup Data",
		})
	}

	result, err := masterTableGroupService.UpdateMasterTableGroup(claims, masterTableGroupId, data)
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

func DeleteMasterTableGroup(c *fiber.Ctx) error {
	masterTableId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err = masterTableGroupService.DeleteMasterTableGroup(claims, masterTableId)
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

func CreateMasterTableItem(c *fiber.Ctx) error {
	var data models.MasterTableItem

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

	err = masterTableGroupService.CreateMasterTableItem(claims, data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Return success response
	return c.JSON(models.Response{
		Code:    fiber.StatusCreated,
		Message: "Master TableItem Created!",
	})
}

func ShowMasterTableItem(c *fiber.Ctx) error {
	groupId := c.Query("group_id")
	result := masterTableGroupService.ShowMasterTableItem(groupId)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowMasterTableItemDetail(c *fiber.Ctx) error {
	masterTableItemId := c.Params("id")
	result, err := masterTableGroupService.ShowMasterTableItemDetail(masterTableItemId)
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

func UpdateMasterTableItem(c *fiber.Ctx) error {
	masterTableItemId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var data models.MasterTableItem
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid Master TableItem Data",
		})
	}

	result, err := masterTableGroupService.UpdateMasterTableItem(claims, masterTableItemId, data)
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

func DeleteMasterTableItem(c *fiber.Ctx) error {
	masterTableId := c.Params("id")

	// Get user role ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err = masterTableGroupService.DeleteMasterTableItem(claims, masterTableId)
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

func ShowFormMasterTableGroup(c *fiber.Ctx) error {
	// Ambil ID tabel dari parameter rute
	groupName := c.Params("group_name")

	claims, err := utils.TakeUsername(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}


	result ,err := masterTableGroupService.ShowFormMasterTableGroup(groupName, claims)
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

func ShowDataMasterTableGroup(c *fiber.Ctx) error {
	tableGroup := c.Params("table_group")
	tableItem := c.Params("table_item")

	claims, err := utils.TakeUsername(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	result, err := masterTableGroupService.ShowDataMasterTableGroup(tableGroup, tableItem, claims, "")
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

func ShowApprovalTableGroupItem(c *fiber.Ctx) error {
	groupName := c.Params("group_name")

	claims, err := utils.TakeUsername(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}


	result ,err := masterTableGroupService.ShowApprovalTableGroupItem(groupName, claims)
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

func ShowDataMasterTableGroupById(c *fiber.Ctx) error {
	tableGroup := c.Params("table_group")
	tableItem := c.Params("table_item")
	id := c.Params("id")

	result, err := masterTableGroupService.ShowDataMasterTableGroup(tableGroup, tableItem, "", id)
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

func CreateDataMasterTableGroup(c *fiber.Ctx) error {
	tableGroup := c.Params("table_group")
	tableItem := c.Params("table_item")

	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	claims, err := utils.TakeUsername(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err, errValidation := masterTableGroupService.CreateDataMasterTableGroup(tableGroup, tableItem, claims, data)
	if errValidation != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Error Validation",
			Data:    errValidation,
		})
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response{
		Code:    fiber.StatusCreated,
		Message: "Success Create",
	})
}

func UpdateDataMasterTableGroup(c *fiber.Ctx) error {
	tableGroup := c.Params("table_group")
	tableItem := c.Params("table_item")
	id := c.Params("id")

	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	claims, err := utils.TakeUsername(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err, errValidation := masterTableGroupService.UpdateDataMasterTableGroup(tableGroup, tableItem, id, claims, data)
	if errValidation != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Error Validation",
			Data:    errValidation,
		})
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success Update",
	})
}

func DeleteDataMasterTableGroup(c *fiber.Ctx) error {
	tableGroup := c.Params("table_group")
	tableItem := c.Params("table_item")
	id := c.Params("id")

	claims, err := utils.TakeUsername(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err = masterTableGroupService.DeleteDataMasterTableGroup(tableGroup, tableItem, id, claims)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success Delete",
	})
}

func SubmitTableGroupItem(c *fiber.Ctx) error {
	var data models.TableGroupItemStatus

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	claims, err := utils.TakeUsername(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	err= masterTableGroupService.SubmitTableGroupItem(claims, data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success Update",
	})

}

func GenerateTableGroup(c *fiber.Ctx) error {
	groupId := c.Params("id")

	err := masterTableGroupService.GenerateTableGroupByGroupId(groupId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Table generated successfully!",
	})
}