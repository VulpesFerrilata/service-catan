package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewResourceCard(game *Game, resourceType datamodel.ResourceType) *ResourceCard {
	resourceCard := new(ResourceCard)
	resourceCard.resourceCard = new(datamodel.ResourceCard)
	resourceCard.resourceCard.Type = resourceType
	resourceCard.SetGame(game)
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

func (rc *ResourceCard) GetGameId() uint {
	return rc.resourceCard.GameID
}

func (rc *ResourceCard) setGameId(gameId uint) {
	if rc.resourceCard.GameID != gameId {
		rc.resourceCard.GameID = gameId
		rc.isModified = true
	}
}

func (rc *ResourceCard) GetPlayerId() *uint {
	return rc.resourceCard.PlayerID
}

func (rc *ResourceCard) SetPlayerId(playerId *uint) {
	if rc.resourceCard.PlayerID != playerId {
		rc.resourceCard.PlayerID = playerId
		rc.isModified = true
	}
}

func (rc *ResourceCard) GetType() datamodel.ResourceType {
	return rc.resourceCard.Type
}

func (rc *ResourceCard) IsModified() bool {
	return rc.isModified
}

func (rc *ResourceCard) SetGame(game *Game) {
	rc.game = game
	game.resourceCards.append(rc)
}

func (rc *ResourceCard) GetPlayer() *Player {
	if rc.GetPlayerId() == nil {
		return nil
	}

	return rc.game.players.Filter(func(player *Player) bool {
		return player.GetId() == *rc.GetPlayerId()
	}).First()
}
