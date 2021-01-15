package model

import (
	"github.com/VulpesFerrilata/library/pkg/model"
	"github.com/google/uuid"
)

type Construction struct {
	model.Model
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	GameID           uuid.UUID `gorm:"type:uuid"`
	Q                int
	R                int
	Location         ConstructionLocation
	PlayerID         *uuid.UUID `gorm:"type:uuid"`
	ConstructionType ConstructionType
}

type ConstructionLocation string

const (
	Top    ConstructionLocation = "Top"
	Bottom ConstructionLocation = "Bottom"
)

type ConstructionType string

const (
	Land       ConstructionType = "Land"
	Settlement ConstructionType = "Settlement"
	City       ConstructionType = "City"
)
