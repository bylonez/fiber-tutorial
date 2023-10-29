package common

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

// ErrorHandler handle panic
func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var value ErrorStruct
	// if custom ErrorStruct, use code
	if errors.As(err, &value) == true {
		code = int(value.err)
	}
	return c.Status(fiber.StatusInternalServerError).JSON(Result{
		Code: code,
		Msg:  err.Error()},
	)
}
