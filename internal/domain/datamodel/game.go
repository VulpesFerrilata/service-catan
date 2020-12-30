package datamodel

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	PlayerInTurn uint
	Turn         int
	Status       GameStatus
}

type GameStatus string

const (
	Waiting  GameStatus = "waiting"
	Started  GameStatus = "started"
	Finished GameStatus = "finished"
)
