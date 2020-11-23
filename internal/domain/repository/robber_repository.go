package repository

import (
	"context"
	"errors"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/db"
	server_errors "github.com/VulpesFerrilata/library/pkg/errors"
	"gorm.io/gorm"
)

type SafeRobberRepository interface {
	GetByGameId(ctx context.Context, gameId uint) (*model.Robber, error)
}

type RobberRepository interface {
	SafeRobberRepository
	InsertOrUpdate(ctx context.Context, robber *model.Robber) error
}

func NewRobberRepository(dbContext *db.DbContext) RobberRepository {
	return &robberRepository{
		dbContext: dbContext,
	}
}

type robberRepository struct {
	dbContext *db.DbContext
}

func (rr *robberRepository) GetByGameId(ctx context.Context, gameId uint) (*model.Robber, error) {
	robber := new(model.Robber)
	err := rr.dbContext.GetDB(ctx).First(&robber, "game_id = ?", gameId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, server_errors.NewNotFoundError("game")
	}
	return robber, err
}

func (rr *robberRepository) InsertOrUpdate(ctx context.Context, robber *model.Robber) error {
	return rr.dbContext.GetDB(ctx).Save(robber).Error
}
