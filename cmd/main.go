package main

import (
	"grape-boilerplate/cmd/baserouter"
	"grape-boilerplate/cmd/config"
	"grape-boilerplate/db/connection"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title Viodash Youtube API
// @version 1.0
// @description This is a server for Viodash Youtube API.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @schemes http https
// @BasePath /v1
func main() {
	client := connection.New()

	app := fiber.New(config.FiberConfig)

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: time.DateTime,
	}))

	// Initialize routes
	baserouter.InitializeRouters(app, client)

	// Start listening on port 8000
	go func() {
		if err := app.Listen(":8000"); err != nil {
			panic(err)
		}
	}()

	// Graceful shutdown
	GracefulShutdown(app, client)
}

func GracefulShutdown(app *fiber.App, client connection.Client) {

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Received terminate,graceful shutdown", sig)

	database, err := client.PostgresConnection.DB()
	if err != nil {
		log.Fatalln("PostgreSQL Closing ERROR :", err)
	}
	database.Close()
	log.Println("PostgreSQL Closed")

	app.Shutdown()
}
