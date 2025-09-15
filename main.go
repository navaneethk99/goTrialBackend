package main

import (
	"crudProject/database"
	"crudProject/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Put("/users/:id", routes.UpdateUser)
	app.Delete("/users/:id", routes.DeleteUser)

	// product endpoints
	app.Post("/product", routes.CreateProduct)
	app.Get("/products", routes.GetProducts)
	app.Get("/products/:id", routes.GetProduct)
	app.Put("/products/:id", routes.UpdateProduct)
	app.Delete("/products/:id", routes.DeleteProduct)

}

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // frontend origin
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	setupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
