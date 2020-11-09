package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type SafeTerrainRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Terrains, error)
}

type TerrainRepository interface {
	SafeTerrainRepository
	InsertOrUpdate(ctx context.Context, terrain *model.Terrain) error
}
