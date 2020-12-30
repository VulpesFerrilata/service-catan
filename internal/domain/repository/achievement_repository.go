package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
)

type SafeAchievementRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Achievements, error)
}

type AchievementRepository interface {
	SafeAchievementRepository
	InsertOrUpdate(ctx context.Context, achievement *model.Achievement) error
}

func NewAchievementRepository(transactionMiddleware *middleware.TransactionMiddleware) AchievementRepository {
	return &achievementRepository{
		transactionMiddleware: transactionMiddleware,
	}
}

type achievementRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
}

func (ar *achievementRepository) FindByGameId(ctx context.Context, gameId uint) (model.Achievements, error) {
	var achievements model.Achievements
	return achievements, ar.transactionMiddleware.Get(ctx).Find(&achievements, "game_id = ?", gameId).Error
}

func (ar *achievementRepository) InsertOrUpdate(ctx context.Context, achievement *model.Achievement) error {
	return ar.transactionMiddleware.Get(ctx).Save(achievement).Error

}
