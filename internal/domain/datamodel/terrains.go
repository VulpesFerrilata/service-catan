package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/pkg/errors"
)

func NewTerrainsFromTerrainModels(terrainModels []*model.Terrain) (Terrains, error) {
	terrains := make(Terrains, 0)

	for _, terrainModel := range terrainModels {
		terrain, err := NewTerrainFromTerrainModel(terrainModel)
		if err != nil {
			return nil, errors.Wrap(err, "datamodel.NewTerrainsFromTerrainModels")
		}
		terrains = append(terrains, terrain)
	}

	return terrains, nil
}

type Terrains []*Terrain

type TerrainFilterFunc func(terrain *Terrain) bool

func (t Terrains) Filter(f TerrainFilterFunc) Terrains {
	var terrains Terrains
	for _, terrain := range t {
		if f(terrain) {
			terrains = append(terrains, terrain)
		}
	}
	return terrains
}

func (t Terrains) First() *Terrain {
	if len(t) > 0 {
		return (t)[0]
	}
	return nil
}

func (t Terrains) Any(f TerrainFilterFunc) bool {
	for _, terrain := range t {
		if f(terrain) {
			return true
		}
	}
	return false
}
