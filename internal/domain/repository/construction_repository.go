package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type SafeConstructionRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Constructions, error)
}

type ConstructionRepository interface {
	SafeConstructionRepository
	InsertOrUpdate(ctx context.Context, construction *model.Construction) error
}
