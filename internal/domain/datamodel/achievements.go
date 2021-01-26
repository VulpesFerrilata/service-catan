package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/pkg/errors"
)

func NewAchievementsFromAchievementModels(achievementModels []*model.Achievement) (Achievements, error) {
	achievements := make(Achievements, 0)

	for _, achievementModel := range achievementModels {
		achievement, err := NewAchievementFromAchievementModel(achievementModel)
		if err != nil {
			return nil, errors.Wrap(err, "datamodel.NewAchievementsFromAchievementModels")
		}
		achievements = append(achievements, achievement)
	}

	return achievements, nil
}

type Achievements []*Achievement

type AchievementFilterFunc func(achievement *Achievement) bool

func (a Achievements) Filter(achievementFilterFunc AchievementFilterFunc) Achievements {
	var achievements Achievements
	for _, achievement := range a {
		if achievementFilterFunc(achievement) {
			achievements = append(achievements, achievement)
		}
	}
	return achievements
}
