package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/db"
)

type SafeDevelopmentCardRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.DevelopmentCards, error)
}

type DevelopmentCardRepository interface {
	SafeDevelopmentCardRepository
	InsertOrUpdate(ctx context.Context, developmentCard *model.DevelopmentCard) error
}

func NewDevelopmentCardRepository(dbContext *db.DbContext) DevelopmentCardRepository {
	return &developmentCardRepository{
		dbContext: dbContext,
	}
}

type developmentCardRepository struct {
	dbContext *db.DbContext
}

func (dcr *developmentCardRepository) FindByGameId(ctx context.Context, gameId uint) (model.DevelopmentCards, error) {
	var developmentCards model.DevelopmentCards
	return developmentCards, dcr.dbContext.GetDB(ctx).Find(&developmentCards, "game_id = ?", gameId).Error
}

func (dcr *developmentCardRepository) InsertOrUpdate(ctx context.Context, developmentCard *model.DevelopmentCard) error {
	return dcr.dbContext.GetDB(ctx).Save(developmentCard).Error
}
