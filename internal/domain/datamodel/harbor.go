package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/pkg/math"
	"github.com/pkg/errors"
)

func NewHarborFromHarborModel(harborModel *model.Harbor) *Harbor {
	harbor := new(Harbor)
	harbor.id = harborModel.ID
	harbor.q = harborModel.Q
	harbor.r = harborModel.R
	harbor.terrainQ = harborModel.TerrainQ
	harbor.terrainR = harborModel.TerrainR
	harbor.harborType = harborModel.HarborType
	harbor.isModified = false
	harbor.isRemoved = false
	return harbor
}

type Harbor struct {
	base
	id         int
	q          int
	r          int
	terrainQ   int
	terrainR   int
	harborType model.HarborType
	game       *Game
}

func (h Harbor) GetTerrain() *Terrain {
	return h.game.terrains.Filter(func(terrain *Terrain) bool {
		return terrain.q == h.terrainQ && terrain.r == h.terrainR
	}).First()
}

func (h *Harbor) GetIntersectroad() *Road {
	terrain := h.GetTerrain()

	if terrain == nil {
		return nil
	}

	return h.game.roads.Filter(func(road *Road) bool {
		if h.q == terrain.q {
			return road.q == h.q && road.r == math.Max(h.r, terrain.r) && road.location == model.TopLeft
		} else if h.r == terrain.r {
			return road.q == math.Max(h.q, terrain.q) && road.r == h.r && road.location == model.MiddleLeft
		}
		return road.q == math.Max(h.q, terrain.q) && road.r == math.Min(h.r, terrain.r) && road.location == model.BottomLeft
	}).First()
}

func (h *Harbor) Persist(f func(harborModel *model.Harbor) error) error {
	harborModel := new(model.Harbor)
	harborModel.ID = h.id
	harborModel.Q = h.q
	harborModel.R = h.r
	harborModel.TerrainQ = h.terrainQ
	harborModel.TerrainR = h.terrainR
	harborModel.HarborType = h.harborType

	if err := f(harborModel); err != nil {
		return errors.Wrap(err, "datamodel.Harbor.Persist")
	}
	h.isModified = false
	h.isRemoved = false

	h.id = harborModel.ID
	h.q = harborModel.Q
	h.r = harborModel.R
	h.terrainQ = harborModel.TerrainQ
	h.terrainR = harborModel.TerrainR
	h.harborType = harborModel.HarborType

	return nil
}
