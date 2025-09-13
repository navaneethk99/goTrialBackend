package main

import (
	"crudProject/database"
	"github.com/gofiber/fiber/v2"
)
import "log"

func welcome(c *fiber.Ctx) error {
	return c.SendString("API is running successfully")
}
func main() {
	database.ConnectDb()
	app := fiber.New()
	app.Get("/", welcome)
	log.Fatal(app.Listen(":8000"))
}
