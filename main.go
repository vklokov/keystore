package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vklokov/keystore/config"
	"github.com/vklokov/keystore/middlewares"
	"github.com/vklokov/keystore/router"
)

func main() {
	config.Boot()

	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})

	middlewares.Apply(app)
	router.New(app)

	app.Listen(":3000")
}
