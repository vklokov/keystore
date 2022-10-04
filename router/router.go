package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vklokov/keystore/controllers"
	mw "github.com/vklokov/keystore/middlewares"
)

func New(app *fiber.App) {
	controller := controllers.Create()

	app.Get("/api/v1/ping", controller.Ping.Ping)

	app.Post("/api/v1/auth", controller.Auth.SignIn)
	app.Delete("/api/v1/auth", controller.Auth.SignOut)

	app.Get("/api/v1/users", mw.WithJWTAuth, controller.Users.Me)
	app.Post("/api/v1/users", controller.Users.Create)

	secrets := app.Group("/api/v1/secrets", mw.WithJWTAuth)
	secrets.Get("/", controller.Secrets.All)
	secrets.Post("/", controller.Secrets.Create)
	secrets.Get("/:id", controller.Secrets.Find)
	secrets.Put("/:id", controller.Secrets.Update)
}
