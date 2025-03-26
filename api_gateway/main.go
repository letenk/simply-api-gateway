package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func rateLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Expiration: time.Second * 60,
		Max:        5,
	})

}

func reverseProxy(target string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		url := target + c.OriginalURL()
		log.Println("Proxying request to:", url)
		return proxy.Do(c, url)
	}
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	// Implement rate limitter
	app.Use(rateLimiter())

	// Implement proxy
	// Redirect request to service product
	app.Use("/product", reverseProxy("http://localhost:5001"))

	// Redirect request to service order
	app.Use("/order", reverseProxy("http://localhost:5002"))

	// Redirect request to service cart
	app.Use("/cart", reverseProxy("http://localhost:5003"))

	port := ":3000"
	log.Printf("Service API Gateway running on port %s...", port)
	log.Fatal(app.Listen(port))
}
