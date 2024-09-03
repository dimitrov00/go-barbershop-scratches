package main

import (
	"barbershop/creativo/configs"
	"barbershop/creativo/i18n"
	"barbershop/creativo/pkg/common"
	"barbershop/creativo/pkg/web"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := configs.InitConfig(".env")
	i18nBundle := i18n.InitI18n()
	fmt.Printf("App config: %+v\n", config)

	code, _ := common.GenerateRandomCode(6)
	hashedCode, _ := common.HashString(code)

	fmt.Printf("Code: %s\nHashed code: %s\n", code, hashedCode)

	app := fiber.New()

	app.Use(
		web.LocalizationMiddleware(i18nBundle),
		web.RecoveryMiddleware,
		web.ErrorHandlingMiddleware,
	)

	web.SetupHandlers(app, config)

	log.Fatal(app.Listen(":3000"))
}
