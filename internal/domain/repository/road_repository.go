package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/db"
)

type SafeRoadRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Roads, error)
}

type RoadRepository interface {
	SafeRoadRepository
	InsertOrUpdate(ctx context.Context, road *model.Road) error
}

func NewRoadRepository(dbContext *db.DbContext) RoadRepository {
	return &roadRepository{
		dbContext: dbContext,
	}
}

type roadRepository struct {
	dbContext *db.DbContext
}

func (rr *roadRepository) FindByGameId(ctx context.Context, gameId uint) (model.Roads, error) {
	var roads model.Roads
	return roads, rr.dbContext.GetDB(ctx).Find(&roads, "game_id = ?", gameId).Error
}

func (rr *roadRepository) InsertOrUpdate(ctx context.Context, road *model.Road) error {
	return rr.dbContext.GetDB(ctx).Save(road).Error
}
