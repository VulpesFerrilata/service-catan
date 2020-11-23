package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/db"
)

type SafeAchievementRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Achievements, error)
}

type AchievementRepository interface {
	SafeAchievementRepository
	InsertOrUpdate(ctx context.Context, achievement *model.Achievement) error
}

func NewAchievementRepository(dbContext *db.DbContext) AchievementRepository {
	return &achievementRepository{
		dbContext: dbContext,
	}
}

type achievementRepository struct {
	dbContext *db.DbContext
}

func (ar *achievementRepository) FindByGameId(ctx context.Context, gameId uint) (model.Achievements, error) {
	var achievements model.Achievements
	return achievements, ar.dbContext.GetDB(ctx).Find(&achievements, "game_id = ?", gameId).Error
}

func (ar *achievementRepository) InsertOrUpdate(ctx context.Context, achievement *model.Achievement) error {
	return ar.dbContext.GetDB(ctx).Save(achievement).Error

}
