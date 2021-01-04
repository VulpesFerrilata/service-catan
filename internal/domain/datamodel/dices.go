package datamodel

func NewDices() Dices {
	var dices Dices

	for i := 1; i <= 2; i++ {
		dice := new(Dice)
		dice.Roll()
		dices.append(dice)
	}

	return dices
}

type Dices []*Dice

func (d *Dices) append(dice *Dice) {
	*d = append(*d, dice)
}

func (d Dices) IsRolled() bool {
	for _, dice := range d {
		if dice.IsRolled {
			return true
		}
	}
	return false
}

func (d *Dices) Roll() int {
	total := 0

	for _, dice := range *d {
		total += dice.Roll()
	}

	return total
}
