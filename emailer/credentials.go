package emailer 

import (
	"fmt"
	"golang.org/x/term"
	"bufio"
	"os"
	"strings"
	"syscall"
)


type Credentials struct {
	Email    string
	Password string
}


/* 
	in case you have 2 factor authentication (2FA) enabled for this email,
	you will have to generate an application password at https://myaccount.google.com/apppasswords 
	and use that here
*/
func readCredentialsFromCommandLine() (string, string, error) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter Username: ")
    username, err := reader.ReadString('\n')
    if err != nil {
        return "", "", err
    }

    fmt.Print("Enter Password: ")
    bytePassword, err := term.ReadPassword(int(syscall.Stdin))
    if err != nil {
        return "", "", err
    }

    password := string(bytePassword)
    return strings.TrimSpace(username), strings.TrimSpace(password), nil
}

func GetCredentials() *Credentials {

	email, password, err := readCredentialsFromCommandLine()
	if err != nil {
		fmt.Println("error in reading/parsing credential inputs: ", err.Error())
	}

	return &Credentials{
		Email:    email,
		Password: password,
	}
}
