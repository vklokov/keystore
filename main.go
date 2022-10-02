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

	// user := entities.User{}
	// db.Conn.Where("email = ?", "klokov.dev@gmail4.com").First(&user)
	// log.Printf("%v", user.ID)
	// db.Conn.Model(&user).Updates(map[string]interface{}{"name": "TTTTFFF FFFFF"})

	app.Listen(":3000")
}
