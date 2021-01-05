package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
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

func NewDevelopmentCards() DevelopmentCards {
	var developmentCards DevelopmentCards

	developmentTypes := map[datamodel.DevelopmentType]int{
		datamodel.DT_KNIGHT:         14,
		datamodel.DT_VICTORY_POINTS: 5,
		datamodel.DT_ROAD_BUILDING:  2,
		datamodel.DT_YEAR_OF_PLENTY: 2,
		datamodel.DT_MONOPOLY:       2,
	}

	for developmentType, quantity := range developmentTypes {
		for i := 1; i <= quantity; i++ {
			developmentCard := new(DevelopmentCard)
			developmentCard.developmentType = developmentType
			developmentCards.append(developmentCard)
		}
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
