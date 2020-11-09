package model

func NewDices(game *Game) Dices {
	var dices Dices

	for i := 1; i <= 2; i++ {
		dice := NewDice(game)
		dices.append(dice)
	}

	return dices
}

type Dices []*Dice

func (d *Dices) append(dice *Dice) {
	*d = append(*d, dice)
}

func (d *Dices) SetGame(game *Game) {
	for _, dice := range *d {
		dice.SetGame(game)
	}
}

func (d *Dices) roll() int {
	sum := 0
	for _, dice := range *d {
		sum += dice.roll()
	}
	return sum
}

func (d *Dices) Total() int {
	total := 0
	for _, dice := range *d {
		total += dice.Number
	}
	return total
}
