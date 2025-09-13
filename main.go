package main

import (
	"crudProject/database"
	"crudProject/routes"
	"github.com/gofiber/fiber/v2"
)
import "log"

func welcome(c *fiber.Ctx) error {
	return c.SendString("API is running successfully")
}

func setupRoutes(app *fiber.App) {
	// welcome endpoint
	app.Get("/", welcome)

	// user endpoints
	app.Post("/users", routes.CreateUser)
	app.Get("/users", routes.GetUsers)
	app.Get("/users/:id", routes.GetUser)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
