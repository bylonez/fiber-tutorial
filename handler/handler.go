package handler

import (
	"fiber-tutorial/handler/user"
	"github.com/gofiber/fiber/v2"
)

func SetupHandler(app *fiber.App) {
	user.Handler(app.Group("/user"))

}
