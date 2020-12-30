package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
)

type SafeConstructionRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Constructions, error)
}

type ConstructionRepository interface {
	SafeConstructionRepository
	InsertOrUpdate(ctx context.Context, construction *model.Construction) error
}

func NewConstructionRepository(transactionMiddleware *middleware.TransactionMiddleware) ConstructionRepository {
	return &constructionRepository{
		transactionMiddleware: transactionMiddleware,
	}
}

type constructionRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
}

func (cr *constructionRepository) FindByGameId(ctx context.Context, gameId uint) (model.Constructions, error) {
	var constructions model.Constructions
	return constructions, cr.transactionMiddleware.Get(ctx).Find(&constructions, "game_id = ?", gameId).Error
}

func (cr *constructionRepository) InsertOrUpdate(ctx context.Context, construction *model.Construction) error {
	return cr.transactionMiddleware.Get(ctx).Save(construction).Error

}
