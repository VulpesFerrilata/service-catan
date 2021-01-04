package datamodel

import (
	"github.com/VulpesFerrilata/library/pkg/model"
)

type Achievement struct {
	model.Model
	GameID      int             `gorm:"primaryKey"`
	Type        AchievementType `gorm:"primaryKey"`
	PlayerID    *int
	BonusPoints int
}

type AchievementType string

const (
	LongestRoad AchievementType = "LongestRoad"
	LargestArmy AchievementType = "LargestArmy"
)
