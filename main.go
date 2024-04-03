package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/routes"
	// "github.com/syahlan1/golos/seeder"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	connection.Connect()

	app := fiber.New()

	routes.Setup(app)
	// seeder.UserSeed()
	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// }))

	app.Listen(":8000")
}
