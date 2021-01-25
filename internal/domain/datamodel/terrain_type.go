package datamodel

import "fmt"

func NewTerrainType(value string) (TerrainType, error) {
	TerrainType := TerrainType(value)
	if _, ok := terrainTypes[TerrainType]; ok {
		return TerrainType, nil
	}
	return TerrainType, fmt.Errorf("terrain type is invalid: %s", value)
}

type TerrainType string

func (t TerrainType) String() string {
	return string(t)
}

const (
	HillTerrain     TerrainType = "Hill"
	FieldTerrain    TerrainType = "Field"
	PastureTerrain  TerrainType = "Pasture"
	MountainTerrain TerrainType = "Mountain"
	ForestTerrain   TerrainType = "Forest"
	DesertTerrain   TerrainType = "Desert"
)

var terrainTypes = map[TerrainType]struct{}{
	HillTerrain:     struct{}{},
	FieldTerrain:    struct{}{},
	PastureTerrain:  struct{}{},
	MountainTerrain: struct{}{},
	ForestTerrain:   struct{}{},
	DesertTerrain:   struct{}{},
}
