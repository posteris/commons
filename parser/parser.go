package parser

import (
	"github.com/gofiber/fiber/v2"
	"github.com/posteris/commons/errors"
	"github.com/posteris/commons/validation"
)

//BodyParser
func BodyParser(model *interface{}, c *fiber.Ctx) error {
	if err := c.BodyParser(model); err != nil {
		return c.Status(fiber.StatusUnsupportedMediaType).JSON(
			errors.CreateDefaultError(err.Error()),
		)
	}

	if err := validation.ValidateModel(model); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return nil
}
