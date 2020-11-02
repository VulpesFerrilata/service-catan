package model

func NewDices() Dices {
	var dices Dices

	for i := 1; i <= 2; i++ {
		dice := NewDice()
		dices.Append(dice)
	}

	return dices
}

type Dices []*Dice

func (d Dices) Append(dice *Dice) {
	d = append(d, dice)
}

func (d Dices) Remove() {
	for _, dice := range d {
		dice.Remove()
	}
}

func (d Dices) Roll() {
	for _, dice := range d {
		dice.Roll()
	}
}

func (d Dices) Total() int {
	total := 0
	for _, dice := range d {
		total += dice.Number
	}
	return total
}
