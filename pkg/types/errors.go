package types

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type (
	ValidationError struct {
		Code           string               `json:"code"`
		LocalizeConfig *i18n.LocalizeConfig `json:"localize_config"`
	}
	// ValidationErrors represents validation errors for a specific field.
	ValidationErrors map[string][]ValidationError
)

func (e ValidationError) Error() string {
	return e.LocalizeConfig.DefaultMessage.ID
}

func (e ValidationErrors) Error() string {
	var message string
	for _, errs := range e {
		for _, err := range errs {
			message += err.Error() + "\n"
		}
	}

	return message
}

type (
	ApiError struct {
		Code           string               `json:"code"`
		Status         int                  `json:"status"`
		LocalizeConfig *i18n.LocalizeConfig `json:"localize_config"`
	}
)

func (e ApiError) Error() string {
	return e.LocalizeConfig.DefaultMessage.ID
}
