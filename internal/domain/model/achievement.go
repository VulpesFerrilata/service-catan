package model

import (
	"github.com/VulpesFerrilata/library/pkg/model"
)

type Achievement struct {
	model.Model
	ID              int `gorm:"primaryKey"`
	GameID          int
	AchievementType AchievementType
	PlayerID        *int
	BonusPoints     int
}

type AchievementType string

const (
	LongestRoad AchievementType = "LongestRoad"
	LargestArmy AchievementType = "LargestArmy"
)
