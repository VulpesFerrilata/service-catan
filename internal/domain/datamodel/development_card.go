package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/pkg/errors"
)

func NewDevelopmentCardFromDevelopmentCardModel(developmentCardModel *model.DevelopmentCard) *DevelopmentCard {
	developmentCard := new(DevelopmentCard)
	developmentCard.id = developmentCardModel.ID
	developmentCard.developmentType = developmentCardModel.DevelopmentType
	developmentCard.playerID = developmentCard.playerID
	developmentCard.isModified = false
	developmentCard.isRemoved = false
	return developmentCard
}

type DevelopmentCard struct {
	base
	id              int
	developmentType model.DevelopmentType
	playerID        *int
	game            *Game
}

func (dc DevelopmentCard) GetPlayer() *Player {
	if dc.playerID == nil {
		return nil
	}

	return dc.game.players.Filter(func(player *Player) bool {
		return player.id == *dc.playerID
	}).First()
}

func (dc *DevelopmentCard) Persist(f func(developmentCardModel *model.DevelopmentCard) error) error {
	developmentCardModel := new(model.DevelopmentCard)
	developmentCardModel.ID = dc.id
	if dc.game != nil {
		developmentCardModel.GameID = dc.game.id
	}
	developmentCardModel.DevelopmentType = dc.developmentType
	developmentCardModel.PlayerID = dc.playerID

	if err := f(developmentCardModel); err != nil {
		return errors.Wrap(err, "datamodel.DevelopmentCard.Persist")
	}
	dc.isModified = false
	dc.isRemoved = false

	dc.id = developmentCardModel.ID
	dc.developmentType = developmentCardModel.DevelopmentType
	dc.playerID = developmentCardModel.PlayerID

	return nil
}
