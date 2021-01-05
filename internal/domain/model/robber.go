package model

import "gorm.io/gorm"

type Robber struct {
	gorm.Model
	ID       int `gorm:"primaryKey"`
	GameID   int
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
