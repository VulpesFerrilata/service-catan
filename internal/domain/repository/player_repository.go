package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/db"
)

type SafePlayerRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Players, error)
}

type PlayerRepository interface {
	SafePlayerRepository
	InsertOrUpdate(ctx context.Context, player *model.Player) error
	Delete(ctx context.Context, player *model.Player) error
}

func NewPlayerRepository(dbContext *db.DbContext) PlayerRepository {
	return &playerRepository{
		dbContext: dbContext,
	}
}

type playerRepository struct {
	dbContext *db.DbContext
}

func (pr *playerRepository) FindByGameId(ctx context.Context, gameId uint) (model.Players, error) {
	var players model.Players
	return players, pr.dbContext.GetDB(ctx).Find(&players, "game_id = ?", gameId).Error
}

func (pr *playerRepository) InsertOrUpdate(ctx context.Context, player *model.Player) error {
	return pr.dbContext.GetDB(ctx).Save(player).Error
}

func (pr *playerRepository) Delete(ctx context.Context, player *model.Player) error {
	return pr.dbContext.GetDB(ctx).Delete(player).Error
}
