package helpers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func DefaultErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	resp := fiber.Map{
		"message": err.Error(),
		"code":    code,
	}

	return c.Status(code).JSON(resp)
}
