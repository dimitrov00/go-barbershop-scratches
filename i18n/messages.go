package i18n

import (
	"barbershop/creativo/pkg/types"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	ErrStringTooShortID = func(fieldName string, minLength int) *i18n.LocalizeConfig {
		return &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "STRING_TOO_SHORT",
				One:   "The {{.FieldName}} must be at least {{.MinLength}} character long",
				Other: "The {{.FieldName}} must be at least {{.MinLength}} characters long",
			},
			TemplateData: map[string]interface{}{
				"FieldName": fieldName,
				"MinLength": minLength,
			},
			PluralCount: 2,
		}
	}
	ErrStringTooLongID = func(fieldName string, maxLength int) *i18n.LocalizeConfig {
		return &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "STRING_TOO_LONG",
				One:   "The {{.FieldName}} must be at most {{.MaxLength}} character long",
				Other: "The {{.FieldName}} must be at most {{.MaxLength}} characters long",
			},
			TemplateData: map[string]interface{}{
				"FieldName": fieldName,
				"MaxLength": maxLength,
			},
			PluralCount: 2,
		}
	}
	ErrStringLettersOnlyID = func(fieldName string) *i18n.LocalizeConfig {
		return &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "STRING_LETTERS_ONLY",
				Other: "The {{.FieldName}} must contain only letters",
			},
			TemplateData: map[string]interface{}{
				"FieldName": fieldName,
			},
		}
	}
	ErrInvalidEmailID       = &i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "INVALID_EMAIL", Other: "Invalid email address"}}
	ErrEmailAlreadyExistsID = &i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "EMAIL_ALREADY_EXISTS", Other: "Email address already exists"}}

	ErrPasswordTooShortID = func(minLength int) *i18n.LocalizeConfig {
		return &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "PASSWORD_TOO_SHORT",
				Other: "Password must be at least {{.MinLength}} characters long",
			},
			TemplateData: map[string]interface{}{
				"MinLength": minLength,
			},
		}
	}
	ErrPasswordTooLongID = func(maxLength int) *i18n.LocalizeConfig {
		return &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "PASSWORD_TOO_LONG",
				Other: "Password must be at most {{.MaxLength}} characters long",
			},
			TemplateData: map[string]interface{}{
				"MaxLength": maxLength,
			},
		}
	}
	ErrPasswordNoLowerID   = &i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "PASSWORD_NO_LOWER", Other: "Password must contain at least one lowercase letter"}}
	ErrPasswordNoUpperID   = &i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "PASSWORD_NO_UPPER", Other: "Password must contain at least one capital letter"}}
	ErrPasswordNoNumberID  = &i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "PASSWORD_NO_NUMBER", Other: "Password must contain at least one number"}}
	ErrPasswordNoSpecialID = &i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "PASSWORD_NO_SPECIAL", Other: "Password must contain at least one special character"}}

	ErrWrongCredentialsID = &i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "WRONG_CREDENTIALS", Other: "Invalid email or password"}}

	ErrInvalidPhoneNumberID = &i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "INVALID_PHONE_NUMBER", Other: "Invalid phone number"}}
	ErrInvalidNameID        = &i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "INVALID_NAME", Other: "Invalid name"}}

	ErrUserAlreadyExistsID = &i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "USER_ALREADY_EXISTS", Other: "User already exists"}}

	ErrInternalServerErrorID = &i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "INTERNAL_SERVER_ERROR", Other: "Something went wrong"}}

	ErrSendingEmailID = &i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "ERROR_SENDING_EMAIL", Other: "Error occurred while sending email"}}

	EmailVerificationCodeSentID = func(ea types.EmailAddress) *i18n.LocalizeConfig {
		return &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "EMAIL_VERIFICATION_CODE_SENT",
				Other: "Email verification code sent to {{.EmailAddress}}",
			},
			TemplateData: map[string]interface{}{
				"EmailAddress": ea,
			},
			PluralCount: 2,
		}
	}

	WelcomeEmailSubjectID = &i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "WELCOME_EMAIL_SUBJECT", Other: "Welcome to Creativo"}}
	WelcomeEmailBodyID    = func(firstName string) *i18n.LocalizeConfig {
		return &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "WELCOME_EMAIL_BODY",
				Other: "Welcome to Creativo, {{.FirstName}}!",
			},
			TemplateData: map[string]interface{}{
				"FirstName": firstName,
			},
		}
	}
)
