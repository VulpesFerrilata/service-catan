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
	HarborType HarborType
	TerrainID  uuid.UUID `gorm:"type:uuid"`
}

type HarborType string

const (
	GeneralHarbor HarborType = "General"
	LumberHarbor  HarborType = "Lumber"
	BrickHarbor   HarborType = "Brick"
	WoolHarbor    HarborType = "Wool"
	GrainHarbor   HarborType = "Grain"
	OreHarbor     HarborType = "Ore"
)
