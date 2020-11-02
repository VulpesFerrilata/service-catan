package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type SafePlayerRepository interface {
	GetByGameIdByUserId(ctx context.Context, gameId uint, userId uint) (*model.Player, error)
	FindByGameId(ctx context.Context, gameId uint) ([]*model.Player, error)
}

type PlayerRepository interface {
	SafePlayerRepository
	InsertOrUpdate(ctx context.Context, player *model.Player) error
	Delete(ctx context.Context, player *model.Player) error
}
