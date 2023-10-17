package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tedbearr/go-learn/config"
	"github.com/tedbearr/go-learn/helper"
	"github.com/tedbearr/go-learn/route"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Use(logger.New())

	helper.Log()

	config.LoadEnv()

	config.DatabaseConfig()

	port := os.Getenv("PORT")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("Welcome")
	})

	prefixRoute := app.Group("/api/v1")
	route.GlobalParameterRoute(prefixRoute)

	app.Use("*", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("what are you looking for ?!")
	})

	log.Fatal(app.Listen(":" + port))
}
