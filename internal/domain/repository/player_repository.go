package repository

import (
	"context"

	model "github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

type ReadOnlyPlayerRepository interface {
	GetByGameIdByUserId(ctx context.Context, gameId uint, userId uint) (*model.Player, error)
	GetByUserId(ctx context.Context, userId uint) ([]*model.Player, error)
	FindByGameId(ctx context.Context, gameId uint) (*model.Player, error)
}

type PlayerRepository interface {
	ReadOnlyPlayerRepository
	Insert(ctx context.Context, players ...*model.Player) error
	Save(ctx context.Context, players ...*model.Player) error
}
