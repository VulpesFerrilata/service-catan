package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type RoadService interface {
	GetRoadRepository() repository.RoadRepository
	InitRoads(terrains datamodel.Terrains) (datamodel.Roads, error)
}

func NewRoadService(roadRepository repository.RoadRepository) RoadService {
	return &roadService{
		roadRepository: roadRepository,
	}
}

type roadService struct {
	roadRepository repository.RoadRepository
}

func (rs roadService) GetRoadRepository() repository.RoadRepository {
	return rs.roadRepository
}

func (rs roadService) InitRoads(terrains datamodel.Terrains) (datamodel.Roads, error) {
	roads := make(datamodel.Roads, 0)

	setHexEdges := make([]*datamodel.HexEdge, 0)
	for _, terrain := range terrains {
		// hex edges may duplicate among terrains
		hexEdges := terrain.GetHex().GetPossibleAdjacentHexEdges()
		for _, hexEdge := range hexEdges {
			isDuplicated := false
			for _, setHexEdge := range setHexEdges {
				if setHexEdge.Equals(hexEdge) {
					isDuplicated = true
					break
				}
			}
			if !isDuplicated {
				setHexEdges = append(setHexEdges, hexEdge)
			}
		}
	}

	for _, setHexEdge := range setHexEdges {
		road, err := datamodel.NewRoad(setHexEdge)
		if err != nil {
			return nil, errors.Wrap(err, "service.RoadService.InitRoads")
		}
		roads = append(roads, road)
	}

	return roads, nil
}
