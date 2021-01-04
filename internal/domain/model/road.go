package datamodel

import "gorm.io/gorm"

type Road struct {
	gorm.Model
	GameID   int          `gorm:"primaryKey"`
	Q        int          `gorm:"primaryKey"`
	R        int          `gorm:"primaryKey"`
	Location RoadLocation `gorm:"primaryKey"`
	PlayerID *int
}

type RoadLocation string

const (
	TopLeft    RoadLocation = "TopLeft"
	MiddleLeft RoadLocation = "MiddleLeft"
	BottomLeft RoadLocation = "BottomLeft"
)
