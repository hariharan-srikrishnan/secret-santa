package emailer

import (
	"fmt"
	"net/smtp"
	"strings"

    "github.com/hariharan-srikrishnan/secret-santa/models"
)

var HOST string = "smtp.gmail.com"
var PORT string = "587" // default SMTP port

var MESSAGE string = `Dear %s,
You are the Secret Santa for %s!`

var SUBJECT string = "Secret Santa 2022"

type Mail struct {
    Sender  string
    To      []string
    Cc      []string
    Bcc     []string
    Subject string
    Body    string
}

func Send(c *Credentials, from, to *models.Player) error {	
    auth := smtp.PlainAuth("", c.Email, c.Password, HOST)
	toList := []string{to.GetEmail()}
	message := fmt.Sprintf(MESSAGE, to.GetName(), from.GetName())

	mail := Mail {
		Sender: c.Email,
		To: toList,
        Subject: SUBJECT,
		Body: message,
	}

	err := smtp.SendMail(HOST+":"+PORT, auth, c.Email, toList, []byte(BuildMessage(mail)))	
	if err != nil {
		fmt.Printf("Something went wrong with sending an email. Error: %v", err)
		return err
	}
	return nil
}


func BuildMessage(mail Mail) string {

    msg := ""
    msg += fmt.Sprintf("From: %s\r\n", mail.Sender)

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