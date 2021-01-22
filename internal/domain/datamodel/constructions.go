package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/pkg/errors"
)

func NewConstructionsFromConstructionModels(constructionModels []*model.Construction) (Constructions, error) {
	constructions := make(Constructions, 0)

	for _, constructionModel := range constructionModels {
		construction, err := NewConstructionFromConstructionModel(constructionModel)
		if err != nil {
			return nil, errors.Wrap(err, "datamodel.NewConstructionsFromConstructionModels")
		}
		constructions = append(constructions, construction)
	}

	return constructions, nil
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
