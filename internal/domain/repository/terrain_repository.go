package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type TerrainRepository interface {
	FindByGameId(ctx context.Context, gameId uuid.UUID) (datamodel.Terrains, error)
	InsertOrUpdate(ctx context.Context, terrain *datamodel.Terrain) error
}

func NewTerrainRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate,
	hexRepository HexRepository) TerrainRepository {
	return &terrainRepository{
		transactionMiddleware: transactionMiddleware,
		validate:              validate,
		hexRepository:         hexRepository,
	}
}

type terrainRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	validate              *validator.Validate
	hexRepository         HexRepository
}

func (t terrainRepository) FindByGameId(ctx context.Context, gameId uuid.UUID) (datamodel.Terrains, error) {
	terrainModels := make([]*model.Terrain, 0)

	err := t.transactionMiddleware.Get(ctx).Find(&terrainModels, "game_id = ?", gameId).Error
	if err != nil {
		return nil, errors.Wrap(err, "repository.TerrainRepository.FindByGameId")
	}

	terrains := make(datamodel.Terrains, 0)
	for _, terrainModel := range terrainModels {
		hex, err := t.hexRepository.GetById(ctx, terrainModel.HexID)
		if err != nil {
			return nil, errors.Wrap(err, "repository.TerrainRepository.FindByGameId")
		}

		terrain, err := datamodel.NewTerrainFromModel(terrainModel, hex)
		if err != nil {
			return nil, errors.Wrap(err, "repository.TerrainRepository.FindByGameId")
		}

		terrains = append(terrains, terrain)
	}

	return terrains, nil
}

func (t terrainRepository) InsertOrUpdate(ctx context.Context, terrain *datamodel.Terrain) error {
	terrainModel := terrain.ToModel()

	if err := t.validate.StructCtx(ctx, terrainModel); err != nil {
		return errors.Wrap(err, "repository.TerrainRepository.InsertOrUpdate")
	}

	if err := t.transactionMiddleware.Get(ctx).Save(terrainModel).Error; err != nil {
		return errors.Wrap(err, "repository.TerrainRepository.InsertOrUpdate")
	}

	hex := terrain.GetHex()
	if err := t.hexRepository.InsertOrUpdate(ctx, hex); err != nil {
		return errors.Wrap(err, "repository.TerrainRepository.InsertOrUpdate")
	}

	return nil
}
