package datamodel

import (
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
)

func NewDiceFromDiceModel(diceModel *model.Dice) *Dice {
	dice := new(Dice)
	dice.id = diceModel.ID
	dice.number = diceModel.Number
	dice.isModified = false
	dice.isRemoved = false
	return dice
}

type Dice struct {
	base
	id         uuid.UUID
	number     int
	game       *Game
	isModified bool
}

func (d *Dice) Roll() {
	d.number = rand.Intn(6) + 1
	d.isModified = true
}

func (d Dice) ToModel() *model.Dice {
	diceModel := new(model.Dice)
	diceModel.ID = d.id
	if d.game != nil {
		diceModel.GameID = d.game.id
	}
	diceModel.Number = d.number
	return diceModel
}
