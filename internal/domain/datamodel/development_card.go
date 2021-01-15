package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
)

func NewDevelopmentCardFromDevelopmentCardModel(developmentCardModel *model.DevelopmentCard) *DevelopmentCard {
	developmentCard := new(DevelopmentCard)
	developmentCard.id = developmentCardModel.ID
	developmentCard.developmentType = developmentCardModel.DevelopmentType
	developmentCard.isModified = false
	developmentCard.isRemoved = false
	return developmentCard
}

type DevelopmentCard struct {
	base
	id              uuid.UUID
	developmentType model.DevelopmentType
	game            *Game
	playerID        *uuid.UUID
}

func (dc DevelopmentCard) ToModel() *model.DevelopmentCard {
	developmentCardModel := new(model.DevelopmentCard)
	developmentCardModel.ID = dc.id
	if dc.game != nil {
		developmentCardModel.GameID = dc.game.id
	}
	developmentCardModel.DevelopmentType = dc.developmentType
	developmentCardModel.PlayerID = dc.playerID
	return developmentCardModel
}
