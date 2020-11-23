package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/db"
)

type SafeConstructionRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Constructions, error)
}

type ConstructionRepository interface {
	SafeConstructionRepository
	InsertOrUpdate(ctx context.Context, construction *model.Construction) error
}

func NewConstructionRepository(dbContext *db.DbContext) ConstructionRepository {
	return &constructionRepository{
		dbContext: dbContext,
	}
}

type constructionRepository struct {
	dbContext *db.DbContext
}

func (cr *constructionRepository) FindByGameId(ctx context.Context, gameId uint) (model.Constructions, error) {
	var constructions model.Constructions
	return constructions, cr.dbContext.GetDB(ctx).Find(&constructions, "game_id = ?", gameId).Error
}

func (cr *constructionRepository) InsertOrUpdate(ctx context.Context, construction *model.Construction) error {
	return cr.dbContext.GetDB(ctx).Save(construction).Error

}
