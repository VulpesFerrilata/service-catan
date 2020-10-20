package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type Game struct {
	*datamodel.Game
	Players []*Player
}

func (g *Game) AddPlayer(player *Player) {
	player.GameID = g.ID
	g.Players = append(g.Players, player)
}
