package model

import (
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
	players   Players
	dices     Dices
	isRemoved bool
}

func (g Game) FilterPlayers(f func(p *Player) bool) Players {
	var players Players
	for _, player := range g.players {
		if f(player) {
			players.Append(player)
		}
	}
	return players
}

func (g *Game) AddPlayer(player *Player) {
	player.setGame(g)
	g.players.Append(player)
}

func (g *Game) GetDices() Dices {
	return g.dices
}

func (g *Game) AddDice(dice *Dice) {
	dice.setGame(g)
	g.dices.Append(dice)
}

func (g Game) IsRemoved() bool {
	return g.isRemoved
}

func (g *Game) Remove() {
	g.isRemoved = true
	g.players.Remove()
	g.dices.Remove()
}

func (g *Game) Init() {
	g.Status = datamodel.GS_STARTED

	dices := NewDices()
	for _, dice := range dices {
		g.AddDice(dice)
	}
}
