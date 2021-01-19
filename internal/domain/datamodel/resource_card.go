package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewResourceCard(resourceType model.ResourceType) (*ResourceCard, error) {
	resourceCard := new(ResourceCard)
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewResourceCard")
	}
	resourceCard.id = id
	resourceCard.resourceType = resourceType
	resourceCard.playerID = nil

	resourceCard.SetModelState(Added)

	return resourceCard, nil
}

func NewResourceCardFromResourceCardModel(resourceCardModel *model.ResourceCard) *ResourceCard {
	resourceCard := new(ResourceCard)
	resourceCard.id = resourceCardModel.ID
	resourceCard.playerID = resourceCardModel.PlayerID
	resourceCard.resourceType = resourceCardModel.ResourceType

	resourceCard.SetModelState(Unchanged)

	return resourceCard
}

type ResourceCard struct {
	base
	id           uuid.UUID
	resourceType model.ResourceType
	playerID     *uuid.UUID
	game         *Game
}

func (rc ResourceCard) GetPlayer() *Player {
	if rc.playerID == nil {
		return nil
	}

	return rc.game.players.Filter(func(player *Player) bool {
		return player.id == *rc.playerID
	}).First()
}

func (rc *ResourceCard) ToModel() *model.ResourceCard {
	rc.SetModelState(Unchanged)

	resourceCardModel := new(model.ResourceCard)
	resourceCardModel.ID = rc.id
	resourceCardModel.PlayerID = rc.playerID
	resourceCardModel.ResourceType = rc.resourceType
	return resourceCardModel
}
