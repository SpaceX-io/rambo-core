package mail

import (
	"strings"

	"gopkg.in/gomail.v2"
)

type Options struct {
	Host    string
	Port    int
	User    string
	Pass    string
	To      string
	Subject string
	Body    string
}

func Send(o *Options) error {
	m := gomail.NewMessage()

	m.SetHeader("From", o.User)

	toMultiMail := strings.Split(o.To, ",")
	m.SetHeader("To", toMultiMail...)

	m.SetHeader("Subject", o.Subject)

	m.SetBody("text/html", o.Body)

	d := gomail.NewDialer(o.Host, o.Port, o.User, o.Pass)

	return d.DialAndSend(m)
}
