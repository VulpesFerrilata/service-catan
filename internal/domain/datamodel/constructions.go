package datamodel

import (
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
