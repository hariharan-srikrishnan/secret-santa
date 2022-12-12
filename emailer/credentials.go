package emailer 

type Credentials struct {
	Email    string
	Password string
}

func GetCredentials() *Credentials {
	return &Credentials{
		Email:    "YOUR_GMAIL_HERE",
		Password: "YOUR_PASSWORD_HERE",
		/* in case you have 2 factor authentication (2FA) enabled for this email,
			you will have to generate an application password at https://myaccount.google.com/apppasswords 
			and use that here
		*/
	}
}
