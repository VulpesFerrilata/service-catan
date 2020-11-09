package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewDevelopmentCards(game *Game) DevelopmentCards {
	var developmentCards DevelopmentCards

	for i := 1; i <= 14; i++ {
		knightCard := NewDevelopmentCard(game, datamodel.DT_KNIGHT)
		developmentCards.append(knightCard)
	}

	developmentTypes := map[datamodel.DevelopmentType]int{
		datamodel.DT_VICTORY_POINTS: 5,
		datamodel.DT_ROAD_BUILDING:  2,
		datamodel.DT_YEAR_OF_PLENTY: 2,
		datamodel.DT_MONOPOLY:       2,
	}

	for developmentType, quantity := range developmentTypes {
		for i := 1; i <= quantity; i++ {
			developmentCard := NewDevelopmentCard(game, developmentType)
			developmentCards.append(developmentCard)
		}
	}

	return developmentCards
}

type DevelopmentCards []*DevelopmentCard

func (dc *DevelopmentCards) append(developmentCard *DevelopmentCard) {
	*dc = append(*dc, developmentCard)
}

func (dc *DevelopmentCards) SetGame(game *Game) {
	for _, developmentCard := range *dc {
		developmentCard.SetGame(game)
	}
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
