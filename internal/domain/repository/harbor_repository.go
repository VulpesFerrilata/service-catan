package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/db"
)

type SafeHarborRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Harbors, error)
}

type HarborRepository interface {
	SafeHarborRepository
	InsertOrUpdate(ctx context.Context, harbor *model.Harbor) error
}

func NewHarborRepository(dbContext *db.DbContext) HarborRepository {
	return &harborRepository{
		dbContext: dbContext,
	}
}

type harborRepository struct {
	dbContext *db.DbContext
}

func (hr *harborRepository) FindByGameId(ctx context.Context, gameId uint) (model.Harbors, error) {
	var harbors model.Harbors
	return harbors, hr.dbContext.GetDB(ctx).Find(&harbors, "game_id = ?", gameId).Error
}

func (hr *harborRepository) InsertOrUpdate(ctx context.Context, harbor *model.Harbor) error {
	return hr.dbContext.GetDB(ctx).Save(harbor).Error
}
