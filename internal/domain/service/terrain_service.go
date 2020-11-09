package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type TerrainService interface {
	GetFieldRepository() repository.SafeTerrainRepository
	Save(ctx context.Context, terrain *model.Terrain) error
}

type terrainService struct {
	terrainRepository repository.TerrainRepository
}

func (ts *terrainService) GetFieldRepository() repository.SafeTerrainRepository {
	return ts.terrainRepository
}

func (ts *terrainService) validate(ctx context.Context, terrain *model.Terrain) error {
	//TODO: validate terrain
	return nil
}

func (ts *terrainService) Save(ctx context.Context, terrain *model.Terrain) error {
	if err := ts.validate(ctx, terrain); err != nil {
		return err
	}

	return ts.terrainRepository.InsertOrUpdate(ctx, terrain)
}
