package main

import (
	"fiber-tutorial/common"
	"fiber-tutorial/handler"
	"fiber-tutorial/model"
	_ "fiber-tutorial/service/servicedi"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
	"runtime/debug"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: common.ErrorHandler,
	})

	// middleware
	// request id
	app.Use(requestid.New())
	// logger
	logger, _ := zap.NewProduction()
	log.SetLogger(fiberzap.NewLogger(fiberzap.LoggerConfig{
		SetLogger: logger,
	}))
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))
	// recover from panic and log
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			log.Errorf("[PANIC RECOVER] %s\n%s", e, debug.Stack())
		},
	}))

	handler.SetupHandler(app)
	model.InitDB()

	log.Fatal(app.Listen(":3000"))
}
