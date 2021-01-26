package datamodel

import "fmt"

func NewAchievementType(value string) (AchievementType, error) {
	achievementType := AchievementType(value)

	if _, ok := achievementTypes[achievementType]; ok {
		return achievementType, nil
	}

	return achievementType, fmt.Errorf("achievement type is invalid: %s", value)
}

type AchievementType string

func (a AchievementType) String() string {
	return string(a)
}

const (
	LongestRoad AchievementType = "LongestRoad"
	LargestArmy AchievementType = "LargestArmy"
)

var achievementTypes = map[AchievementType]struct{}{
	LongestRoad: {},
	LargestArmy: {},
}
