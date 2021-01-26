package service

import (
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type TerrainService interface {
	GetTerrainRepository() repository.TerrainRepository
	InitTerrains() (datamodel.Terrains, error)
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
	//circle vectors
	hexVectors := []*datamodel.HexVector{
		{Q: 0, R: -1}, //top left
		{Q: 1, R: -1}, //top right
		{Q: 1, R: 0},  //middle right
		{Q: 0, R: 1},  //bottom right
		{Q: -1, R: 1}, //bottom left
		{Q: -1, R: 0}, //middle left
	}
	if rand.Intn(2) == 0 {
		//reverse direction
		for i, j := 0, len(hexVectors)-1; i < j; i, j = i+1, j-1 {
			hexVectors[i], hexVectors[j] = hexVectors[j], hexVectors[i]
		}
	}
	//random corner in circle directions
	randIdx := rand.Intn(len(hexVectors))
	rootVector := hexVectors[randIdx]
	//direction will start moving at next 2 index confront root direction
	//example: top left corner hex will start moving from middle right or bottom left
	randIdx += 2
	if randIdx > len(hexVectors)-1 {
		randIdx -= len(hexVectors)
	}
	hexVectors = append(hexVectors[randIdx:], hexVectors[:randIdx]...)

	//spiral hexes start at specified root direction's hex - corner - and end at center
	spiralHexes := make([]*datamodel.Hex, 0)
	centerHex, err := datamodel.NewHex(0, 0)
	if err != nil {
		return nil, errors.Wrap(err, "service.TerrainService.InitTerrains")
	}
	for radius := 2; radius >= 1; radius-- {
		//corner hex
		vector := &datamodel.HexVector{
			Q: rootVector.Q * radius,
			R: rootVector.R * radius,
		}
		hex, err := datamodel.NewHexFromVector(centerHex, vector)
		if err != nil {
			return nil, errors.Wrap(err, "service.TerrainService.InitTerrains")
		}
		spiralHexes = append(spiralHexes, hex)

		//circle hexes
		for _, hexVector := range hexVectors {
			for i := 1; i <= radius; i++ {
				hex, err = datamodel.NewHexFromVector(hex, hexVector)
				if err != nil {
					return nil, errors.Wrap(err, "service.TerrainService.InitTerrains")
				}
				spiralHexes = append(spiralHexes, hex)
			}
		}
	}
	spiralHexes = append(spiralHexes, centerHex)

	numbers := []int{
		5,  //A:****
		2,  //B:*
		6,  //C:*****
		3,  //D:**
		8,  //E:*****
		10, //F:***
		9,  //G:****
		12, //H:*
		11, //I:**
		4,  //J:***
		8,  //K:*****
		10, //L:***
		9,  //M:****
		4,  //N:***
		5,  //O:****
		6,  //P:*****
		3,  //Q:**
		11, //R:**
	}

	terrainTypes := []datamodel.TerrainType{
		datamodel.DesertTerrain, //1 x desert terrain
		datamodel.FieldTerrain,  //4 x field terrain
		datamodel.FieldTerrain,
		datamodel.FieldTerrain,
		datamodel.FieldTerrain,
		datamodel.ForestTerrain, //4 x forest terrain
		datamodel.ForestTerrain,
		datamodel.ForestTerrain,
		datamodel.ForestTerrain,
		datamodel.PastureTerrain, //4 x pasture terrain
		datamodel.PastureTerrain,
		datamodel.PastureTerrain,
		datamodel.PastureTerrain,
		datamodel.MountainTerrain, //3 x mountain terrain
		datamodel.MountainTerrain,
		datamodel.MountainTerrain,
		datamodel.HillTerrain, //3 x hill terrain
		datamodel.HillTerrain,
		datamodel.HillTerrain,
	}
	rand.Shuffle(len(terrainTypes), func(i, j int) {
		terrainTypes[i], terrainTypes[j] = terrainTypes[j], terrainTypes[i]
	})

	terrains := make(datamodel.Terrains, 0)
	for idx, hex := range spiralHexes {
		terrainType := terrainTypes[idx]

		var number int
		if terrainType == datamodel.DesertTerrain {
			number = 7
		} else {
			number = numbers[0]
			numbers = numbers[1:]
		}

		terrain, err := datamodel.NewTerrain(hex, number, terrainType)
		if err != nil {
			return nil, errors.Wrap(err, "service.TerrainService.InitTerrains")
		}

		terrains = append(terrains, terrain)
	}

	return terrains, nil
}
