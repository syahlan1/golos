package main

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/routes"
)

func main() {
	connection.Connect()

	app := fiber.New()

	routes.Setup(app)

	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// }))

	app.Listen(":8000")
}
