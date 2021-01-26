package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type AchievementService interface {
	GetAchievementRepository() repository.AchievementRepository
	InitAchievements() (datamodel.Achievements, error)
}

func NewAchievementService(achievementRepository repository.AchievementRepository) AchievementService {
	return &achievementService{
		achievementRepository: achievementRepository,
	}
}

type achievementService struct {
	achievementRepository repository.AchievementRepository
}

func (a achievementService) GetAchievementRepository() repository.AchievementRepository {
	return a.achievementRepository
}

func (a achievementService) InitAchievements() (datamodel.Achievements, error) {
	achievements := make(datamodel.Achievements, 0)

	longestRoadAchievement, err := datamodel.NewAchievement(datamodel.LongestRoad, 2)
	if err != nil {
		return nil, errors.Wrap(err, "service.AchievementService.InitAchievements")
	}
	achievements = append(achievements, longestRoadAchievement)

	largestArmyAchievement, err := datamodel.NewAchievement(datamodel.LargestArmy, 2)
	if err != nil {
		return nil, errors.Wrap(err, "service.AchievementService.InitAchievements")
	}
	achievements = append(achievements, largestArmyAchievement)

	return achievements, nil
}
