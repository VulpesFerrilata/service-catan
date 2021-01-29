package datamodel

type Dices []*Dice

func (d Dices) Roll() {
	for _, dice := range d {
		dice.Roll()
	}
}
