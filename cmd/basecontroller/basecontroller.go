package basecontroller

import (
	"github.com/gofiber/fiber/v2"
)

func welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to Viodash Youtube API")
}

func healthcheck(ctx *fiber.Ctx) error {
	return ctx.SendString("keycloak is running")
}

func HealthCheck(app fiber.Router) {

	app.Get("/", welcome)

	app.Get("/health-check", healthcheck)
}
