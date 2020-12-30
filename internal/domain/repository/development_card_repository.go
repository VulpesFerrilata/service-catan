package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
)

type SafeDevelopmentCardRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.DevelopmentCards, error)
}

type DevelopmentCardRepository interface {
	SafeDevelopmentCardRepository
	InsertOrUpdate(ctx context.Context, developmentCard *model.DevelopmentCard) error
}

func NewDevelopmentCardRepository(transactionMiddleware *middleware.TransactionMiddleware) DevelopmentCardRepository {
	return &developmentCardRepository{
		transactionMiddleware: transactionMiddleware,
	}
}

type developmentCardRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
}

func (dcr *developmentCardRepository) FindByGameId(ctx context.Context, gameId uint) (model.DevelopmentCards, error) {
	var developmentCards model.DevelopmentCards
	return developmentCards, dcr.transactionMiddleware.Get(ctx).Find(&developmentCards, "game_id = ?", gameId).Error
}

func (dcr *developmentCardRepository) InsertOrUpdate(ctx context.Context, developmentCard *model.DevelopmentCard) error {
	return dcr.transactionMiddleware.Get(ctx).Save(developmentCard).Error
}
