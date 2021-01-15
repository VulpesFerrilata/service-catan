package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewResourceCardFromResourceCardModel(resourceCardModel *model.ResourceCard) *ResourceCard {
	resourceCard := new(ResourceCard)
	resourceCard.id = resourceCardModel.ID
	resourceCard.playerID = resourceCardModel.PlayerID
	resourceCard.resourceType = resourceCardModel.ResourceType
	resourceCard.isModified = false
	resourceCard.isRemoved = false
	return resourceCard
}

type ResourceCard struct {
	base
	id           uuid.UUID
	resourceType model.ResourceType
	game         *Game
	player       *Player
}

func (rc ResourceCard) GetPlayer() *Player {
	if rc.playerID == nil {
		return nil
	}

	return rc.game.players.Filter(func(player *Player) bool {
		return player.id == *rc.playerID
	}).First()
}

func (rc *ResourceCard) Persist(f func(resourceCardModel *model.ResourceCard) error) error {
	resourceCardModel := new(model.ResourceCard)
	resourceCardModel.ID = rc.id
	resourceCardModel.PlayerID = rc.playerID
	resourceCardModel.ResourceType = rc.resourceType

	if err := f(resourceCardModel); err != nil {
		return errors.Wrap(err, "model.ResourceCard.Persist")
	}
	rc.isModified = false
	rc.isRemoved = false

	rc.id = resourceCardModel.ID
	rc.playerID = resourceCardModel.PlayerID
	rc.resourceType = resourceCardModel.ResourceType

	return nil
}
