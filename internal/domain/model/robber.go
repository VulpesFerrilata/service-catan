package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewRobber(terrains Terrains) *Robber {
	robber := new(Robber)
	robber.Status = datamodel.RS_IDLE

	desertTerrain := terrains.Filter(func(terrain *Terrain) bool {
		if terrain.Type == datamodel.TT_DESERT {
			return true
		}
		return false
	}).First()
	robber.TerrainID = &desertTerrain.ID

	return robber
}

type Robber struct {
	datamodel.Robber
	game *Game
}

func (r *Robber) SetGame(game *Game) {
	if game != nil {
		r.GameID = &game.id
	}
	r.game = game
}
