package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/pkg/math"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewHarbor(hex *Hex, harborType HarborType, terrain *Terrain) (*Harbor, error) {
	harbor := new(Harbor)

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewHarbor")
	}
	harbor.id = id

	harbor.hex = hex
	harbor.harborType = harborType
	harbor.terrainID = terrain.id

	return harbor, nil
}

func NewHarborFromHarborModel(harborModel *model.Harbor) (*Harbor, error) {
	harbor := new(Harbor)
	harbor.id = harborModel.ID

	hex := NewHex(harborModel.Q, harborModel.R)
	harbor.hex = hex

	harborType, err := NewHarborType(harborModel.HarborType)
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewHarborFromHarborModel")
	}
	harbor.harborType = harborType

	harbor.terrainID = harborModel.TerrainID
	return harbor, nil
}

type Harbor struct {
	id         uuid.UUID
	hex        *Hex
	harborType HarborType
	terrainID  uuid.UUID
	game       *Game
}

func (h Harbor) GetHex() *Hex {
	return h.hex
}

func (h Harbor) GetIntersectRoad() *Road {
	terrain := h.game.terrains.Filter(func(terrain *Terrain) bool {
		return terrain.id == h.terrainID
	}).First()

	return h.game.roads.Filter(func(road *Road) bool {
		if h.q == terrain.q {
			return road.q == h.q && road.r == math.Max(h.r, terrain.r) && road.location == model.TopLeft
		} else if h.r == terrain.r {
			return road.q == math.Max(h.q, terrain.q) && road.r == h.r && road.location == model.MiddleLeft
		}
		return road.q == math.Max(h.q, terrain.q) && road.r == math.Min(h.r, terrain.r) && road.location == model.BottomLeft
	}).First()
}

func (h Harbor) ToModel() *model.Harbor {
	harborModel := new(model.Harbor)
	harborModel.ID = h.id
	harborModel.Q = h.q
	harborModel.R = h.r
	harborModel.TerrainID = h.terrainID
	harborModel.HarborType = h.harborType
	return harborModel
}
