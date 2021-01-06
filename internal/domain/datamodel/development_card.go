package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/pkg/errors"
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
	id              int
	developmentType model.DevelopmentType
	game            *Game
	player          *Player
}

func (dc *DevelopmentCard) Persist(f func(developmentCardModel *model.DevelopmentCard) error) error {
	developmentCardModel := new(model.DevelopmentCard)
	developmentCardModel.ID = dc.id
	if dc.game != nil {
		developmentCardModel.GameID = dc.game.id
	}
	developmentCardModel.DevelopmentType = dc.developmentType
	if dc.player != nil {
		developmentCardModel.PlayerID = &dc.player.id
	}

	if err := f(developmentCardModel); err != nil {
		return errors.Wrap(err, "datamodel.DevelopmentCard.Persist")
	}
	dc.isModified = false
	dc.isRemoved = false

	dc.id = developmentCardModel.ID
	dc.developmentType = developmentCardModel.DevelopmentType

	return nil
}
