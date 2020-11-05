package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewResourceCard(game *Game, resourceType datamodel.ResourceType) *ResourceCard {
	resourceCard := new(ResourceCard)
	resourceCard.Type = resourceType
	resourceCard.SetGame(game)
	return resourceCard
}

type ResourceCard struct {
	*datamodel.ResourceCard
	game *Game
}

func (rc *ResourceCard) SetGame(game *Game) {
	rc.GameID = game.ID
	rc.game = game

	rc.game.resourceCards.append(rc)
}

func (rc *ResourceCard) GetPlayer() *Player {
	return rc.game.players.Filter(func(player *Player) bool {
		return player.ID == *rc.PlayerID
	}).First()
}
