package model

import (
	"github.com/VulpesFerrilata/library/pkg/model"
	"github.com/google/uuid"
)

type Game struct {
	model.Model
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey"`
	PlayerInTurn  *uuid.UUID `gorm:"type:uuid"`
	Turn          int
	Status        string
	IsRolledDices bool
}
