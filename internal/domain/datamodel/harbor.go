package datamodel

import (
	"fmt"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
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

func NewHarborFromModel(harborModel *model.Harbor, hex *Hex) (*Harbor, error) {
	harbor := new(Harbor)
	harbor.id = harborModel.ID

	harborType, err := NewHarborType(harborModel.HarborType)
	if err != nil {
		return nil, fmt.Errorf("harbor type is invalid: %s", harborModel.HarborType)
	}
	harbor.harborType = harborType

	harbor.terrainID = harborModel.TerrainID
	harbor.hex = hex
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

	hexEdge := h.hex.GetPossibleIntersectHexEdge(terrain.GetHex())

	return h.game.roads.Filter(func(road *Road) bool {
		if road.GetHexEdge().Equals(hexEdge) {
			return true
		}
		return false
	}).First()
}

func (h Harbor) ToModel() *model.Harbor {
	harborModel := new(model.Harbor)
	harborModel.ID = h.id

	if h.hex != nil {
		harborModel.HexID = h.hex.id
	}

	harborModel.TerrainID = h.terrainID
	harborModel.HarborType = h.harborType.String()
	return harborModel
}
