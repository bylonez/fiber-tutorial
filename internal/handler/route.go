package handler

import (
	"github.com/bylonez/fiber-tutorial/internal/handler/user"
	"github.com/gofiber/fiber/v3"
)

func Route(app *fiber.App) {
	user.Handler(app.Group("/user"))

}
