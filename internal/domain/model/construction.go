package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type Construction struct {
	datamodel.Construction
	game *Game
}

func (c *Construction) SetGame(game *Game) {
	if game != nil {
		c.GameID = &game.ID
	}
	c.game = game
}

func (c *Construction) GetPlayer() *Player {
	if c.PlayerID == nil {
		return nil
	}

	return c.game.players.Filter(func(player *Player) bool {
		return player.ID == *c.PlayerID
	}).First()
}
