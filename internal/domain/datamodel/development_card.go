package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewDevelopmentCard(developmentType model.DevelopmentType) (*DevelopmentCard, error) {
	developmentCard := new(DevelopmentCard)

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewDevelopmentCard")
	}
	developmentCard.id = id

	developmentCard.developmentType = developmentType
	developmentCard.playerID = nil
	return developmentCard, nil
}

func NewDevelopmentCardFromDevelopmentCardModel(developmentCardModel *model.DevelopmentCard) *DevelopmentCard {
	developmentCard := new(DevelopmentCard)
	developmentCard.id = developmentCardModel.ID
	developmentCard.developmentType = developmentCardModel.DevelopmentType
	developmentCard.playerID = developmentCardModel.PlayerID
	return developmentCard
}

type DevelopmentCard struct {
	id              uuid.UUID
	developmentType model.DevelopmentType
	playerID        *uuid.UUID
	game            *Game
}

func (d DevelopmentCard) ToModel() *model.DevelopmentCard {
	developmentCardModel := new(model.DevelopmentCard)
	developmentCardModel.ID = d.id

	if d.game != nil {
		developmentCardModel.GameID = d.game.id
	}

	developmentCardModel.DevelopmentType = d.developmentType
	developmentCardModel.PlayerID = d.playerID
	return developmentCardModel
}
