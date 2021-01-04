package datamodel

import "gorm.io/gorm"

type Robber struct {
	gorm.Model
	GameID   int `gorm:"primaryKey"`
	TerrainQ int
	TerrainR int
	Status   RobberStatus
}

type RobberStatus string

const (
	Idle    RobberStatus = "Idle"
	Moving  RobberStatus = "Moving"
	Robbing RobberStatus = "Robbing"
)
