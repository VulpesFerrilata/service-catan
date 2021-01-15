package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Robber struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	GameID    uuid.UUID `gorm:"type:uuid"`
	Status    RobberStatus
	TerrainID uuid.UUID `gorm:"type:uuid"`
}

type RobberStatus string

const (
	Idle    RobberStatus = "Idle"
	Moving  RobberStatus = "Moving"
	Robbing RobberStatus = "Robbing"
)
