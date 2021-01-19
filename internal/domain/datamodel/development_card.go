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

	developmentCard.SetModelState(Added)

	return developmentCard, nil
}

func NewDevelopmentCardFromDevelopmentCardModel(developmentCardModel *model.DevelopmentCard) *DevelopmentCard {
	developmentCard := new(DevelopmentCard)
	developmentCard.id = developmentCardModel.ID
	developmentCard.developmentType = developmentCardModel.DevelopmentType
	developmentCard.playerID = developmentCardModel.PlayerID

	developmentCard.SetModelState(Unchanged)

	return developmentCard
}

type DevelopmentCard struct {
	base
	id              uuid.UUID
	developmentType model.DevelopmentType
	playerID        *uuid.UUID
	game            *Game
}

func (dc *DevelopmentCard) ToModel() *model.DevelopmentCard {
	dc.SetModelState(Unchanged)

	developmentCardModel := new(model.DevelopmentCard)
	developmentCardModel.ID = dc.id
	if dc.game != nil {
		developmentCardModel.GameID = dc.game.id
	}
	developmentCardModel.DevelopmentType = dc.developmentType
	developmentCardModel.PlayerID = dc.playerID
	return developmentCardModel
}
