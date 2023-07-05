# Prerequisites:
- go 1.16 (might work with earlier versions of go but I haven't checked)
- A Gmail account 

# Running the application:
```go run main.go```

# Configurations required:
- Initialize all the players playing by editing the file `players.txt` in the `config` directory
- Initialize the game invite sender's email account by editing the file `credentials.txt` in the `config` directory 
- If you have 2FA enabled on the email account that sends out the emails, you will have to generate an app password at https://myaccount.google.com/apppasswords and use the generated password
