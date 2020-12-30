package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type ResourceCard struct {
	datamodel.ResourceCard
	game *Game
}

func (rc *ResourceCard) SetGame(game *Game) {
	if game != nil {
		rc.GameID = &game.id
	}
	rc.game = game
}

func (rc *ResourceCard) GetPlayer() *Player {
	if rc.PlayerID == nil {
		return nil
	}

	return rc.game.players.Filter(func(player Player) bool {
		return player.id == *rc.PlayerID
	}).First()
}
