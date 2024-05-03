package ex

import (
	"errors"
	"fiber-tutorial/pkg/dto"
	"github.com/gofiber/fiber/v2"
)

// ErrorHandler handle panic
func ErrorHandler(c *fiber.Ctx, err error) error {
	statusCode := fiber.StatusInternalServerError
	resultCode := fiber.StatusInternalServerError
	var value ExceptionPanic
	// if custom ExceptionPanic, use code
	if errors.As(err, &value) == true {
		statusCode = fiber.StatusBadRequest
		resultCode = int(value.Ex)
	}
	return c.Status(statusCode).JSON(dto.Result{
		Code: resultCode,
		Msg:  err.Error()},
	)
}
