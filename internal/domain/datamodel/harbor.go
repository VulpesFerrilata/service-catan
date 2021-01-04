package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/pkg/math"
)

type Harbor struct {
	datamodel.Harbor
	game *Game
}

func (h *Harbor) SetGame(game *Game) {
	if game != nil {
		h.GameID = &game.id
	}
	h.game = game
}

func (h *Harbor) GetTerrain() *Terrain {
	if h.TerrainID == nil {
		return nil
	}
	return h.game.terrains.Filter(func(terrain *Terrain) bool {
		return terrain.ID == *h.TerrainID
	}).First()
}

func (h *Harbor) GetIntersectRoad() *Road {
	terrain := h.GetTerrain()

	if terrain == nil {
		return nil
	}

	return h.game.roads.Filter(func(road *Road) bool {
		if h.Q == terrain.Q {
			return road.Q == h.Q && road.R == math.Max(h.R, terrain.R) && road.Location == datamodel.RL_TOP_LEFT
		} else if h.R == terrain.R {
			return road.Q == math.Max(h.Q, terrain.Q) && road.R == h.R && road.Location == datamodel.RL_MID_LEFT
		}
		return road.Q == math.Max(h.Q, terrain.Q) && road.R == math.Min(h.R, terrain.R) && road.Location == datamodel.RL_BOT_LEFT
	}).First()
}
