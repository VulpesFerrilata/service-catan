package model

import (
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

func NewDice() *Dice {
	dice := new(Dice)
	dice.Number = 1
	return &Dice{}
}

type Dice struct {
	*datamodel.Dice
	game      *Game
	isRemoved bool
}

func (d *Dice) setGame(game *Game) {
	d.GameID = game.ID
	d.game = game
}

func (d Dice) IsRemoved() bool {
	return d.isRemoved
}

func (d *Dice) Remove() {
	d.isRemoved = true
}

func (d *Dice) Roll() {
	d.Number = rand.Intn(6) + 1
	d.IsRolled = true
}
