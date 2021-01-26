package service

import (
	"math/rand"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type HarborService interface {
	GetHarborRepository() repository.HarborRepository
	InitHarbors(terrains datamodel.Terrains) (datamodel.Harbors, error)
}

func NewHarborService(harborRepository repository.HarborRepository) HarborService {
	return &harborService{
		harborRepository: harborRepository,
	}
}

type harborService struct {
	harborRepository repository.HarborRepository
}

func (h harborService) GetHarborRepository() repository.HarborRepository {
	return h.harborRepository
}

func (h harborService) InitHarbors(terrains datamodel.Terrains) (datamodel.Harbors, error) {
	harbors := make(datamodel.Harbors, 0)

	harborTypes := []datamodel.HarborType{
		datamodel.GeneralHarbor,
		datamodel.GeneralHarbor,
		datamodel.GeneralHarbor,
		datamodel.GeneralHarbor,
		datamodel.GeneralHarbor,
		datamodel.LumberHarbor,
		datamodel.BrickHarbor,
		datamodel.WoolHarbor,
		datamodel.GrainHarbor,
		datamodel.OreHarbor,
	}
	rand.Shuffle(len(harborTypes), func(i, j int) {
		harborTypes[i], harborTypes[j] = harborTypes[j], harborTypes[i]
	})

	//circle vectors
	hexVectors := []*datamodel.HexVector{
		{Q: 0, R: -1}, //top left
		{Q: 1, R: -1}, //top right
		{Q: 1, R: 0},  //middle right
		{Q: 0, R: 1},  //bottom right
		{Q: -1, R: 1}, //bottom left
		{Q: -1, R: 0}, //middle left
	}

	rootVector := hexVectors[0]
	hexVectors = append(hexVectors[2:], hexVectors[:2]...)

	circleHexes := make([]*datamodel.Hex, 0)
	centerHex, err := datamodel.NewHex(0, 0)
	if err != nil {
		return nil, errors.Wrap(err, "service.HarborService.InitHarbors")
	}
	radius := 3
	vector := &datamodel.HexVector{
		Q: rootVector.Q * radius,
		R: rootVector.R * radius,
	}
	hex, err := datamodel.NewHexFromVector(centerHex, vector)
	if err != nil {
		return nil, errors.Wrap(err, "service.HarborService.InitHarbors")
	}
	circleHexes = append(circleHexes, hex)

	for _, hexVector := range hexVectors {
		for i := 1; i <= 3; i++ {
			hex, err = datamodel.NewHexFromVector(hex, hexVector)
			if err != nil {
				return nil, errors.Wrap(err, "service.HarborService.InitHarbors")
			}
			circleHexes = append(circleHexes, hex)
		}
	}

	remainder := rand.Intn(2)
	for idx, circleHex := range circleHexes {
		if idx%2 == remainder {
			adjacentHexes, err := circleHex.GetPossibleAdjacentHexes()
			if err != nil {
				return nil, errors.Wrap(err, "service.HarborService.InitHarbors")
			}
			adjacentTerrains := terrains.Filter(func(terrain *datamodel.Terrain) bool {
				for _, adjacentHex := range adjacentHexes {
					if terrain.GetHex().Equals(adjacentHex) {
						return true
					}
				}
				return false
			})
			randAdjacentTerrain := adjacentTerrains[rand.Intn(len(adjacentTerrains))]

			harborType := harborTypes[0]
			harborTypes = harborTypes[1:]

			harbor, err := datamodel.NewHarbor(circleHex, harborType, randAdjacentTerrain)
			if err != nil {
				return nil, errors.Wrap(err, "service.HarborService.InitHarbors")
			}
			harbors = append(harbors, harbor)
		}
	}

	return harbors, nil
}
