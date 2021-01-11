package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type TerrainRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (datamodel.Terrains, error)
	Save(ctx context.Context, terrain *datamodel.Terrain) error
}

func NewTerrainRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate) TerrainRepository {
	return &terrainRepository{
		transactionMiddleware: transactionMiddleware,
		validate:              validate,
	}
}

type terrainRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	validate              *validator.Validate
}

func (tr terrainRepository) FindByGameId(ctx context.Context, gameId uint) (datamodel.Terrains, error) {
	terrainModels := make([]*model.Terrain, 0)
	err := tr.transactionMiddleware.Get(ctx).Find(&terrainModels, "game_id = ?", gameId).Error
	return datamodel.NewTerrainsFromTerrainModels(terrainModels), errors.Wrap(err, "repository.TerrainRepository.FindByGameId")
}

func (tr terrainRepository) insertOrUpdate(ctx context.Context, terrain *datamodel.Terrain) error {
	return terrain.Persist(func(terrainModel *model.Terrain) error {
		if err := tr.validate.StructCtx(ctx, terrainModel); err != nil {
			return errors.Wrap(err, "repository.TerrainRepository.InsertOrUpdate")
		}
		err := tr.transactionMiddleware.Get(ctx).Save(terrainModel).Error
		return errors.Wrap(err, "repository.TerrainRepository.InsertOrUpdate")
	})
}

func (tr terrainRepository) delete(ctx context.Context, terrain *datamodel.Terrain) error {
	return terrain.Persist(func(terrainModel *model.Terrain) error {
		err := tr.transactionMiddleware.Get(ctx).Delete(terrainModel).Error
		return errors.Wrap(err, "repository.TerrainRepository.Delete")
	})
}

func (tr terrainRepository) Save(ctx context.Context, terrain *datamodel.Terrain) error {
	if terrain.IsRemoved() {
		err := tr.delete(ctx, terrain)
		return errors.Wrap(err, "service.TerrainRepository.Save")
	}
	if terrain.IsModified() {
		err := tr.insertOrUpdate(ctx, terrain)
		return errors.Wrap(err, "service.TerrainRepository.Save")
	}
	return nil
}
