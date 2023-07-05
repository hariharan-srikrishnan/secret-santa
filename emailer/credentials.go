package emailer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)


type Credentials struct {
	Email    string    `json:"email"`
	Password string    `json:"password"`
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

func readCredentialsFromConfigFile() (string, string, error) {
	credentials := Credentials{}
	fileBytes, err := os.ReadFile("config/credentials.txt")
	if err != nil {
		return "", "", err
	}
	err = json.Unmarshal(fileBytes, &credentials)
	if err != nil {
		return "", "", err
	}
	return credentials.Email, credentials.Password, nil
}

func GetCredentials(credentialsReadMode CredentialsReadMode) *Credentials {

	var email, password string
	var err error

	if credentialsReadMode == CredentialsReadMode_CommandLine {
		email, password, err = readCredentialsFromCommandLine()
		if err != nil {
			fmt.Println("error in reading/parsing credential inputs: ", err.Error())
			os.Exit(1)
		}
	} else if credentialsReadMode == CredentialsReadMode_ConfigFile {
		email, password, err = readCredentialsFromConfigFile()
		if err != nil {
			fmt.Println("error in reading/parsing credential inputs: ", err.Error())
			os.Exit(1)
		}
	}

	return &Credentials{
		Email:    email,
		Password: password,
	}
}

type CredentialsReadMode int 
const (
	CredentialsReadMode_CommandLine CredentialsReadMode = 0
	CredentialsReadMode_ConfigFile CredentialsReadMode = 1
)