package emailer

import (
	"fmt"
	"net/smtp"
	"strings"
    "time"
)

var HOST string = "smtp.gmail.com"
var PORT string = "587" // default SMTP port

var MESSAGE string = `Dear %s,
You are the Secret Santa for %s!`

var SUBJECT string = fmt.Sprintf("Secret Santa %d", time.Now().Year())

type Mail struct {
    Sender  string
    To      []string
    Cc      []string
    Bcc     []string
    Subject string
    Body    string
}

type Mailer struct {
    credentials *Credentials
    mail *Mail
}

const ERROR_SENDING_EMAIL string = "Something went wrong with sending an email. Error: %v"

/* OPTIONS FOR CONFIGURATION */

// Functional option for configuring Credentials.
func WithCredentials(c *Credentials) func(*Mailer) {
	return func(m *Mailer) {
		m.credentials = c
	}
}

// Functional option for configuring Message Content
func WithContent(content *Mail) func(*Mailer) {
    return func(m *Mailer) {
        m.mail = content
    }
}

func NewMailer(options ...func(*Mailer)) *Mailer {
    mailer := &Mailer{}
    for _, option := range options {
        option(mailer)
    }
    return mailer
}


func (m *Mailer) Send() error {
    auth := smtp.PlainAuth("", m.credentials.Email, m.credentials.Password, HOST)
    toList := m.mail.To

	err := smtp.SendMail(buildSocketAddress(HOST, PORT), 
                            auth, 
                            m.credentials.Email, 
                            toList, 
                            []byte(buildMessage(m.mail)),
                        )	
	return err

}

func buildMessage(mail *Mail) string {

    // DO NOT EDIT THIS, FORMAT HAS TO BE EXACTLY LIKE THIS FOR THE MAIL TO BE DELIVERED CORRECTLY
    msg := fmt.Sprintf("From: %s\r\n", mail.Sender)

    if len(mail.To) > 0 {
        msg += fmt.Sprintf("To: %s\r\n", mail.To[0])
    }

    if len(mail.Cc) > 0 {
        msg += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
    }

    msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
    msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

    return msg
}

func buildSocketAddress(host string, port string) string {
    return HOST+":"+PORT
}