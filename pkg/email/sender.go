package email

import (
	"barbershop/creativo/configs"
	"barbershop/creativo/pkg/types"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func ConfigureEmailSender(config *configs.EmailConfig) types.EmailSenderFn {
	auth := smtp.PlainAuth("", string(config.FromAddress), config.Password, config.Host)

	return func(
		subject string,
		content string,
		to types.EmailAddresses,
		cc types.EmailAddresses,
		bcc types.EmailAddresses,
		attachFiles []string,
	) error {
		e := email.NewEmail()

		e.From = fmt.Sprintf("%s <%s>", config.FromName, config.FromAddress)
		e.Subject = subject
		e.HTML = []byte(content)
		e.To = to.Strings()
		e.Cc = cc.Strings()
		e.Bcc = bcc.Strings()

		for _, file := range attachFiles {
			if _, err := e.AttachFile(file); err != nil {
				return err
			}
		}

		return e.Send(fmt.Sprintf("%s:%d", config.Host, config.Port), auth)
	}
}
