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
	InsertOrUpdate(ctx context.Context, terrain *datamodel.Terrain) error
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
	if err != nil {
		return nil, errors.Wrap(err, "repository.TerrainRepository.FindByGameId")
	}

	terrains, err := datamodel.NewTerrainsFromTerrainModels(terrainModels)
	return terrains, errors.Wrap(err, "repository.TerrainRepository.FindByGameId")
}

func (tr terrainRepository) InsertOrUpdate(ctx context.Context, terrain *datamodel.Terrain) error {
	terrainModel := terrain.ToModel()

	if err := tr.validate.StructCtx(ctx, terrainModel); err != nil {
		return errors.Wrap(err, "repository.TerrainRepository.InsertOrUpdate")
	}

	err := tr.transactionMiddleware.Get(ctx).Save(terrainModel).Error
	return errors.Wrap(err, "repository.TerrainRepository.InsertOrUpdate")
}
