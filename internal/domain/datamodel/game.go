package datamodel

import "gorm.io/gorm"

type Game struct {
	*gorm.Model
	Host         uint
	PlayerInTurn uint
	Turn         int
	Status       GameStatus
}

type GameStatus string

const (
	GS_WAITED   = "WAITED"
	GS_STARTED  = "STARTED"
	GS_FINISHED = "FINISHED"
)
