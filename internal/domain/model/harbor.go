package datamodel

import (
	"github.com/VulpesFerrilata/library/pkg/model"
)

type Harbor struct {
	model.Model
	GameID   int `gorm:"primaryKey"`
	Q        int `gorm:"primaryKey"`
	R        int `gorm:"primaryKey"`
	TerrainQ int
	TerrainR int
	Type     HarborType
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
