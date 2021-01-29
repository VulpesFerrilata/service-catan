package datamodel

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
