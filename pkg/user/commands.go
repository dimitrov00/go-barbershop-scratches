package user

import (
	messages "barbershop/creativo/i18n"
	"barbershop/creativo/internal/validation"
	"barbershop/creativo/pkg/auth"
	"barbershop/creativo/pkg/errors"
	"barbershop/creativo/pkg/types"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func HandleUserRegistration(
	checkIfExistsFn types.CheckIfUserExistsByEmailFn,
	sendEmailConfirmationFn types.SendEmailConfirmationFn,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input types.UserRegistrationInput
		if err := c.BodyParser(&input); err != nil {
			return err
		}

		validatedEmail, validationErr := validation.NewEmailAddress(input.Email)
		if validationErr != nil {
			return validationErr
		}

		if exists, err := checkIfExistsFn(validatedEmail); err != nil {
			return err
		} else if exists {
			return errors.ErrUserAlreadyExists
		}

		if err := sendEmailConfirmationFn(validatedEmail); err != nil {
			return err
		}

		localizer := c.Locals("localizer").(*i18n.Localizer)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": localizer.MustLocalize(messages.EmailVerificationCodeSentID(validatedEmail)),
		})
	}
}

func HandleUserConfirmation(
	passwordValidators []types.StringValidatorFn,
	checkIfExistsFn types.CheckIfUserExistsByEmailOrPhoneFn,
	hashPasswordFn types.HashPasswordFn,
	timeNowFn types.TimeNowFn,
	auditUpdaterFn types.AuditUpdaterFn,
	createUserFn types.CreateUserFn,
	sendWelcomeEmailFn types.SendWelcomeEmailFn,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input types.UserConfirmationInput
		if err := c.BodyParser(&input); err != nil {
			return err
		}

		validationErrors := make(types.ValidationErrors)

		validatedEmail, validationErr := validation.NewEmailAddress(input.Email)
		if validationErr != nil {
			validationErrors["email"] = []types.ValidationError{*validationErr}
		}

		validatedPassword, passwordErr := validation.NewPassword(input.Password, passwordValidators...)
		if passwordErr != nil {
			validationErrors["password"] = *passwordErr
		}

		validatedPhone, validationErr := validation.NewPhoneNumber(input.Phone)
		if validationErr != nil {
			validationErrors["phone"] = []types.ValidationError{*validationErr}
		}

		validatedName, nameErr := validation.NewName(input.FirstName, input.LastName)
		if nameErr != nil {
			validationErrors["name"] = *nameErr
		}

		// Check if there are validation errors and return them as a composed error
		if len(validationErrors) > 0 {
			return &validationErrors
		}

		// Check if user already exists
		if exists, err := checkIfExistsFn(validatedEmail, validatedPhone); err != nil {
			return err
		} else if exists {
			return errors.ErrUserAlreadyExists
		}

		hashedPassword, err := hashPasswordFn(validatedPassword)
		if err != nil {
			return err
		}

		audit := auditUpdaterFn(SystemUserID, timeNowFn(), true)(types.Audit{})

		newUser, creationErr := createUserFn(types.User{
			Name: *validatedName,
			Contact: types.ContactInfo{
				Email: validatedEmail,
				Phone: validatedPhone,
			},
			Password: &hashedPassword,
			Status:   UserStatusActive,
			Roles:    []types.UserRole{auth.UserRoleUser},
			Audit:    audit,
		})

		if creationErr != nil {
			return creationErr
		}

		fmt.Printf("New user created: %+v\n", newUser)
		// if err := sendWelcomeEmailFn(newUser); err != nil {
		// 	return err
		// }

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":  "success",
			"message": "User created successfully",
		})
	}
}

func HandleUserLogin(
	passwordValidators []types.StringValidatorFn,
	getUserByEmailFn types.GetUserByEmailFn,
	compareHashAndPasswordFn types.CompareHashAndPasswordFn,
	jwtSecret types.JWTSecret,
	generateJWTTokenFn types.GenerateJWTTokenFn,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input types.UserLoginInput
		if err := c.BodyParser(&input); err != nil {
			return err
		}

		validationErrors := make(types.ValidationErrors)

		validatedEmail, validationErr := validation.NewEmailAddress(input.Email)
		if validationErr != nil {
			validationErrors["email"] = []types.ValidationError{*validationErr}
		}

		validatedPassword, passwordErr := validation.NewPassword(input.Password)
		if passwordErr != nil {
			validationErrors["password"] = *passwordErr
		}

		// Check if there are validation errors and return them as a composed error
		if len(validationErrors) > 0 {
			return &validationErrors
		}

		user, err := getUserByEmailFn(validatedEmail)
		if err != nil {
			return err
		}

		if err := compareHashAndPasswordFn(*user.Password, validatedPassword); err != nil {
			return err
		}

		token, err := generateJWTTokenFn(*user, jwtSecret)
		if err != nil {
			return err
		}

		return c.JSON(types.UserLoginOutput{
			Token: token,
			User:  *user,
		})
	}
}
