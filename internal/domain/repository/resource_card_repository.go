package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/db"
)

type SafeResourceCardRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.ResourceCards, error)
}

type ResourceCardRepository interface {
	SafeResourceCardRepository
	InsertOrUpdate(ctx context.Context, resourceCard *model.ResourceCard) error
}

func NewResourceCardRepository(dbContext *db.DbContext) ResourceCardRepository {
	return &resourceCardRepository{
		dbContext: dbContext,
	}
}

type resourceCardRepository struct {
	dbContext *db.DbContext
}

func (rcr *resourceCardRepository) FindByGameId(ctx context.Context, gameId uint) (model.ResourceCards, error) {
	var resourceCards model.ResourceCards
	return resourceCards, rcr.dbContext.GetDB(ctx).Find(&resourceCards, "game_id = ?", gameId).Error
}

func (rcr *resourceCardRepository) InsertOrUpdate(ctx context.Context, resourceCard *model.ResourceCard) error {
	return rcr.dbContext.GetDB(ctx).Save(resourceCard).Error
}
