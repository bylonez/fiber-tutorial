package handler

import (
	"fiber-tutorial/handler/user"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	user.Handler(app.Group("/userservice"))

}
