package repository

import (
	"context"
	"errors"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/db"
	server_errors "github.com/VulpesFerrilata/library/pkg/errors"
	"gorm.io/gorm"
)

type ReadOnlyPlayerRepository interface {
	GetByGameIdByUserId(ctx context.Context, gameId uint, userId uint) (*model.Player, error)
	FindByGameId(ctx context.Context, gameId uint) ([]*model.Player, error)
}

type PlayerRepository interface {
	ReadOnlyPlayerRepository
	Insert(ctx context.Context, players ...*model.Player) error
	Save(ctx context.Context, players ...*model.Player) error
}

func NewPlayerRepository(db *db.DbContext) PlayerRepository {
	return &playerRepository{
		db: db,
	}
}

type playerRepository struct {
	db *db.DbContext
}

func (pr playerRepository) GetByGameIdByUserId(ctx context.Context, gameId uint, userId uint) (*model.Player, error) {
	player := new(model.Player)
	err := pr.db.GetDB(ctx).First(player, "game_id = ?", gameId, "user_id = ?", userId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = server_errors.NewNotFoundError("player")
	}
	return player, err
}

func (pr playerRepository) Insert(ctx context.Context, players ...*model.Player) error {
	return pr.db.GetDB(ctx).Create(players).Error
}

func (pr playerRepository) Save(ctx context.Context, players ...*model.Player) error {
	return pr.db.GetDB(ctx).Save(players).Error
}
