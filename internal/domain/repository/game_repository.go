package repository

import (
	"context"
	"errors"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/db"
	server_errors "github.com/VulpesFerrilata/library/pkg/errors"
	"gorm.io/gorm"
)

type SafeGameRepository interface {
	GetById(ctx context.Context, id uint) (*model.Game, error)
	FindByGameStatus(ctx context.Context, gameStatus datamodel.GameStatus) ([]*model.Game, error)
}

type GameRepository interface {
	SafeGameRepository
	InsertOrUpdate(ctx context.Context, game *model.Game) error
	Delete(ctx context.Context, game *model.Game) error
}

func NewGameRepository(dbContext *db.DbContext) GameRepository {
	return &gameRepository{
		dbContext: dbContext,
	}
}

type gameRepository struct {
	dbContext *db.DbContext
}

func (gr *gameRepository) GetById(ctx context.Context, gameId uint) (*model.Game, error) {
	game := new(model.Game)
	err := gr.dbContext.GetDB(ctx).First(&game, "game_id = ?", gameId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, server_errors.NewNotFoundError("game")
	}
	return game, err
}

func (gr *gameRepository) FindByGameStatus(ctx context.Context, gameStatus datamodel.GameStatus) ([]*model.Game, error) {
	games := make([]*model.Game, 0)
	return games, gr.dbContext.GetDB(ctx).Find(&games, "status", gameStatus).Error
}

func (gr *gameRepository) InsertOrUpdate(ctx context.Context, game *model.Game) error {
	return gr.dbContext.GetDB(ctx).Save(game).Error
}

func (gr *gameRepository) Delete(ctx context.Context, game *model.Game) error {
	return gr.dbContext.GetDB(ctx).Delete(game).Error
}
