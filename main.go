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
	santaTo := getPlayerMap(players)

	for player := range santaTo {
		emailer.Send(credentials, player, santaTo[player])
	}
}

func createAllPlayers() []*models.Player {
	players := []*models.Player{}

	/* 
	initialize all players one-by-one by calling
	models.CreatePlayer("<player name>", "<player email>") and append to players
	*/
	return players
}

func getPlayerMap(players []*models.Player) map[*models.Player]*models.Player {
	playerCount := len(players)
	santaMap := make(map[*models.Player]*models.Player)

	arrangement := GetDerangements(playerCount)

	for i := 0; i < playerCount; i++ {
		santaMap[players[i]] = players[arrangement[i]-1]
	}

	return santaMap
}
