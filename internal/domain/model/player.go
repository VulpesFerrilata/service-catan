package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewPlayer(userId uint) *Player {
	player := new(Player)
	player.Player = new(datamodel.Player)
	player.UserID = userId
	return player
}

type Player struct {
	*datamodel.Player
}
