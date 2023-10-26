package router

import (
	"github.com/aalvrz/shelfy/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// Users
	users := api.Group("/users")
	users.Post("/", handler.CreateUser)
}
