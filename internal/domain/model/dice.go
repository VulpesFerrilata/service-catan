package model

import (
	"github.com/VulpesFerrilata/library/pkg/model"
	"github.com/google/uuid"
)

type Dice struct {
	model.Model
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	GameID uuid.UUID `gorm:"type:uuid"`
	Number int
}
