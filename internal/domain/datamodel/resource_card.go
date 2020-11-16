package datamodel

import "gorm.io/gorm"

type ResourceCard struct {
	gorm.Model
	GameID   *uint
	PlayerID *uint
	Type     ResourceType
}

type ResourceType string

const (
	RT_LUMBER ResourceType = "LUMBER"
	RT_BRICK  ResourceType = "BRICK"
	RT_WOOL   ResourceType = "WOOL"
	RT_GRAIN  ResourceType = "GRAIN"
	RT_ORE    ResourceType = "ORE"
)
