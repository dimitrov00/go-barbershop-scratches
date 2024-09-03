package web

import (
	messages "barbershop/creativo/i18n"
	"barbershop/creativo/pkg/errors"
	"barbershop/creativo/pkg/types"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func RecoveryMiddleware(c *fiber.Ctx) error {
	localizer := c.Locals("localizer").(*i18n.Localizer)

	defer func() {
		if r := recover(); r != nil {
			err := errors.ErrInternalServerError
			c.Status(err.Status).JSON(fiber.Map{
				"status":  "error",
				"code":    err.Code,
				"message": localizer.MustLocalize(err.LocalizeConfig),
			})
		}
	}()
	return c.Next()
}

func ErrorHandlingMiddleware(c *fiber.Ctx) error {
	localizer := c.Locals("localizer").(*i18n.Localizer)

	err := c.Next()
	if err == nil {
		return nil
	}

	log.Printf("Error: %+v", err)

	if e, ok := err.(*types.ApiError); ok {
		return c.Status(e.Status).JSON(fiber.Map{
			"status":  "error",
			"code":    e.Code,
			"message": localizer.MustLocalize(e.LocalizeConfig),
		})
	} else if e, ok := err.(*types.ValidationError); ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"code":    e.Code,
			"message": localizer.MustLocalize(e.LocalizeConfig),
		})
	} else if e, ok := err.(*types.ValidationErrors); ok {
		localizedErrors := make(map[string][]map[string]string)
		for fieldName, validationErrors := range *e {
			localizedValidationErrors := make([]map[string]string, len(validationErrors))
			for i, validationError := range validationErrors {
				localizedValidationErrors[i] = map[string]string{
					"code":    validationError.Code,
					"message": localizer.MustLocalize(validationError.LocalizeConfig),
				}
			}
			localizedErrors[fieldName] = localizedValidationErrors
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"code":   "VALIDATION_ERROR",
			"errors": localizedErrors,
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":  "error",
		"code":    "INTERNAL_SERVER_ERROR",
		"message": localizer.MustLocalize(messages.ErrInternalServerErrorID),
	})
}

func LocalizationMiddleware(bundle *i18n.Bundle) fiber.Handler {
	return func(c *fiber.Ctx) error {
		preferredLanguage := c.Get("Accept-Language")

		localizer := i18n.NewLocalizer(bundle, preferredLanguage)

		c.Locals("localizer", localizer)

		return c.Next()
	}
}
