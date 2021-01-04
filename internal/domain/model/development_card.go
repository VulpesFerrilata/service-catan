package model

import "github.com/VulpesFerrilata/library/pkg/model"

type DevelopmentCard struct {
	model.Model
	ID     int `gorm:"primaryKey"`
	GameID int
	Type   DevelopmentType
	UserID *int
}

type DevelopmentType string

const (
	Knight        DevelopmentType = "Knight"
	Monopoly      DevelopmentType = "Monopoly"
	RoadBuilding  DevelopmentType = "RoadBuilding"
	YearOfPlenty  DevelopmentType = "YearOfPlenty"
	VictoryPoints DevelopmentType = "VictoryPoints"
)
