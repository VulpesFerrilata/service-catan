package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
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
