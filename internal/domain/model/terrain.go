package model

import "gorm.io/gorm"

type Terrain struct {
	gorm.Model
	ID          int `gorm:"primaryKey"`
	GameID      int
	Q           int
	R           int
	Number      int
	TerrainType TerrainType
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
