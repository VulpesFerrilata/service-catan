package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type AchievementService interface {
	GetAchievementRepository() repository.AchievementRepository
}

func NewAchievementService(achievementRepository repository.AchievementRepository) AchievementService {
	return &achievementService{
		achievementRepository: achievementRepository,
	}
}

type achievementService struct {
	achievementRepository repository.AchievementRepository
}

func (as achievementService) GetAchievementRepository() repository.AchievementRepository {
	return as.achievementRepository
}

func (as achievementService) InitAchievements() (datamodel.Achievements, error) {
	achievements := make(datamodel.Achievements, 0)

	longestRoadAchievement, err := datamodel.NewAchievement(model.LongestRoad, 2)
	if err != nil {
		return nil, errors.Wrap(err, "service.AchievementService.InitAchievements")
	}
	achievements = append(achievements, longestRoadAchievement)

	largestArmyAchievement, err := datamodel.NewAchievement(model.LargestArmy, 2)
	if err != nil {
		return nil, errors.Wrap(err, "service.AchievementService.InitAchievements")
	}
	achievements = append(achievements, largestArmyAchievement)

	return achievements, nil
}
