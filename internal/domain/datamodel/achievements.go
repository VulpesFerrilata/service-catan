package datamodel

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
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

func NewAchievements() Achievements {
	var achievements Achievements

	longestRoadAchievement := new(Achievement)
	longestRoadAchievement.Type = datamodel.AT_LONGEST_ROAD
	longestRoadAchievement.BonusPoints = 2
	achievements.append(longestRoadAchievement)

	largestArmyAchievement := new(Achievement)
	largestArmyAchievement.Type = datamodel.AT_LARGEST_ARMY
	largestArmyAchievement.BonusPoints = 2
	achievements.append(largestArmyAchievement)

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
