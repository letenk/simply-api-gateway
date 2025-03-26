package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/order", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"service": "Service Order", "message": "Hello from Service Order!"})
	})

	log.Println("Service order running on port 5002...")
	app.Listen(":5002")
}
