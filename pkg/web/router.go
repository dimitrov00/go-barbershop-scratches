package web

import (
	"barbershop/creativo/configs"
	"barbershop/creativo/pkg/email"
	"barbershop/creativo/pkg/types"
	"barbershop/creativo/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func SetupHandlers(app *fiber.App, config *configs.AppConfig) error {
	emailSender := email.ConfigureEmailSender(&config.Email)

	app.Post("/test-email", func(c *fiber.Ctx) error {
		return emailSender(
			"Test email",
			"This is a test email",
			[]types.EmailAddress{
				"dimitar.dimitrov.21f@gmail.com",
			},
			nil,
			nil,
			nil,
		)
	})

	user.SetupHandlers(app, nil, &config.Auth)

	return nil
}
