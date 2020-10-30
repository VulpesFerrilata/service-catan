package model

import (
	"errors"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

func NewGame() *Game {
	game := new(Game)
	game.Game = new(datamodel.Game)
	game.Status = datamodel.GS_WAITING
	return game
}

type Game struct {
	*datamodel.Game
	players   []*Player
	IsDeleted bool
}

func (g *Game) GetPlayerQuantity() int {
	quantity := 0

	for _, player := range g.players {
		if !player.IsDeleted {
			quantity++
		}
	}
	return quantity
}

func (g *Game) GetPlayerByUserId(userId uint) (*Player, error) {
	for _, player := range g.players {
		if player.UserID == userId && !player.IsDeleted {
			return player, nil
		}
	}
	return nil, errors.New("player is not exist")
}

func (g *Game) GetPlayers() []*Player {
	players := make([]*Player, 0)
	for _, player := range g.players {
		if !player.IsDeleted {
			players = append(players, player)
		}
	}
	return players
}

func (g *Game) AddPlayer(player *Player) {
	player.setGame(g)
	g.players = append(g.players, player)

	if g.Host == 0 {
		g.Host = player.ID
	}
}

func (g *Game) SwitchHost() {
	for _, player := range g.players {
		if !player.IsDeleted {
			g.Host = player.UserID
			return
		}
	}
}

func (g *Game) Init() {
	g.Status = datamodel.GS_STARTED
}
