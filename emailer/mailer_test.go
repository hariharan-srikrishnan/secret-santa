package emailer

import (
	"reflect"
	"testing"
)

func TestMailerCreationMustHaveCorrectCredentials(t *testing.T) {
	c := &Credentials{
		Email: "test@xyz",
		Password: "testpassword",
	}
	mailer := NewMailer(
		WithCredentials(c),
	)

	if c.Email != mailer.credentials.Email || 
		c.Password != mailer.credentials.Password {
		t.Error("mailer object is not initialized correctly with given credentials")
		t.Fail()
	}
}

func TestMailerCreationMustHaveCorrectMailContent(t *testing.T) {
	
	content := &Mail{
		Sender: "sender@sender",
		To: []string{"a@a, b@b, c@c"},
		Cc: []string{"cc1@cc, cc2@cc"},
		Bcc: []string{"bcc1@bcc, bcc2@bcc"},
		Subject: "subject",
		Body: "body",
	}
	mailer := NewMailer(
		WithContent(content),
	)

	if content.Sender != mailer.mail.Sender {
		t.Error("Sender field in mailer is not populated correctly")
	}

	if !reflect.DeepEqual(content.To, mailer.mail.To) {
		t.Error("To field in mailer is not populated correctly")
	}

	if !reflect.DeepEqual(content.Cc, mailer.mail.Cc) {
		t.Error("Cc field in mailer is not populated correctly")
	}

	if !reflect.DeepEqual(content.Bcc, mailer.mail.Bcc) {
		t.Error("Bcc field in mailer is not populated correctly")
	}

	if content.Subject != mailer.mail.Subject {
		t.Error("Subject field in mailer is not populated correctly")
	}

	if content.Body != mailer.mail.Body {
		t.Error("Content field in mailer is not populated correctly")
	}
}