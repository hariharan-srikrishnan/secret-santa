package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/hariharan-srikrishnan/secret-santa/emailer"
	"github.com/hariharan-srikrishnan/secret-santa/models"
	"github.com/hariharan-srikrishnan/secret-santa/permute"


)

func main() {
	players := createAllPlayers()
	// TODO: move this credentials retrieval mode to a CLI parameter
	credentials := emailer.GetCredentials(emailer.CredentialsReadMode_ConfigFile)

	// generate a derangement and map it to players
	santaTo := getGiftingMap(players)
	mailTemplate := &emailer.Mail{
		Sender: credentials.Email,
		Subject: emailer.SUBJECT,
	}


	for gifter, giftee := range santaTo {
		mailTemplate.To = []string{giftee.Email()}
		mailTemplate.Body = fmt.Sprintf(emailer.MESSAGE, gifter.Name(), giftee.Name())
		mailer := emailer.NewMailer(
			emailer.WithCredentials(credentials),
			emailer.WithContent(mailTemplate),
		)
		err := mailer.Send()
		if err != nil {
			log.Error().Msgf(emailer.ERROR_SENDING_EMAIL, err)
		}
	}
}

/*
createAllPlayers creates all the players from the config file located at config/players.txt
*/
func createAllPlayers() []*models.Player {
	players := []*models.Player{}

	f, err := os.Open("config/players.txt")
	if err != nil {
		log.Error().Msgf("Error opening file: %s", err)
		os.Exit(1)
	}
	reader := csv.NewReader(f)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Error().Msgf("Error reading file: %s", err)
			os.Exit(1)
		}
		player := models.CreatePlayer(row[0], row[1])
		players = append(players, player)
	}

	
	return players
}

func getGiftingMap(players []*models.Player) map[*models.Player]*models.Player {
	playerCount := len(players)
	santaMap := make(map[*models.Player]*models.Player)

	arrangement := permute.GetRandomDerangement(playerCount)

	for i := 0; i < playerCount; i++ {
		santaMap[players[i]] = players[arrangement[i]-1]
	}

	return santaMap
}
