package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

func NewDevelopmentCardsFromDevelopmentCardModels(developmentCardModels []*model.DevelopmentCard) DevelopmentCards {
	developmentCards := make(DevelopmentCards, 0)

	for _, developmentCardModel := range developmentCardModels {
		developmentCard := NewDevelopmentCardFromDevelopmentCardModel(developmentCardModel)
		developmentCards = append(developmentCards, developmentCard)
	}

	return developmentCards
}

type DevelopmentCards []*DevelopmentCard

func (dc *DevelopmentCards) append(developmentCard *DevelopmentCard) {
	*dc = append(*dc, developmentCard)
}

type DevelopmentCardFilterFunc func(developmentCard *DevelopmentCard) bool

func (dc DevelopmentCards) Filter(developmentCardFilterFunc DevelopmentCardFilterFunc) DevelopmentCards {
	var developmentCards DevelopmentCards
	for _, developmentCard := range dc {
		if developmentCardFilterFunc(developmentCard) {
			developmentCards.append(developmentCard)
		}
	}
	return developmentCards
}
