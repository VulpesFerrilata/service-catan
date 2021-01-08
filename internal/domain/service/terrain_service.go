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
