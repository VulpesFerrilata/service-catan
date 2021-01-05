package model

import "github.com/VulpesFerrilata/library/pkg/model"

type ResourceCard struct {
	model.Model
	ID           int `gorm:"primaryKey"`
	GameID       int
	PlayerID     *int
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
