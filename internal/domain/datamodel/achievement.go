package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewAchievement(achievementType model.AchievementType, bonusPoint int) (*Achievement, error) {
	achievement := new(Achievement)

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewAchievement")
	}
	achievement.id = id

	achievement.achievementType = achievementType
	achievement.bonusPoints = bonusPoint
	achievement.playerID = nil
	return achievement, nil
}

func NewAchievementFromAchievementModel(achievementModel *model.Achievement) *Achievement {
	achievement := new(Achievement)
	achievement.id = achievementModel.ID
	achievement.achievementType = achievementModel.AchievementType
	achievement.bonusPoints = achievementModel.BonusPoints
	achievement.playerID = achievementModel.PlayerID
	return achievement
}

type Achievement struct {
	id              uuid.UUID
	achievementType model.AchievementType
	bonusPoints     int
	playerID        *uuid.UUID
	game            *Game
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
