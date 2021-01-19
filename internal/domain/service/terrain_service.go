package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
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

func (ts terrainService) InitTerrains() (datamodel.Terrains, error) {
	terrains := make(datamodel.Terrains, 0)

	numbers := []int{2, 3, 3, 4, 4, 5, 5, 6, 6, 8, 8, 9, 9, 10, 10, 11, 11, 12}
	terrainTypes := []model.TerrainType{
		model.DesertTerrain,
		model.FieldTerrain,
		model.FieldTerrain,
		model.FieldTerrain,
		model.FieldTerrain,
		model.ForestTerrain,
		model.ForestTerrain,
		model.ForestTerrain,
		model.ForestTerrain,
		model.PastureTerrain,
		model.PastureTerrain,
		model.PastureTerrain,
		model.PastureTerrain,
		model.MountainTerrain,
		model.MountainTerrain,
		model.MountainTerrain,
		model.HillTerrain,
		model.HillTerrain,
		model.HillTerrain,
	}

	minQ := 1
	maxQ := 3
	for r := 1; r <= 5; r++ {
		for q := minQ; q <= maxQ; q++ {
			var terrainType model.TerrainType
			var number int
			terrainType, terrainTypes = terrainTypes[0], terrainTypes[1:]
			if terrainType == model.DesertTerrain {
				number = 7
			} else {
				number, numbers = numbers[0], numbers[1:]
			}
			terrain, err := datamodel.NewTerrain(q, r, number, terrainType)
			if err != nil {
				return nil, errors.Wrap(err, "service.TerrainService.InitTerrains")
			}
			terrains = append(terrains, terrain)
		}

		if r < 3 {
			minQ--
		} else {
			maxQ--
		}
	}

	return terrains, nil
}
