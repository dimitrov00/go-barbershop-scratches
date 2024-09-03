package validation

import (
	"barbershop/creativo/pkg/errors"
	"barbershop/creativo/pkg/types"
	"regexp"
)

func IsMatchingRegex(regexPattern string, s string) bool {
	regex := regexp.MustCompile(regexPattern)

	return regex.MatchString(s)
}

func IsValidEmail(email string) bool {
	return IsMatchingRegex(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
}

func NewEmailAddress(emailAddress string) (types.EmailAddress, *types.ValidationError) {
	if !IsValidEmail(emailAddress) {
		return "", errors.ErrInvalidEmailAddress
	}

	return types.EmailAddress(emailAddress), nil
}

func WithMinLength(minLength int, err *types.ValidationError) types.StringValidatorFn {
	return func(s string) *types.ValidationError {
		if len(s) < minLength {
			return err
		}

		return nil
	}
}

func WithMaxLength(maxLength int, err *types.ValidationError) types.StringValidatorFn {
	return func(s string) *types.ValidationError {
		if len(s) > maxLength {
			return err
		}

		return nil
	}
}

func MatchesRegex(regexPattern string, err *types.ValidationError) types.StringValidatorFn {
	return func(s string) *types.ValidationError {
		if IsMatchingRegex(regexPattern, s) {
			return nil
		}

		return err
	}
}

func WithAtLeastOneLowercase(err *types.ValidationError) types.StringValidatorFn {
	return MatchesRegex(`[а-яa-z]`, err)
}

func WithAtLeastOneCapital(err *types.ValidationError) types.StringValidatorFn {
	return MatchesRegex(`[А-ЯA-Z]`, err)
}

func WithAtLeastOneNumber(err *types.ValidationError) types.StringValidatorFn {
	return MatchesRegex(`[0-9]`, err)
}

func WithAtLeastOneSpecialCharacter(err *types.ValidationError) types.StringValidatorFn {
	return MatchesRegex(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`, err)
}

func WithLettersOnly(err *types.ValidationError) types.StringValidatorFn {
	return MatchesRegex(`^[a-zA-Zа-яА-Я]+$`, err)
}

var DefaultPasswordValidators = []types.StringValidatorFn{
	WithMinLength(8, errors.ErrPasswordTooShort(8)),
	WithMaxLength(64, errors.ErrPasswordTooLong(64)),
	WithAtLeastOneLowercase(errors.ErrPasswordMissingLowercase),
	WithAtLeastOneCapital(errors.ErrPasswordMissingCapital),
	WithAtLeastOneNumber(errors.ErrPasswordMissingNumber),
	WithAtLeastOneSpecialCharacter(errors.ErrPasswordMissingSpecialCharacter),
}

func NewPassword(password string, validators ...types.StringValidatorFn) (types.Password, *[]types.ValidationError) {
	validationErrors := make([]types.ValidationError, 0, len(validators))

	for _, validator := range validators {
		if err := validator(password); err != nil {
			validationErrors = append(validationErrors, *err)
		}
	}

	if len(validationErrors) > 0 {
		return "", &validationErrors
	}

	return types.Password(password), nil
}

func NewString50(fieldName, s string) (types.String50, *types.ValidationError) {
	if err := WithMaxLength(50, errors.ErrStringMaxLength(fieldName, 50))(s); err != nil {
		return "", err
	}

	return types.String50(s), nil
}

func NewString500(fieldName, s string) (types.String500, *types.ValidationError) {
	if err := WithMaxLength(500, errors.ErrStringMaxLength(fieldName, 500))(s); err != nil {
		return "", err
	}

	return types.String500(s), nil
}

func NewString50LettersOnly(fieldName, s string) (types.String50LettersOnly, *types.ValidationError) {
	validators := []types.StringValidatorFn{
		WithLettersOnly(errors.ErrStringLettersOnly(fieldName)),
		WithMaxLength(50, errors.ErrStringMaxLength(fieldName, 50)),
	}

	for _, validator := range validators {
		if err := validator(s); err != nil {
			return "", err
		}
	}

	return types.String50LettersOnly(s), nil
}

func NewName(fName, lName string) (*types.Name, *[]types.ValidationError) {
	validationErrors := make([]types.ValidationError, 0, 2)

	newFName, err := NewString50LettersOnly("FirstName", fName)
	if err != nil {
		validationErrors = append(validationErrors, *err)
	}

	newLName, err := NewString50LettersOnly("LastName", lName)
	if err != nil {
		validationErrors = append(validationErrors, *err)
	}

	if len(validationErrors) > 0 {
		return nil, &validationErrors
	}

	return &types.Name{
		FirstName: newFName,
		LastName:  newLName,
	}, nil
}
