package repository

import (
	"context"
	"errors"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/db"
	server_errors "github.com/VulpesFerrilata/library/pkg/errors"
	"gorm.io/gorm"
)

type ReadOnlyGameRepository interface {
	GetById(ctx context.Context, id uint) (*model.Game, error)
}

type GameRepository interface {
	ReadOnlyGameRepository
	Insert(ctx context.Context, game *model.Game) error
}

func NewGameRepository(db *db.DbContext) GameRepository {
	return &gameRepository{
		db: db,
	}
}

type gameRepository struct {
	db *db.DbContext
}

func (gr gameRepository) GetById(ctx context.Context, id uint) (*model.Game, error) {
	game := new(model.Game)
	err := gr.db.GetDB(ctx).First(game, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = server_errors.NewNotFoundError("game")
	}
	return game, gr.db.GetDB(ctx).First(game, id).Error
}

func (gr gameRepository) Insert(ctx context.Context, game *model.Game) error {
	return gr.db.GetDB(ctx).Create(game).Error
}
