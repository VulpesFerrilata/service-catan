package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewConstruction() *Construction {
	construction := new(Construction)
	construction.construction = new(datamodel.Construction)
	construction.construction.IsUpgradedCastle = false
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

func (c *Construction) GetGameId() *uint {
	return c.construction.GameID
}

func (c *Construction) setGame(game *Game) {
	if game != nil {
		c.construction.GameID = &game.game.ID
		c.game = game
	}
}

func (c *Construction) GetPlayerId() *uint {
	return c.construction.PlayerID
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

func (c *Construction) GetPlayer() *Player {
	if c.GetPlayerId() == nil {
		return nil
	}

	return c.game.players.Filter(func(player *Player) bool {
		return player.GetId() == *c.GetPlayerId()
	}).First()
}

func (c *Construction) SetPlayer(player *Player) {
	if player == nil {
		c.construction.PlayerID = nil
	} else {
		c.construction.PlayerID = &player.player.ID
	}
	c.isModified = true
}
