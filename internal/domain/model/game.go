package model

import (
	"github.com/VulpesFerrilata/library/pkg/model"
	"github.com/google/uuid"
)

type Game struct {
	model.Model
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey"`
	playerInTurn  *uuid.UUID `gorm:"type:uuid"`
	CurrentTurn   int
	Status        GameStatus
	IsRolledDices bool
}

type GameStatus string

const (
	Waiting  GameStatus = "Waiting"
	Started  GameStatus = "Started"
	Finished GameStatus = "Finished"
)
