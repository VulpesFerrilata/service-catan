package model

import (
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

func NewDice() *Dice {
	dice := new(Dice)
	dice.dice = new(datamodel.Dice)
	dice.dice.Number = 1
	return dice
}

type Dice struct {
	dice       *datamodel.Dice
	game       *Game
	isModified bool
}

func (d *Dice) GetDice() datamodel.Dice {
	return *d.dice
}

func (d *Dice) GetId() uint {
	return d.dice.ID
}

func (d *Dice) GetGameId() *uint {
	return d.dice.GameID
}

func (d *Dice) setGame(game *Game) {
	if game != nil {
		d.dice.GameID = &game.game.ID
		d.game = game
	}
}

func (d *Dice) GetNumber() int {
	return d.dice.Number
}

func (d *Dice) IsRolled() bool {
	return d.dice.IsRolled
}

func (d *Dice) IsModified() bool {
	return d.isModified
}

func (d *Dice) Roll() {
	d.dice.Number = rand.Intn(6) + 1
	d.dice.IsRolled = true
	d.isModified = true
}
