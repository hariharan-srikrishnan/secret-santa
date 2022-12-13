package main

import (
	"santa.com/secret/emailer"
	"santa.com/secret/models"
	"santa.com/secret/permute"
)

var NUM_PLAYERS = 8

func main() {
	players := createAllPlayers()
	credentials := emailer.GetCredentials()

	// generate a derangement and map it to players
	santaTo := permute.GetPlayerMap(players)

	for player := range santaTo {
		emailer.Send(credentials, player, santaTo[player])
	}
}

func createAllPlayers() []*models.Player {
	players := []*models.Player{}

	/* 
	initialize all players one-by-one by calling
	models.CreatePlayer() and append to players
	*/
	return players
}
