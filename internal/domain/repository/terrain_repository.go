package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/db"
)

type SafeTerrainRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Terrains, error)
}

type TerrainRepository interface {
	SafeTerrainRepository
	InsertOrUpdate(ctx context.Context, terrain *model.Terrain) error
}

func NewTerrainRepository(dbContext *db.DbContext) TerrainRepository {
	return &terrainRepository{
		dbContext: dbContext,
	}
}

type terrainRepository struct {
	dbContext *db.DbContext
}

func (tr *terrainRepository) FindByGameId(ctx context.Context, gameId uint) (model.Terrains, error) {
	var terrains model.Terrains
	return terrains, tr.dbContext.GetDB(ctx).Find(&terrains, "game_id = ?", gameId).Error
}

func (tr *terrainRepository) InsertOrUpdate(ctx context.Context, terrain *model.Terrain) error {
	return tr.dbContext.GetDB(ctx).Save(terrain).Error
}
