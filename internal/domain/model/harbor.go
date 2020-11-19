package model

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/pkg/math"
)

func NewHarbor() *Harbor {
	harbor := new(Harbor)
	harbor.harbor = new(datamodel.Harbor)
	return harbor
}

type Harbor struct {
	harbor     *datamodel.Harbor
	game       *Game
	isModified bool
}

func (h *Harbor) GetHarbor() datamodel.Harbor {
	return *h.harbor
}

func (h *Harbor) GetId() uint {
	return h.harbor.ID
}

func (h *Harbor) GetGameId() *uint {
	return h.harbor.GameID
}

func (h *Harbor) setGame(game *Game) {
	if game != nil {
		h.harbor.GameID = &game.game.ID
		h.game = game
	}
}

func (h *Harbor) GetQ() int {
	return h.harbor.Q
}

func (h *Harbor) GetR() int {
	return h.harbor.R
}

func (h *Harbor) GetType() datamodel.HarborType {
	return h.harbor.Type
}

func (h *Harbor) GetTerrainId() *uint {
	return h.harbor.TerrainID
}

func (h *Harbor) GetTerrain() *Terrain {
	return h.game.terrains.Filter(func(terrain *Terrain) bool {
		if h.harbor.TerrainID == nil {
			return false
		}
		return terrain.GetId() == *h.GetTerrainId()
	}).First()
}

func (h *Harbor) SetTerrain(terrain *Terrain) {
	h.harbor.TerrainID = &terrain.terrain.ID
}

func (h *Harbor) IsModified() bool {
	return h.isModified
}

func (h *Harbor) GetIntersectRoad() *Road {
	terrain := h.GetTerrain()

	if terrain == nil {
		return nil
	}

	return h.game.roads.Filter(func(road *Road) bool {
		if h.GetQ() == terrain.GetQ() {
			return road.GetQ() == h.GetQ() && road.GetR() == math.Max(h.GetR(), terrain.GetR()) && road.GetLocation() == datamodel.RL_TOP_LEFT
		} else if h.GetR() == terrain.GetR() {
			return road.GetQ() == math.Max(h.GetQ(), terrain.GetQ()) && road.GetR() == h.GetR() && road.GetLocation() == datamodel.RL_MID_LEFT
		}
		return road.GetQ() == math.Max(h.GetQ(), terrain.GetQ()) && road.GetR() == math.Min(h.GetR(), terrain.GetR()) && road.GetLocation() == datamodel.RL_BOT_LEFT
	}).First()
}
