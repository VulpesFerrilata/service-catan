package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type ReadOnlyGameRepository interface {
	GetById(ctx context.Context, id uint) (*model.Game, error)
}

type GameRepository interface {
	ReadOnlyGameRepository
	Insert(ctx context.Context, game *model.Game) error
}
