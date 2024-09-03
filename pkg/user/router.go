package user

import (
	"barbershop/creativo/configs"
	"barbershop/creativo/internal/validation"
	"barbershop/creativo/pkg/auth"
	"barbershop/creativo/pkg/common"
	"barbershop/creativo/pkg/types"
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SetupHandlers(app *fiber.App, db *sql.DB, config *configs.AuthConfig) error {
	userRoutes := app.Group("/users")

	// Arrange dependencies
	checkUserExistsByEmailOrPhoneFn := func(ea types.EmailAddress, pn types.PhoneNumber) (bool, error) {
		return true, nil
	}
	checkUserExistsByEmailFn := func(ea types.EmailAddress) (bool, error) {
		return true, nil
	}
	createUserFn := func(u types.User) (types.User, error) {
		return u, nil
	}
	sendEmailConfirmationFn := func(ea types.EmailAddress) error {
		return nil
	}
	testPass, _ := bcrypt.GenerateFromPassword([]byte("1Aabcdefg-"), bcrypt.DefaultCost)
	hashedPass := types.HashedPassword(testPass)
	getUserByEmailFn := func(ea types.EmailAddress) (*types.User, error) {
		return &types.User{
			Contact: types.ContactInfo{
				Email: ea,
			},
			Password: &hashedPass,
		}, nil
	}

	userRoutes.Post("/register", HandleUserRegistration(
		checkUserExistsByEmailFn,
		sendEmailConfirmationFn,
	))

	userRoutes.Post("/register/confirm", HandleUserConfirmation(
		validation.DefaultPasswordValidators,
		checkUserExistsByEmailOrPhoneFn,
		auth.HashPassword,
		time.Now,
		common.NewAuditUpdater,
		createUserFn,
		nil,
	))

	userRoutes.Post("/login", HandleUserLogin(
		validation.DefaultPasswordValidators,
		getUserByEmailFn,
		auth.CompareHashAndPassword,
		config.JWTSecret,
		auth.GenerateJWTToken,
	))

	return nil
}
