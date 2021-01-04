package model

import "github.com/VulpesFerrilata/library/pkg/model"

type Construction struct {
	model.Model
	GameID           int                  `gorm:"primaryKey"`
	Q                int                  `gorm:"primaryKey"`
	R                int                  `gorm:"primaryKey"`
	Location         ConstructionLocation `gorm:"primaryKey"`
	UserID           *int
	IsUpgradedCastle bool
}

type ConstructionLocation string

const (
	Top    ConstructionLocation = "Top"
	Bottom ConstructionLocation = "Bottom"
)
