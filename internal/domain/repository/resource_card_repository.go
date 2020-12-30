package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
)

type SafeResourceCardRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.ResourceCards, error)
}

type ResourceCardRepository interface {
	SafeResourceCardRepository
	InsertOrUpdate(ctx context.Context, resourceCard *model.ResourceCard) error
}

func NewResourceCardRepository(transactionMiddleware *middleware.TransactionMiddleware) ResourceCardRepository {
	return &resourceCardRepository{
		transactionMiddleware: transactionMiddleware,
	}
}

type resourceCardRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
}

func (rcr *resourceCardRepository) FindByGameId(ctx context.Context, gameId uint) (model.ResourceCards, error) {
	var resourceCards model.ResourceCards
	return resourceCards, rcr.transactionMiddleware.Get(ctx).Find(&resourceCards, "game_id = ?", gameId).Error
}

func (rcr *resourceCardRepository) InsertOrUpdate(ctx context.Context, resourceCard *model.ResourceCard) error {
	return rcr.transactionMiddleware.Get(ctx).Save(resourceCard).Error
}
