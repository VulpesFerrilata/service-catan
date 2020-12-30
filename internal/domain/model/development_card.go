package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type DevelopmentCard struct {
	datamodel.DevelopmentCard
	game *Game
}

func (dc *DevelopmentCard) SetGame(game *Game) {
	if game != nil {
		dc.GameID = &game.id
	}
	dc.game = game
}

func (dc *DevelopmentCard) GetPlayer() *Player {
	if dc.PlayerID == nil {
		return nil
	}

	return dc.game.players.Filter(func(player Player) bool {
		return player.id == *dc.PlayerID
	}).First()
}
