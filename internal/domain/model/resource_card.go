package model

import (
	"github.com/VulpesFerrilata/library/pkg/model"
	"github.com/google/uuid"
)

type ResourceCard struct {
	model.Model
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey"`
	GameID       uuid.UUID  `gorm:"type:uuid"`
	PlayerID     *uuid.UUID `gorm:"type:uuid"`
	ResourceType ResourceType
}

type ResourceType string

const (
	LumberResource ResourceType = "Lumber"
	BrickResource  ResourceType = "Brick"
	WoolResource   ResourceType = "Wool"
	GrainResource  ResourceType = "Grain"
	OreResource    ResourceType = "Ore"
)
