package model

import (
	"github.com/VulpesFerrilata/library/pkg/model"
	"github.com/google/uuid"
)

type Harbor struct {
	model.Model
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	GameID     uuid.UUID `gorm:"type:uuid"`
	HarborType string
	HexID      uuid.UUID `gorm:"type:uuid"`
	TerrainID  uuid.UUID `gorm:"type:uuid"`
}
