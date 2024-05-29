package templates

func HandlerCRUDTemplate() string {
	return `package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"gorm.io/gorm"
		"{{.PackagePath}}/models"
		"{{.PackagePath}}/utils"
	)
	
	func Create{{.TableName | ToCamel}}(c *fiber.Ctx) error {
		var data map[string]interface{}
	
		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Response{
				Code:    fiber.StatusBadRequest,
				Message: err.Error(),
			})
		}
	
		var {{.TableName}} models.{{.TableName | ToCamel}}
	
		{{range $column := .Columns}}
		{{$.TableName | ToCamel}}.{{$column.FieldName | ToCamel}} = utils.Get{{mapFieldTypeModel $column.FieldType | ToCamel}}Value(data, "{{$column.FieldName}}")
		{{end}}
	
		// Buat data ke database
		if err := connection.DB.Create(&{{.TableName}}).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Response{
				Code:    fiber.StatusBadRequest,
				Message: err.Error(),
			})
		}
	
		return c.JSON(models.Response{
			Code:    fiber.StatusOK,
			Message: "Success",
		})
	}
	
	func Get{{.TableName | ToCamel}}(c *fiber.Ctx) error {
		db := c.Locals("db").(*gorm.DB)
		var {{.TableName}} []models.{{.TableName | ToCamel}}
	
		if err := db.Find(&{{.TableName}}).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	
		return c.JSON({{.TableName}})
	}
	
	func Get{{.TableName | ToCamel}}ByID(c *fiber.Ctx) error {
		db := c.Locals("db").(*gorm.DB)
		id := c.Params("id")
		var {{.TableName}} models.{{.TableName | ToCamel}}
	
		if err := db.First(&{{.TableName}}, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Data Not Found"})
		}
	
		return c.JSON({{.TableName}})
	}
	
	func Update{{.TableName | ToCamel}}(c *fiber.Ctx) error {
		id := c.Params("id")
	
		var {{.TableName}} models.{{.TableName | ToCamel}}
		if err := connection.DB.First(&{{.TableName}}, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(models.Response{
				Code:    fiber.StatusNotFound,
				Message: "{{.TableName | ToCamel}} not found",
			})
		}
	
		var data map[string]interface{}
		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Response{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid request payload",
			})
		}
	
		{{range $column := .Columns}}
		{{$.TableName | ToCamel}}.{{$column.FieldName | ToCamel}} = utils.Get{{mapFieldTypeModel $column.FieldType | ToCamel}}Value(data, "{{$column.FieldName}}")
		{{end}}
	
		if err := connection.DB.Save(&{{.TableName}}).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
				Code:    fiber.StatusInternalServerError,
				Message: "Failed to update the {{.TableName | ToCamel}} data",
			})
		}
	
		return c.JSON(models.Response{
			Code:    fiber.StatusOK,
			Message: "Success",
			Data:    {{.TableName}},
		})
	}
	
	func Delete{{.TableName | ToCamel}}(c *fiber.Ctx) error {
		db := c.Locals("db").(*gorm.DB)
		id := c.Params("id")
		var {{.TableName}} models.{{.TableName | ToCamel}}
	
		if err := db.First(&{{.TableName}}, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Data Not Found"})
		}
	
		if err := db.Delete(&{{.TableName}}).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	
		return c.SendStatus(fiber.StatusNoContent)
	}
	`
}
