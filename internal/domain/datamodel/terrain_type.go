package datamodel

import "fmt"

func NewTerrainType(value string) (terrainType, error) {
	terrainType := terrainType(value)
	if _, ok := terrainTypes[terrainType]; ok {
		return terrainType, nil
	}
	return terrainType, fmt.Errorf("terrain type is invalid: %s", value)
}

type terrainType string

func (t terrainType) String() string {
	return string(t)
}

const (
	HillTerrain     terrainType = "Hill"
	FieldTerrain    terrainType = "Field"
	PastureTerrain  terrainType = "Pasture"
	MountainTerrain terrainType = "Mountain"
	ForestTerrain   terrainType = "Forest"
	DesertTerrain   terrainType = "Desert"
)

var terrainTypes = map[terrainType]struct{}{
	HillTerrain:     struct{}{},
	FieldTerrain:    struct{}{},
	PastureTerrain:  struct{}{},
	MountainTerrain: struct{}{},
	ForestTerrain:   struct{}{},
	DesertTerrain:   struct{}{},
}
