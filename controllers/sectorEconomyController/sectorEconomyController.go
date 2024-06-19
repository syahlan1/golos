package sectorEconomyController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/golos/models"
	"github.com/syahlan1/golos/services/sectorEconomyService"
)

func ShowSektorEkonomi1(c *fiber.Ctx) error {
	result := sectorEconomyService.ShowSektorEkonomi1()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowSektorEkonomi2(c *fiber.Ctx) error {
	id := c.Query("id")
	result := sectorEconomyService.ShowSektorEkonomi2(id)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowSektorEkonomi3(c *fiber.Ctx) error {
	id := c.Query("id")
	result := sectorEconomyService.ShowSektorEkonomi3(id)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowSektorEkonomiOjk(c *fiber.Ctx) error {
	id := c.Query("id")
	result := sectorEconomyService.ShowSektorEkonomiOjk(id)

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowLokasiPabrik(c *fiber.Ctx) error {
	result := sectorEconomyService.ShowLokasiPabrik()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowLokasiDati2(c *fiber.Ctx) error {
	result := sectorEconomyService.ShowLokasiDati2()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowHubunganNasabahBank(c *fiber.Ctx) error {
	result := sectorEconomyService.ShowHubunganNasabahBank()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func ShowHubunganKeluarga(c *fiber.Ctx) error {
	result := sectorEconomyService.ShowHubunganKeluarga()

	return c.JSON(models.Response{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}
