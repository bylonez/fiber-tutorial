package common

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

// ErrorHandler handle panic
func ErrorHandler(c *fiber.Ctx, err error) error {
	statusCode := fiber.StatusInternalServerError
	resultCode := fiber.StatusInternalServerError
	var value ErrorStruct
	// if custom ErrorStruct, use code
	if errors.As(err, &value) == true {
		statusCode = fiber.StatusBadRequest
		resultCode = int(value.err)
	}
	return c.Status(statusCode).JSON(Result{
		Code: resultCode,
		Msg:  err.Error()},
	)
}
