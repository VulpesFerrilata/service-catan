package model

import (
	"github.com/VulpesFerrilata/library/pkg/model"
	"github.com/google/uuid"
)

type DevelopmentCard struct {
	model.Model
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	GameID          uuid.UUID `gorm:"type:uuid"`
	DevelopmentType DevelopmentType
	PlayerID        *uuid.UUID `gorm:"type:uuid"`
}

type DevelopmentType string

const (
	Knight        DevelopmentType = "Knight"
	Monopoly      DevelopmentType = "Monopoly"
	RoadBuilding  DevelopmentType = "RoadBuilding"
	YearOfPlenty  DevelopmentType = "YearOfPlenty"
	VictoryPoints DevelopmentType = "VictoryPoints"
)
