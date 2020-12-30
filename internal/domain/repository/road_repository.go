package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
)

type SafeRoadRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Roads, error)
}

type RoadRepository interface {
	SafeRoadRepository
	InsertOrUpdate(ctx context.Context, road *model.Road) error
}

func NewRoadRepository(transactionMiddleware *middleware.TransactionMiddleware) RoadRepository {
	return &roadRepository{
		transactionMiddleware: transactionMiddleware,
	}
}

type roadRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
}

func (rr *roadRepository) FindByGameId(ctx context.Context, gameId uint) (model.Roads, error) {
	var roads model.Roads
	return roads, rr.transactionMiddleware.Get(ctx).Find(&roads, "game_id = ?", gameId).Error
}

func (rr *roadRepository) InsertOrUpdate(ctx context.Context, road *model.Road) error {
	return rr.transactionMiddleware.Get(ctx).Save(road).Error
}
