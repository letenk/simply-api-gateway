package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/product/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		return c.JSON(fiber.Map{
			"service": "Service Product",
			"message": "Hello from Service Product!",
			"id":      id,
		})
	})

	log.Println("Service product running on port 5001...")
	app.Listen(":5001")
}
