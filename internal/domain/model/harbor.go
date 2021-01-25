package model

import (
	"github.com/VulpesFerrilata/library/pkg/model"
	"github.com/google/uuid"
)

type Harbor struct {
	model.Model
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	GameID     uuid.UUID `gorm:"type:uuid"`
	Q          int
	R          int
	HarborType string
	TerrainID  uuid.UUID `gorm:"type:uuid"`
}
