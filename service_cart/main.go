package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/cart", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"service": "Service Cart", "message": "Hello from service Cart!"})
	})

	log.Println("Service cart running on port 5003...")
	app.Listen(":5003")
}
