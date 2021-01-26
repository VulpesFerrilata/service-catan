package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewAchievement(achievementType AchievementType, bonusPoint int) (*Achievement, error) {
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

func NewAchievementFromAchievementModel(achievementModel *model.Achievement) (*Achievement, error) {
	achievement := new(Achievement)
	achievement.id = achievementModel.ID

	achievementType, err := NewAchievementType(achievementModel.AchievementType)
	if err != nil {
		return nil, errors.Wrap(err, "datamodel.NewAchievementFromAchievementModel")
	}
	achievement.achievementType = achievementType

	achievement.bonusPoints = achievementModel.BonusPoints
	achievement.playerID = achievementModel.PlayerID
	return achievement, nil
}

type Achievement struct {
	id              uuid.UUID
	achievementType AchievementType
	bonusPoints     int
	playerID        *uuid.UUID
	game            *Game
}

func (a Achievement) ToModel() *model.Achievement {
	achievementModel := new(model.Achievement)
	achievementModel.ID = a.id
	achievementModel.GameID = a.game.id
	achievementModel.AchievementType = a.achievementType.String()
	achievementModel.BonusPoints = a.bonusPoints
	achievementModel.PlayerID = a.playerID
	return achievementModel
}
