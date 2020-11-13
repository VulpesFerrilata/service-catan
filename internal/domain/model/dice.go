package model

import (
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

func NewDice(game *Game) *Dice {
	dice := new(Dice)
	dice.dice = new(datamodel.Dice)
	dice.dice.Number = 1
	dice.SetGame(game)
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

func (d *Dice) GetGameId() uint {
	return d.dice.GameID
}

func (d *Dice) setGameId(gameId uint) {
	if d.dice.GameID != gameId {
		d.dice.GameID = gameId
		d.isModified = true
	}
}

func (d *Dice) GetNumber() int {
	return d.dice.Number
}

func (d *Dice) IsModified() bool {
	return d.isModified
}

func (d *Dice) SetGame(game *Game) {
	d.game = game
	game.dices.append(d)
}

func (d *Dice) Roll() {
	d.dice.Number = rand.Intn(6) + 1
	d.isModified = true
}
