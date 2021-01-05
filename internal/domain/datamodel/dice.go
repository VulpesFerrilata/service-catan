package datamodel

import (
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/pkg/errors"
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
	id         int
	number     int
	game       *Game
	isModified bool
}

func (d *Dice) Roll() {
	d.number = rand.Intn(6) + 1
	d.isModified = true
}

func (d *Dice) Persist(f func(diceModel *model.Dice) error) error {
	diceModel := new(model.Dice)
	diceModel.ID = d.id
	if d.game != nil {
		diceModel.GameID = d.game.id
	}
	diceModel.Number = d.number

	if err := f(diceModel); err != nil {
		return errors.Wrap(err, "datamodel.Dice.Persist")
	}
	d.isModified = false
	d.isRemoved = false

	d.id = diceModel.ID
	d.number = diceModel.Number

	return nil
}
