package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewConstruction(game *Game) *Construction {
	construction := new(Construction)
	construction.construction = new(datamodel.Construction)
	construction.construction.IsUpgradedCastle = false
	construction.SetGame(game)
	return construction
}

type Construction struct {
	construction *datamodel.Construction
	game         *Game
	isModified   bool
}

func (c *Construction) GetConstruction() datamodel.Construction {
	return *c.construction
}

func (c *Construction) GetId() uint {
	return c.construction.ID
}

func (c *Construction) GetGameId() uint {
	return c.construction.GameID
}

func (c *Construction) setGameId(gameId uint) {
	if c.construction.GameID != gameId {
		c.construction.GameID = gameId
		c.isModified = true
	}
}

func (c *Construction) GetPlayerId() *uint {
	return c.construction.PlayerID
}

func (c *Construction) SetPlayerId(playerId *uint) {
	if c.construction.PlayerID != playerId {
		c.construction.PlayerID = playerId
		c.isModified = true
	}
}

func (c *Construction) GetQ() int {
	return c.construction.Q
}

func (c *Construction) GetR() int {
	return c.construction.R
}

func (c *Construction) GetLocation() datamodel.ConstructionLocation {
	return c.construction.Location
}

func (c *Construction) IsUpgradedCastle() bool {
	return c.construction.IsUpgradedCastle
}

func (c *Construction) UpgradeCastle() {
	if !c.construction.IsUpgradedCastle {
		c.construction.IsUpgradedCastle = true
		c.isModified = true
	}
}

func (c *Construction) IsModified() bool {
	return c.isModified
}

func (c *Construction) SetGame(game *Game) {
	c.game = game
	game.constructions.append(c)
}

func (c *Construction) GetPlayer() *Player {
	if c.GetPlayerId() == nil {
		return nil
	}

	return c.game.players.Filter(func(player *Player) bool {
		return player.GetId() == *c.GetPlayerId()
	}).First()
}
