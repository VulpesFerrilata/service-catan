package model

import (
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

func NewDice(game *Game) *Dice {
	dice := new(Dice)
	dice.Dice = new(datamodel.Dice)
	dice.Number = 1
	dice.SetGame(game)
	return dice
}

type Dice struct {
	*datamodel.Dice
	game *Game
}

func (d *Dice) SetGame(game *Game) {
	d.game = game
	game.dices.append(d)
}

func (d *Dice) Roll() int {
	d.Number = rand.Intn(6) + 1
	return d.Number
}
