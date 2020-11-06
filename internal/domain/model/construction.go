package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewConstruction(game *Game) *Construction {
	construction := new(Construction)
	construction.Construction = new(datamodel.Construction)
	construction.IsUpgradedCastle = false
	construction.SetGame(game)
	return construction
}

type Construction struct {
	*datamodel.Construction
	game *Game
}

func (c *Construction) SetGame(game *Game) {
	c.game = game
	game.constructions.append(c)
}

func (c *Construction) GetPlayer() *Player {
	if c.PlayerID == nil {
		return nil
	}

	return c.game.players.Filter(func(player *Player) bool {
		return player.ID == *c.PlayerID
	}).First()
}
