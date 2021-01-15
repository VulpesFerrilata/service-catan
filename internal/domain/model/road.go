package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Road struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	GameID   uuid.UUID `gorm:"type:uuid"`
	Q        int
	R        int
	Location RoadLocation
	PlayerID *uuid.UUID `gorm:"type:uuid"`
}

type RoadLocation string

const (
	TopLeft    RoadLocation = "TopLeft"
	MiddleLeft RoadLocation = "MiddleLeft"
	BottomLeft RoadLocation = "BottomLeft"
)
