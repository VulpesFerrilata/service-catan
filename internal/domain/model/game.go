package model

import "github.com/VulpesFerrilata/library/pkg/model"

type Game struct {
	model.Model
	ID            int `gorm:"primaryKey"`
	PlayerInTurn  int
	Turn          int
	Status        GameStatus
	IsRolledDices bool
}

type GameStatus string

const (
	Waiting  GameStatus = "Waiting"
	Started  GameStatus = "Started"
	Finished GameStatus = "Finished"
)
