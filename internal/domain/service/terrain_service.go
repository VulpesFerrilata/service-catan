package service

import (
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type Hex struct {
	Q int
	R int
}

func (h Hex) GetNeightbor(hex Hex) *Hex {
	neightborHex := new(Hex)
	neightborHex.Q = h.Q + hex.Q
	neightborHex.R = h.R + hex.R
	return neightborHex
}

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
	//circle directions
	directions := [][]int{
		{0, -1}, //top left
		{1, -1}, //top right
		{1, 0},  //middle right
		{0, 1},  //bottom right
		{-1, 1}, //bottom left
		{-1, 0}, //middle left
	}
	if rand.Intn(2) == 0 {
		//reverse direction
		for i, j := 0, len(directions)-1; i < j; i, j = i+1, j-1 {
			directions[i], directions[j] = directions[j], directions[i]
		}
	}
	//random corner in circle directions
	randIdx := rand.Intn(len(directions))
	rootDirection := directions[randIdx]
	randIdx += 2
	if randIdx > len(directions) - 1 {
		randIdx -= len(directions)
	}
	directions = append(directions[randIdx:], directions[:randIdx]...)

	spiralCoords := make([][]int, 0)
	centerCoord := []int{0, 0}
	for radius := 2, radius >= 0, radius-- {
		coord := []int {rootDirection[0] 
	}

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
