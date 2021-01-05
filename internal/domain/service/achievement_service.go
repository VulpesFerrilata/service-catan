package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type AchievementService interface {
	GetAchievementRepository() repository.SafeAchievementRepository
	Save(ctx context.Context, achievement *datamodel.Achievement) error
}

func NewAchievementService(achievementRepository repository.AchievementRepository) AchievementService {
	return &achievementService{
		achievementRepository: achievementRepository,
	}
}

type achievementService struct {
	achievementRepository repository.AchievementRepository
}

func (as achievementService) GetAchievementRepository() repository.SafeAchievementRepository {
	return as.achievementRepository
}

func (as achievementService) Save(ctx context.Context, achievement *datamodel.Achievement) error {
	if achievement.IsRemoved() {
		return nil
	}
	if achievement.IsModified() {
		err := as.achievementRepository.InsertOrUpdate(ctx, achievement)
		return errors.Wrap(err, "service.AchievementService.Save")
	}
	return nil
}
