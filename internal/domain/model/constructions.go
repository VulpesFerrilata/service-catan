package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewConstructions(game *Game) Constructions {
	var constructions Constructions

	minQ := 1
	maxQ := 4
	for r := 0; r <= 6; r++ {
		for q := minQ; q <= maxQ; q++ {
			if (r > 0 && q != minQ && q != maxQ) || r > 3 {
				topConstruction := NewConstruction(game)
				topConstruction.Q = q
				topConstruction.R = r
				topConstruction.Location = datamodel.CL_TOP
				constructions.append(topConstruction)
			}
			if (r < 6 && q != minQ && q != maxQ) || r < 3 {
				botConstruction := NewConstruction(game)
				botConstruction.Q = q
				botConstruction.R = r
				botConstruction.Location = datamodel.CL_BOT
				constructions.append(botConstruction)
			}
		}

		if r < 3 {
			minQ--
		} else {
			maxQ--
		}
	}

	return constructions
}

type Constructions []*Construction

func (c *Constructions) append(construction *Construction) {
	*c = append(*c, construction)
}

func (c *Constructions) SetGame(game *Game) {
	for _, construction := range *c {
		construction.SetGame(game)
	}
}
