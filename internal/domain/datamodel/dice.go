package datamodel

import (
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewDice(maxNumber int) (*Dice, error) {
	dice := new(Dice)

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewDice")
	}
	dice.id = id

	dice.maxNumber = maxNumber
	dice.Roll()
	return dice, nil
}

func NewDiceFromDiceModel(diceModel *model.Dice) *Dice {
	dice := new(Dice)
	dice.id = diceModel.ID
	dice.number = diceModel.Number
	dice.maxNumber = diceModel.MaxNumber

	return dice
}

type Dice struct {
	id        uuid.UUID
	number    int
	maxNumber int
	game      *Game
}

func (d *Dice) Roll() {
	d.number = rand.Intn(d.maxNumber) + 1
}

func (d Dice) ToModel() *model.Dice {
	diceModel := new(model.Dice)
	diceModel.ID = d.id

	if d.game != nil {
		diceModel.GameID = d.game.id
	}

	diceModel.Number = d.number
	diceModel.MaxNumber = d.maxNumber
	return diceModel
}
