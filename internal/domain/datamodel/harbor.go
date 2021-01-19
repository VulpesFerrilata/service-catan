package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/pkg/math"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewHarbor(q int, r int, harborType model.HarborType, terrain *Terrain) (*Harbor, error) {
	harbor := new(Harbor)
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewHarbor")
	}
	harbor.id = id
	harbor.q = q
	harbor.r = r
	harbor.harborType = harborType
	harbor.terrainID = terrain.id

	harbor.SetModelState(Added)

	return harbor, nil
}

func NewHarborFromHarborModel(harborModel *model.Harbor) *Harbor {
	harbor := new(Harbor)
	harbor.id = harborModel.ID
	harbor.q = harborModel.Q
	harbor.r = harborModel.R
	harbor.harborType = harborModel.HarborType
	harbor.terrainID = harborModel.TerrainID

	harbor.SetModelState(Unchanged)

	return harbor
}

type Harbor struct {
	base
	id         uuid.UUID
	q          int
	r          int
	harborType model.HarborType
	terrainID  uuid.UUID
	game       *Game
}

func (h Harbor) GetIntersectRoad() *Road {
	terrain := h.game.terrains.Filter(func(terrain *Terrain) bool {
		return terrain.id == h.terrainID
	}).First()

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

func (h *Harbor) ToModel() *model.Harbor {
	h.SetModelState(Unchanged)

	harborModel := new(model.Harbor)
	harborModel.ID = h.id
	harborModel.Q = h.q
	harborModel.R = h.r
	harborModel.TerrainID = h.terrainID
	harborModel.HarborType = h.harborType
	return harborModel
}
