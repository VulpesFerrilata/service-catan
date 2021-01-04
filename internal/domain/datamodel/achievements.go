package datamodel

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

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

func (a *Achievements) append(achievement *Achievement) {
	*a = append(*a, achievement)
}

type AchievementFilterFunc func(achievement *Achievement) bool

func (a Achievements) Filter(achievementFilterFunc AchievementFilterFunc) Achievements {
	var achievements Achievements
	for _, achievement := range a {
		if achievementFilterFunc(achievement) {
			achievements.append(achievement)
		}
	}
	return achievements
}
