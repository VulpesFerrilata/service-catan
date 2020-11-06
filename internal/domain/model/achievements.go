package model

func NewAchievements(game *Game) Achievements {
	var achievements Achievements

	longestRoadAchievement := NewLongestRoadAchievement(game)
	achievements.append(longestRoadAchievement)

	largestArmyAchievement := NewLargestArmyAchievement(game)
	achievements.append(largestArmyAchievement)

	return achievements
}

type Achievements []*Achievement

func (a *Achievements) append(achievement *Achievement) {
	*a = append(*a, achievement)
}

func (a *Achievements) SetGame(game *Game) {
	for _, achievement := range *a {
		achievement.SetGame(game)
	}
}

type AchievementFilterFunc func(achievement *Achievement) bool

func (a *Achievements) Filter(achievementFilterFunc AchievementFilterFunc) Achievements {
	var achievements Achievements
	for _, achievement := range *a {
		if achievementFilterFunc(achievement) {
			achievements.append(achievement)
		}
	}
	return achievements
}
