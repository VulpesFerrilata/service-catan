package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type TerrainService interface {
	GetTerrainRepository() repository.SafeTerrainRepository
	Save(ctx context.Context, terrain *datamodel.Terrain) error
}

func NewTerrainService(terrainRepository repository.TerrainRepository) TerrainService {
	return &terrainService{
		terrainRepository: terrainRepository,
	}
}

type terrainService struct {
	terrainRepository repository.TerrainRepository
}

func (ts terrainService) GetTerrainRepository() repository.SafeTerrainRepository {
	return ts.terrainRepository
}

func (ts terrainService) Save(ctx context.Context, terrain *datamodel.Terrain) error {
	if terrain.IsRemoved() {
		err := ts.terrainRepository.Delete(ctx, terrain)
		return errors.Wrap(err, "service.TerrainService.Save")
	}
	if terrain.IsModified() {
		err := ts.terrainRepository.InsertOrUpdate(ctx, terrain)
		return errors.Wrap(err, "service.TerrainService.Save")
	}
	return nil
}
