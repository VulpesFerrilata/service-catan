package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type TerrainService interface {
	GetTerrainRepository() repository.TerrainRepository
}

func NewTerrainService(terrainRepository repository.TerrainRepository) TerrainService {
	return &terrainService{
		terrainRepository: terrainRepository,
	}
}

type terrainService struct {
	terrainRepository repository.TerrainRepository
}

func (ts terrainService) GetTerrainRepository() repository.TerrainRepository {
	return ts.terrainRepository
}
