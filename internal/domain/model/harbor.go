package model

import (
	"github.com/VulpesFerrilata/library/pkg/model"
)

type Harbor struct {
	model.Model
	ID         int `gorm:"primaryKey"`
	GameID     int
	Q          int
	R          int
	TerrainQ   int
	TerrainR   int
	HarborType HarborType
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
