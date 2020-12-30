package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
)

type SafeHarborRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Harbors, error)
}

type HarborRepository interface {
	SafeHarborRepository
	InsertOrUpdate(ctx context.Context, harbor *model.Harbor) error
}

func NewHarborRepository(transactionMiddleware *middleware.TransactionMiddleware) HarborRepository {
	return &harborRepository{
		transactionMiddleware: transactionMiddleware,
	}
}

type harborRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
}

func (hr *harborRepository) FindByGameId(ctx context.Context, gameId uint) (model.Harbors, error) {
	var harbors model.Harbors
	return harbors, hr.transactionMiddleware.Get(ctx).Find(&harbors, "game_id = ?", gameId).Error
}

func (hr *harborRepository) InsertOrUpdate(ctx context.Context, harbor *model.Harbor) error {
	return hr.transactionMiddleware.Get(ctx).Save(harbor).Error
}
