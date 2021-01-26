package model

import (
	"github.com/VulpesFerrilata/library/pkg/model"
	"github.com/google/uuid"
)

type Achievement struct {
	model.Model
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	GameID          uuid.UUID `gorm:"type:uuid"`
	AchievementType string
	PlayerID        *uuid.UUID `gorm:"type:uuid"`
	BonusPoints     int
}
