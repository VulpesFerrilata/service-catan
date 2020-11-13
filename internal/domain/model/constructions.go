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
				topConstruction.construction.Q = q
				topConstruction.construction.R = r
				topConstruction.construction.Location = datamodel.CL_TOP
				constructions.append(topConstruction)
			}
			if (r < 6 && q != minQ && q != maxQ) || r < 3 {
				botConstruction := NewConstruction(game)
				botConstruction.construction.Q = q
				botConstruction.construction.R = r
				botConstruction.construction.Location = datamodel.CL_BOT
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

type ConstructionFilterFunc func(construction *Construction) bool

func (c Constructions) Filter(constructionFilterFunc ConstructionFilterFunc) Constructions {
	var constructions Constructions

	for _, construction := range c {
		if constructionFilterFunc(construction) {
			constructions.append(construction)
		}
	}

	return constructions
}
