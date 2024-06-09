package helloRoutes

import (
	helloHandler "github.com/Central-University-IT-prod/backend-oleg-top/api/internal/handlers/hello"
	"github.com/gofiber/fiber/v3"
)

func SetupHelloRoutes(router fiber.Router) {
	hello := router.Group("/hello")
	hello.Get("/", helloHandler.HelloWorld)
}
