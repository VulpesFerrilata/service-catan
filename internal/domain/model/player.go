package model

import (
	"github.com/VulpesFerrilata/library/pkg/model"
	"github.com/google/uuid"
)

type Player struct {
	model.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	GameID    uuid.UUID `gorm:"type:uuid"`
	UserID    uuid.UUID `gorm:"type:uuid"`
	Color     string
	TurnOrder int
	IsLeft    bool
}
