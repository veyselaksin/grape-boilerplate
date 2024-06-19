package baserouter

import (
	"grape-boilerplate/cmd/basecontroller"
	"grape-boilerplate/db/connection"

	"github.com/gofiber/fiber/v2"
)

func InitializeRouters(app fiber.Router, client connection.Client) {

	// API v1 routes
	api := app.Group("/v1")

	// Health check routes
	basecontroller.HealthCheck(api)

}
