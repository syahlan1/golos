package businessController

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/businessService"
	"github.com/syahlan1/golos/utils"
)

func ShowBusinessApplicant(c *fiber.Ctx) error {

	result := businessService.ShowBusinessApplicant()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func BusinessShow(c *fiber.Ctx) error {
	result, err := businessService.BusinessShow()
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

func BusinessShowDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := businessService.BusinessShowDetail(id)
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

func BusinessCreate(c *fiber.Ctx) error {
	var data models.CreateBusiness

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	username, err := utils.TakeUsername(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err = businessService.BusinessCreate(username, data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
	})
}

func BusinessUpdate(c *fiber.Ctx) error {
	businessID := c.Params("id")

	var updatedBusiness models.CreateBusiness
	if err := c.BodyParser(&updatedBusiness); err != nil {
		log.Println("Error parsing request payload:", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request payload",
		})
	}

	result, err := businessService.BusinessUpdate(businessID, updatedBusiness)
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

func BusinessDelete(c *fiber.Ctx) error {
	businessID := c.Params("id")

	err := businessService.BusinessDelete(businessID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Deleted!",
	})
}

func BusinessUploadFile(c *fiber.Ctx) error {

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	result, err := businessService.BusinessUploadFile(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success Upload!",
		Data:    result,
	})
}

func BusinessShowFile(c *fiber.Ctx) error {

	id := c.Params("id")
	result, err := businessService.BusinessShowFile(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	c.Set("Content-Disposition", "attachment; filename=\""+result.DocumentFile+"\"")
	return c.SendFile(result.DocumentPath)
}

func BusinessApproveUpdate(c *fiber.Ctx) error {
	businessID := c.Params("id")

	result, err := businessService.BusinessApproveUpdate(businessID)
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

func ShowCompanyFirstName(c *fiber.Ctx) error {
	result := businessService.ShowCompanyFirstName()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowCompanyType(c *fiber.Ctx) error {
	result := businessService.ShowCompanyType()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowBusinessAddressType(c *fiber.Ctx) error {
	result, err := businessService.ShowBusinessAddressType()
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

func ShowExternalRatingCompany(c *fiber.Ctx) error {
	result:= businessService.ShowExternalRatingCompany()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowRatingClass(c *fiber.Ctx) error {
	externalRatingId := c.Query("external_rating_id")
	result:= businessService.ShowRatingClass(externalRatingId)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowKodeBursa(c *fiber.Ctx) error {
	result:= businessService.ShowKodeBursa()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowBusinessType(c *fiber.Ctx) error {
	result := businessService.ShowBusinessType()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

// zipcode
func GetProvinces(c *fiber.Ctx) error {
	result, err := businessService.GetProvinces()
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

func GetCitiesByProvince(c *fiber.Ctx) error {
	province := c.Query("province")
	result, err := businessService.GetCitiesByProvince(province)
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

func GetDistrictByCity(c *fiber.Ctx) error {
	city := c.Query("city")
	result, err := businessService.GetDistrictByCity(city)
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

func GetSubdistrictByDistrict(c *fiber.Ctx) error {
	district := c.Query("district")
	result, err := businessService.GetSubdistrictByDistrict(district)
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

func GetZipCodesBySubdistrict(c *fiber.Ctx) error {
	subdistrict := c.Query("subdistrict")
	result, err := businessService.GetZipCodesBySubdistrict(subdistrict)
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
