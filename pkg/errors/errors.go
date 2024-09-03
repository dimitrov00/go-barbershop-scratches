package errors

import (
	messages "barbershop/creativo/i18n"
	"barbershop/creativo/pkg/types"

	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	ErrStringMinLength = func(fieldName string, minLength int) *types.ValidationError {
		return NewValidationError(messages.ErrStringTooShortID(fieldName, minLength))
	}
	ErrStringMaxLength = func(fieldName string, maxLength int) *types.ValidationError {
		return NewValidationError(messages.ErrStringTooLongID(fieldName, maxLength))
	}
	ErrStringLettersOnly = func(fieldName string) *types.ValidationError {
		return NewValidationError(messages.ErrStringLettersOnlyID(fieldName))
	}
	ErrPasswordTooShort = func(minLength int) *types.ValidationError {
		return NewValidationError(messages.ErrPasswordTooShortID(minLength))
	}
	ErrPasswordTooLong = func(maxLength int) *types.ValidationError {
		return NewValidationError(messages.ErrPasswordTooLongID(maxLength))
	}
	ErrPasswordMissingLowercase        = NewValidationError(messages.ErrPasswordNoLowerID)
	ErrPasswordMissingCapital          = NewValidationError(messages.ErrPasswordNoUpperID)
	ErrPasswordMissingNumber           = NewValidationError(messages.ErrPasswordNoNumberID)
	ErrPasswordMissingSpecialCharacter = NewValidationError(messages.ErrPasswordNoSpecialID)
	ErrInvalidEmailAddress             = NewValidationError(messages.ErrInvalidEmailID)
	ErrInvalidPhoneNumber              = NewValidationError(messages.ErrInvalidPhoneNumberID)

	ErrWrongCredentials    = NewApiError(fiber.StatusUnauthorized, messages.ErrWrongCredentialsID)
	ErrUserAlreadyExists   = NewApiError(fiber.StatusConflict, messages.ErrUserAlreadyExistsID)
	ErrInternalServerError = NewApiError(fiber.StatusInternalServerError, messages.ErrInternalServerErrorID)

	ErrSendingEmail = NewApiError(fiber.StatusInternalServerError, messages.ErrSendingEmailID)
)

func NewApiError(status int, localizeConfig *i18n.LocalizeConfig) *types.ApiError {
	return &types.ApiError{Status: status, Code: localizeConfig.DefaultMessage.ID, LocalizeConfig: localizeConfig}
}

func NewValidationError(localizeConfig *i18n.LocalizeConfig) *types.ValidationError {
	return &types.ValidationError{Code: localizeConfig.DefaultMessage.ID, LocalizeConfig: localizeConfig}
}
