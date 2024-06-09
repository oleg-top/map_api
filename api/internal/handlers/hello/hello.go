package helloHandler

import "github.com/gofiber/fiber/v3"

func HelloWorld(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "is working!",
		"message": "Hello, world!",
	})
}
