package main

import (
	"fiber-tutorial/common"
	"fiber-tutorial/handler"
	"fiber-tutorial/model"
	_ "fiber-tutorial/service/servicedi"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
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
	app.Use(logger.New(logger.Config{
		Format:     "${time} ${status} - ${method} ${path} ${locals:requestid} ${latency}\n",
		TimeFormat: "2006-01-02T15:04:05Z07:00",
		TimeZone:   "Asia/Shanghai",
	}))
	// recover from panic and log
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			log.Errorf("[PANIC RECOVER] %s\n%s", e, debug.Stack())
		},
	}))
	// todo logger format

	handler.SetupHandler(app)
	model.InitDB()

	log.Fatal(app.Listen(":3000"))
}
