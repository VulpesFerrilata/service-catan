package model

import "github.com/VulpesFerrilata/library/pkg/model"

type Construction struct {
	model.Model
	ID               int `gorm:"primaryKey"`
	GameID           int
	Q                int
	R                int
	Location         ConstructionLocation
	PlayerID         *int
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
