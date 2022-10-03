package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vklokov/keystore/config"
	"github.com/vklokov/keystore/controllers"
	"github.com/vklokov/keystore/middlewares"
)

func main() {
	config.Boot()

	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})

	middlewares.Apply(app)
	controllers.Init(app)

	app.Listen(":3000")
}
