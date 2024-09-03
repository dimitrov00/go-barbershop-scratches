package types

import (
	"time"
)

type (
	String50            string
	String500           string
	String50LettersOnly string
	PhoneNumber         string
	ImageURL            string
	ErrCode             string
)

type (
	Audit struct {
		CreatedAt time.Time `json:"created_at"`
		CreatedBy UserID    `json:"created_by"`

		UpdatedAt time.Time `json:"updated_at"`
		UpdatedBy UserID    `json:"updated_by"`
	}

	Pagination struct {
		Total    uint    `json:"total"`
		Limit    uint    `json:"limit"`
		Offset   uint    `json:"offset"`
		Next     *string `json:"next,omitempty"`
		Previous *string `json:"prev,omitempty"`
	}

	Interval struct {
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	}

	Location struct {
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
		AddressLine string  `json:"address_line"`
	}
)

type (
	LoggerFn func(format string, v ...interface{})

	StringValidatorFn      func(s string) *ValidationError
	ValidateEmailAddressFn func(string) (EmailAddress, error)
	ValidatePhoneNumberFn  func(string) (PhoneNumber, error)

	TimeNowFn              func() time.Time
	AuditChangeProcessorFn func(Audit) Audit
	AuditUpdaterFn         func(updaterID UserID, now time.Time, forNewRecord bool) AuditChangeProcessorFn

	SendEmailFn func(to []EmailAddress, subject, body string) error
)
