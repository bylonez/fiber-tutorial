package app

import (
	_ "fiber-tutorial/internal/database"
	"fiber-tutorial/internal/handler"
	"fiber-tutorial/internal/pkg"
	_ "fiber-tutorial/internal/service/init"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
	"runtime/debug"
)

func Run() {
	app := fiber.New(fiber.Config{
		ErrorHandler: pkg.ErrorHandler,
	})

	// middleware
	// request id
	app.Use(requestid.New())
	// logger
	logger, _ := zap.NewDevelopment()
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
			_, ok := e.(pkg.ErrorPanic)
			if !ok {
				log.Errorf("[PANIC RECOVER] %s\n%s", e, debug.Stack())
			}
		},
	}))

	handler.Route(app)

	log.Fatal(app.Listen(":3000"))
}
