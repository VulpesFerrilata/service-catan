package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type SafeAchievementRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Achievements, error)
}

type AchievementRepository interface {
	SafeAchievementRepository
	InsertOrUpdate(ctx context.Context, achievement *model.Achievement) error
}
