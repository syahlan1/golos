package roleController

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/roleService"
	"github.com/syahlan1/golos/utils"
)

func ShowRole(c *fiber.Ctx) error {

	result, err := roleService.ShowRole()
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

func ShowAllPermissions(c *fiber.Ctx) error {
	result, err := roleService.ShowAllPermissions()
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

func ShowPermissions(c *fiber.Ctx) error {
	roleID := c.Params("id")
	result, err := roleService.ShowPermissions(roleID)
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

func CreateRole(c *fiber.Ctx) error {
	// Parse request body
	var data models.CreateRole
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
		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	err = roleService.CreateRole(claims, data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Role created successfully",
	})
}

func DeleteRole(c *fiber.Ctx) error {
	roleID := c.Params("id")

	err := roleService.DeleteRole(roleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Role deleted successfully",
	})
}

func UpdateRole(c *fiber.Ctx) error {
	// Parse request body
	var data models.CreateRole
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
		return errors.New("status unauthorized")
	}

	roleID := c.Params("id")

	err = roleService.UpdateRole(claims, roleID, data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Role updated successfully",
	})
}

func CreateRoleModules(c *fiber.Ctx) error {
	var data models.CreateRoleModules
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
		return errors.New("status unauthorized")
	}

	err = roleService.CreateRoleModules(claims, data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success Create Role Modules",
	})
}

func ShowAllRoleMenu(c *fiber.Ctx) error {
	roleId := c.Params("id")
	result, err := roleService.ShowAllRoleMenu(roleId)
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

func CreateRoleMenu(c *fiber.Ctx) error {
	var data []models.RoleMenu
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	roleId := c.Params("id")

	// Get user ID
	claims, err := utils.ExtractJWT(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return errors.New("status unauthorized")
	}

	err = roleService.CreateRoleMenu(claims,roleId, data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success Create Menu",
	})
}

func ShowRoleModules(c *fiber.Ctx) error {
	roleID := c.Params("id")
	result, err := roleService.ShowRoleModules(roleID)
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

func CreateRoleModuleTables(c *fiber.Ctx) error {
	var data []models.CreateRoleModuleTables
	roleID := c.Params("id")
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
		return errors.New("status unauthorized")
	}

	err = roleService.CreateRoleModuleTables(claims, roleID, data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success Create Role Modules & Tables",
	})
}