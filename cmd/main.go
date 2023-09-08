package main

import (
	"app/database"
	"app/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Fiber App",
	})

	database.Connect()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
