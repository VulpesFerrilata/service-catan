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
	TerrainType string
}
