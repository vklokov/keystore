package middlewares

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/vklokov/keystore/config"
	"github.com/vklokov/keystore/utils"
)

func Apply(app *fiber.App) {
	writer := config.CreateLogger(fmt.Sprintf("log/%v.log", os.Getenv("APP_ENV")))

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path} ${queryParams} ${body}\n",
		TimeFormat: utils.DATE_TIME_FORMAT,
		TimeZone:   "Local",
		Output:     writer,
	}))
}
