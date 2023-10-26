package main

import (
	"log"

	"github.com/aalvrz/shelfy/database"
	"github.com/aalvrz/shelfy/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// app.Use(cors.New())

	// Setup Database
	if err := database.ConnectDB(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}

	router.SetupRoutes(app)
	app.Listen(":8080")
}
