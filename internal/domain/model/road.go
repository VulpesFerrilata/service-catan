package model

import "gorm.io/gorm"

type Road struct {
	gorm.Model
	ID       int `gorm:"primaryKey"`
	GameID   int
	Q        int
	R        int
	Location RoadLocation
	PlayerID *int
}

type RoadLocation string

const (
	TopLeft    RoadLocation = "TopLeft"
	MiddleLeft RoadLocation = "MiddleLeft"
	BottomLeft RoadLocation = "BottomLeft"
)
