package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewResourceCard() *ResourceCard {
	resourceCard := new(ResourceCard)
	resourceCard.resourceCard = new(datamodel.ResourceCard)
	return resourceCard
}

type ResourceCard struct {
	resourceCard *datamodel.ResourceCard
	game         *Game
	isModified   bool
}

func (rc *ResourceCard) GetResourceCard() datamodel.ResourceCard {
	return *rc.resourceCard
}

func (rc *ResourceCard) GetId() uint {
	return rc.resourceCard.ID
}

func (rc *ResourceCard) GetGameId() *uint {
	return rc.resourceCard.GameID
}

func (rc *ResourceCard) setGame(game *Game) {
	if game != nil {
		rc.resourceCard.GameID = &game.game.ID
		rc.game = game
	}
}

func (rc *ResourceCard) GetPlayerId() *uint {
	return rc.resourceCard.PlayerID
}

func (rc *ResourceCard) GetType() datamodel.ResourceType {
	return rc.resourceCard.Type
}

func (rc *ResourceCard) IsModified() bool {
	return rc.isModified
}

func (rc *ResourceCard) GetPlayer() *Player {
	if rc.GetPlayerId() == nil {
		return nil
	}

	return rc.game.players.Filter(func(player *Player) bool {
		return player.GetId() == *rc.GetPlayerId()
	}).First()
}

func (rc *ResourceCard) SetPlayer(player *Player) {
	if player == nil {
		rc.resourceCard.PlayerID = nil
	} else {
		rc.resourceCard.PlayerID = &player.player.ID
	}
	rc.isModified = true
}
