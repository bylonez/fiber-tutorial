package app

import (
	"github.com/bylonez/fiber-tutorial/internal/database"
	"github.com/bylonez/fiber-tutorial/internal/handler"
	_ "github.com/bylonez/fiber-tutorial/internal/service/init"
	"github.com/bylonez/fiber-tutorial/pkg/ex"
	"github.com/bylonez/fiber-tutorial/pkg/fiberzap"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"go.uber.org/zap"
	"runtime/debug"
)

func Run() {
	app := fiber.New(fiber.Config{
		ErrorHandler: ex.ErrorHandler,
	})
	database.Init()

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
		StackTraceHandler: func(c fiber.Ctx, e any) {
			_, ok := e.(ex.ExceptionPanic)
			if !ok {
				log.Errorf("[PANIC RECOVER] %s\n%s", e, debug.Stack())
			}
		},
	}))
	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	handler.Route(app)

	log.Fatal(app.Listen(":3000"))
}
