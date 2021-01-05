package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

func NewConstructionsFromConstructionModels(constructionModels []*model.Construction) Constructions {
	constructions := make(Constructions, 0)

	for _, constructionModel := range constructionModels {
		construction := NewConstructionFromConstructionModel(constructionModel)
		constructions = append(constructions, construction)
	}

	return constructions
}

func NewConstructions() Constructions {
	var constructions Constructions

	minQ := 1
	maxQ := 4
	for r := 0; r <= 6; r++ {
		for q := minQ; q <= maxQ; q++ {
			if (r > 0 && q != minQ && q != maxQ) || r > 3 {
				topConstruction := new(Construction)
				topConstruction.Q = q
				topConstruction.R = r
				topConstruction.Location = datamodel.CL_TOP
				constructions.append(topConstruction)
			}
			if (r < 6 && q != minQ && q != maxQ) || r < 3 {
				botConstruction := new(Construction)
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

type ConstructionFilterFunc func(construction *Construction) bool

func (c Constructions) Filter(constructionFilterFunc ConstructionFilterFunc) Constructions {
	var constructions Constructions

	for _, construction := range c {
		if constructionFilterFunc(construction) {
			constructions = append(constructions, construction)
		}
	}

	return constructions
}
