package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
)

func NewAchievementFromAchievementModel(achievementModel *model.Achievement) *Achievement {
	achievement := new(Achievement)
	achievement.id = achievementModel.ID
	achievement.achievementType = achievementModel.AchievementType
	achievement.bonusPoints = achievementModel.BonusPoints
	achievement.isModified = false
	achievement.isRemoved = false
	return achievement
}

type Achievement struct {
	base
	id              uuid.UUID
	achievementType model.AchievementType
	bonusPoints     int
	game            *Game
	playerID        *uuid.UUID
}

func (a Achievement) ToModel() *model.Achievement {
	achievementModel := new(model.Achievement)
	achievementModel.ID = a.id
	if a.game != nil {
		achievementModel.GameID = a.game.id
	}
	achievementModel.AchievementType = a.achievementType
	achievementModel.BonusPoints = a.bonusPoints
	achievementModel.PlayerID = a.playerID
	return achievementModel
}
