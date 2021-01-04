package datamodel

import (
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

type Dice struct {
	datamodel.Dice
	game *Game
}

func (d *Dice) SetGame(game *Game) {
	if game != nil {
		d.GameID = &game.id
	}
	d.game = game
}

func (d *Dice) Roll() int {
	d.Number = rand.Intn(6) + 1
	d.IsRolled = true
	return d.Number
}
