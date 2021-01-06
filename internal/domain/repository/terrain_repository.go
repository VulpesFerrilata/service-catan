package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/app_error"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type SafeTerrainRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (datamodel.Terrains, error)
}

type TerrainRepository interface {
	SafeTerrainRepository
	InsertOrUpdate(ctx context.Context, terrain *datamodel.Terrain) error
	Delete(ctx context.Context, terrain *datamodel.Terrain) error
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

func (tr terrainRepository) InsertOrUpdate(ctx context.Context, terrain *datamodel.Terrain) error {
	return terrain.Persist(func(terrainModel *model.Terrain) error {
		if err := tr.validate.StructCtx(ctx, terrainModel); err != nil {
			if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
				err = app_error.NewEntityValidationError(terrainModel, fieldErrors)
			}
			return errors.Wrap(err, "repository.TerrainRepository.InsertOrUpdate")
		}
		err := tr.transactionMiddleware.Get(ctx).Save(terrainModel).Error
		return errors.Wrap(err, "repository.TerrainRepository.InsertOrUpdate")
	})
}

func (tr terrainRepository) Delete(ctx context.Context, terrain *datamodel.Terrain) error {
	return terrain.Persist(func(terrainModel *model.Terrain) error {
		err := tr.transactionMiddleware.Get(ctx).Delete(terrainModel).Error
		return errors.Wrap(err, "repository.TerrainRepository.Delete")
	})
}
