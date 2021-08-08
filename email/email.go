package email

import (
	"os"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendEmail(from string, to []string, subject string, body string, isAuth bool) error {
	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)

	m.SetBody("text/html", body)

	var dialer *gomail.Dialer
	if isAuth {
		dialer = gomail.NewDialer(viper.GetString("mail.smtp-host"), viper.GetInt("mail.smtp-port"), from, os.Getenv("EMAIL_APIKEY"))
	} else {
		dialer = &gomail.Dialer{
			Host: viper.GetString("mail.smtp-host"),
			Port: viper.GetInt("mail.smtp-port"),
		}
	}
	if err := dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
