package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewResourceCard(game *Game, resourceType datamodel.ResourceType) *ResourceCard {
	resourceCard := new(ResourceCard)
	resourceCard.Type = resourceType
	resourceCard.SetGame(game)
	resourceCard.SetPlayer(nil)
	return resourceCard
}

type ResourceCard struct {
	*datamodel.ResourceCard
	game   *Game
	player *Player
}

func (rc *ResourceCard) SetGame(game *Game) {
	rc.GameID = game.ID
	rc.game = game

	if rc.player == nil {
		rc.game.resourceCards.append(rc)
	}
}

func (rc *ResourceCard) SetPlayer(player *Player) {
	if rc.player != nil {
		rc.player.resourceCards.remove(rc)
	}
	if rc.game != nil {
		rc.game.resourceCards.remove(rc)
	}

	if player != nil {
		rc.PlayerID = player.ID
		rc.player = player
		player.resourceCards.append(rc)
	} else {
		rc.PlayerID = 0
		rc.player = nil
		rc.game.resourceCards.append(rc)
	}
}
