package datamodel

import "github.com/VulpesFerrilata/catan/internal/domain/model"

func NewDicesFromDiceModels(diceModels []*model.Dice) Dices {
	dices := make(Dices, 0)

	for _, diceModel := range diceModels {
		dice := NewDiceFromDiceModel(diceModel)
		dices = append(dices, dice)
	}

	return dices
}

type Dices []*Dice

func (d Dices) Roll() {
	for _, dice := range d {
		dice.Roll()
	}
}
