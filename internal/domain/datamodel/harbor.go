package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/pkg/math"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewHarborFromHarborModel(harborModel *model.Harbor) *Harbor {
	harbor := new(Harbor)
	harbor.id = harborModel.ID
	harbor.q = harborModel.Q
	harbor.r = harborModel.R
	harbor.harborType = harborModel.HarborType
	harbor.isModified = false
	harbor.isRemoved = false
	return harbor
}

type Harbor struct {
	base
	id         uuid.UUID
	q          int
	r          int
	harborType model.HarborType
	game       *Game
	terrain    *Terrain
}

func (h *Harbor) GetIntersectRoad() *Road {
	return h.game.roads.Filter(func(road *Road) bool {
		if h.q == h.terrain.q {
			return road.q == h.q && road.r == math.Max(h.r, h.terrain.r) && road.location == model.TopLeft
		} else if h.r == h.terrain.r {
			return road.q == math.Max(h.q, h.terrain.q) && road.r == h.r && road.location == model.MiddleLeft
		}
		return road.q == math.Max(h.q, h.terrain.q) && road.r == math.Min(h.r, h.terrain.r) && road.location == model.BottomLeft
	}).First()
}

func (h *Harbor) Persist(f func(harborModel *model.Harbor) error) error {
	harborModel := new(model.Harbor)
	harborModel.ID = h.id
	harborModel.Q = h.q
	harborModel.R = h.r
	if h.terrain != nil {
		harborModel.TerrainID = h.terrain.id
	}
	harborModel.HarborType = h.harborType

	if err := f(harborModel); err != nil {
		return errors.Wrap(err, "datamodel.Harbor.Persist")
	}
	h.isModified = false
	h.isRemoved = false

	h.id = harborModel.ID
	h.q = harborModel.Q
	h.r = harborModel.R
	h.harborType = harborModel.HarborType

	return nil
}
