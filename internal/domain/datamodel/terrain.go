package datamodel

import "gorm.io/gorm"

type Terrain struct {
	gorm.Model
	GameID uint
	Q      int
	R      int
	Number int
	Type   TerrainType
}

type TerrainType string

const (
	TT_HILL     TerrainType = "HILL"
	TT_FIELD    TerrainType = "FIELD"
	TT_PASTURE  TerrainType = "PASTURE"
	TT_MOUNTAIN TerrainType = "MOUNTAIN"
	TT_FOREST   TerrainType = "FOREST"
	TT_DESERT   TerrainType = "DESERT"
)
