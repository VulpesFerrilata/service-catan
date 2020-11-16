package model

func NewAchievements() Achievements {
	var achievements Achievements

	longestRoadAchievement := NewLongestRoadAchievement()
	achievements.append(longestRoadAchievement)

	largestArmyAchievement := NewLargestArmyAchievement()
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
