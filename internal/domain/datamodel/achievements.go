package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

func NewAchievementsFromAchievementModels(achievementModels []*model.Achievement) Achievements {
	achievements := make(Achievements, 0)

	for _, achievementModel := range achievementModels {
		achievement := NewAchievementFromAchievementModel(achievementModel)
		achievements = append(achievements, achievement)
	}

	return achievements
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
