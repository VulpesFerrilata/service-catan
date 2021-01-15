package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Terrain struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	GameID      uuid.UUID `gorm:"type:uuid"`
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
