package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/syahlan1/golos/controllers"
)

func Setup(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5174",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	//business
	app.Post("/api/business/create", controllers.BusinessCreate)
	app.Get("/api/business/show", controllers.BusinessShow)
	app.Put("/api/business/update/:id", controllers.BusinessUpdate)
	app.Delete("/api/business/delete/:id", controllers.BusinessDelete)

	//applicant
	app.Post("/api/applicant/create", controllers.ApplicantCreate)
	app.Get("/api/applicant/show", controllers.ApplicantShow)
	app.Put("/api/applicant/update/:id", controllers.ApplicantUpdate)
	app.Delete("/api/applicant/delete/:id", controllers.ApplicantDelete)
}
