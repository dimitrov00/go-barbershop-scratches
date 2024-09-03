package types

import "time"

type (
	EmailAddress   string
	EmailAddresses []EmailAddress

	EmailConfirmation struct {
		Email     EmailAddress
		Code      string
		ExpiresAt time.Time
	}
)

func (e EmailAddress) String() string {
	return string(e)
}

func (e EmailAddresses) Strings() []string {
	var result []string
	for _, email := range e {
		result = append(result, email.String())
	}
	return result
}

type (
	EmailSenderFn func(
		subject string,
		content string,
		to EmailAddresses,
		cc EmailAddresses,
		bcc EmailAddresses,
		attachFiles []string,
	) error

	SendEmailConfirmationFn func(EmailAddress) error
	SendWelcomeEmailFn      func(User, EmailSenderFn) error
)
