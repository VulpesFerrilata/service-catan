package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
)

type SafeTerrainRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Terrains, error)
}

type TerrainRepository interface {
	SafeTerrainRepository
	InsertOrUpdate(ctx context.Context, terrain *model.Terrain) error
}

func NewTerrainRepository(transactionMiddleware *middleware.TransactionMiddleware) TerrainRepository {
	return &terrainRepository{
		transactionMiddleware: transactionMiddleware,
	}
}

type terrainRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
}

func (tr *terrainRepository) FindByGameId(ctx context.Context, gameId uint) (model.Terrains, error) {
	var terrains model.Terrains
	return terrains, tr.transactionMiddleware.Get(ctx).Find(&terrains, "game_id = ?", gameId).Error
}

func (tr *terrainRepository) InsertOrUpdate(ctx context.Context, terrain *model.Terrain) error {
	return tr.transactionMiddleware.Get(ctx).Save(terrain).Error
}
