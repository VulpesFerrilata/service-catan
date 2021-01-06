package model

import "gorm.io/gorm"

type Robber struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	GameID    int
	Status    RobberStatus
	TerrainID int
}

type RobberStatus string

const (
	Idle    RobberStatus = "Idle"
	Moving  RobberStatus = "Moving"
	Robbing RobberStatus = "Robbing"
)
