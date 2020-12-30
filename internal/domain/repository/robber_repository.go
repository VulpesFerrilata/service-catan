package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/app_error"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"gorm.io/gorm"
)

type SafeRobberRepository interface {
	GetByGameId(ctx context.Context, gameId uint) (*model.Robber, error)
}

type RobberRepository interface {
	SafeRobberRepository
	InsertOrUpdate(ctx context.Context, robber *model.Robber) error
}

func NewRobberRepository(transactionMiddleware *middleware.TransactionMiddleware) RobberRepository {
	return &robberRepository{
		transactionMiddleware: transactionMiddleware,
	}
}

type robberRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
}

func (rr *robberRepository) GetByGameId(ctx context.Context, gameId uint) (*model.Robber, error) {
	robber := new(model.Robber)
	err := rr.transactionMiddleware.Get(ctx).First(&robber, "game_id = ?", gameId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, app_error.NewNotFoundError("game")
	}
	return robber, err
}

func (rr *robberRepository) InsertOrUpdate(ctx context.Context, robber *model.Robber) error {
	return rr.transactionMiddleware.Get(ctx).Save(robber).Error
}
