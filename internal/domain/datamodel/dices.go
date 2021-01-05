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

func NewDices() Dices {
	var dices Dices

	for i := 1; i <= 2; i++ {
		dice := new(Dice)
		dice.Roll()
	}

	return dices
}

type Dices []*Dice

func (d Dices) Roll() {
	for _, dice := range d {
		dice.Roll()
	}
}
