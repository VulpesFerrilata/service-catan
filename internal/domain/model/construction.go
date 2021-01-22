package model

import (
	"github.com/VulpesFerrilata/library/pkg/model"
	"github.com/google/uuid"
)

type Construction struct {
	model.Model
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	GameID           uuid.UUID `gorm:"type:uuid"`
	Q                int
	R                int
	Location         string
	PlayerID         *uuid.UUID `gorm:"type:uuid"`
	ConstructionType string
}
