package templates

func ControllerTemplate() string {
	return `package {{ .TableName | ToLowerCamel}}Controller

import (
	"github.com/gofiber/fiber/v2"
	"{{.PackagePath}}/models"
	"{{.PackagePath}}/services/{{ .TableName | ToLowerCamel }}Service"
	"{{.PackagePath}}/utils"
)

func Show{{ .TableName | ToCamel }}(c *fiber.Ctx) error {
	result, err := {{ .TableName | ToLowerCamel }}Service.Show{{ .TableName | ToCamel }}()
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

func ShowDetail{{ .TableName | ToCamel }}(c *fiber.Ctx) error {
	Id := c.Params("id")

	result := {{ .TableName | ToLowerCamel }}Service.ShowDetail{{ .TableName | ToCamel }}(Id)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func Create{{ .TableName | ToCamel}}(c *fiber.Ctx) error {
	var data models.{{ .TableName | ToCamel}}

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

	result, err := {{ .TableName | ToLowerCamel}}Service.Create{{ .TableName | ToCamel}}(claims, data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Return success response
	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Created!",
		Data:    result,
	})
}

func Update{{ .TableName | ToCamel}}(c *fiber.Ctx) error {
	Id := c.Params("id")

	// Get user role ID
	claims, err := utils.TakeUsername(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var data models.{{ .TableName | ToCamel}}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid Master Code Data",
		})
	}

	result, err := {{ .TableName | ToLowerCamel}}Service.Update{{ .TableName | ToCamel}}(Id, claims, data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Message: "Updated!",
		Data:    result,
	})
}

func Delete{{ .TableName | ToCamel}}(c *fiber.Ctx) error {
	Id := c.Params("id")

	// Get user role ID
	claims, err := utils.TakeUsername(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}

	var data models.{{ .TableName | ToCamel}}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid Master Code Data",
		})
	}

	err = {{ .TableName | ToLowerCamel}}Service.Delete{{ .TableName | ToCamel}}(claims, Id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Message: "Deleted!",
	})
}
`
}
