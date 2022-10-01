package core

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/vklokov/keystore/db"
	"github.com/vklokov/keystore/handlers"
	"github.com/vklokov/keystore/router"
	"github.com/vklokov/keystore/utils"
)

type App struct {
	Fiber *fiber.App
}

func (app *App) Start(port int) {
	app.Fiber.Listen(fmt.Sprintf(":%v", port))
}

func New() *App {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("* Error loading .env file, %v", err)
	}

	logWriter := utils.LogWriter{
		Location: fmt.Sprintf("log/%v.log", os.Getenv("APP_ENV")),
	}

	handler := handlers.New()
	handler.Logger = &logWriter

	database := db.Conn()
	database.AutoMigrate(
		&db.User{},
	)

	app := App{
		Fiber: fiber.New(),
	}
	app.Fiber.Use(recover.New())
	app.Fiber.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path} ${queryParams} ${body}\n",
		TimeFormat: utils.DATE_TIME_FORMAT,
		TimeZone:   "Local",
		Output:     &logWriter,
	}))

	router.Register(app.Fiber, handler)
	return &app
}
