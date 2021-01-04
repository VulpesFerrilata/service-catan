package datamodel

import "gorm.io/gorm"

type Terrain struct {
	gorm.Model
	GameID int `gorm:"primaryKey"`
	Q      int `gorm:"primaryKey"`
	R      int `gorm:"primaryKey"`
	Number int
	Type   TerrainType
}

type TerrainType string

const (
	HillTerrain     TerrainType = "Hill"
	FieldTerrain    TerrainType = "Field"
	PastureTerrain  TerrainType = "Pasture"
	MountainTerrain TerrainType = "Mountain"
	ForestTerrain   TerrainType = "Forest"
	DesertTerrain   TerrainType = "Desert"
)
