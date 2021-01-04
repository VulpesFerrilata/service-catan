package datamodel

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type Achievement struct {
	datamodel.Achievement
	game *Game
}

func (a *Achievement) SetGame(game *Game) {
	if game != nil {
		a.GameID = &game.id
	}
	a.game = game
}

func (a *Achievement) GetPlayer() *Player {
	if a.PlayerID == nil {
		return nil
	}

	return a.game.players.Filter(func(player Player) bool {
		return player.id == *a.PlayerID
	}).First()
}
