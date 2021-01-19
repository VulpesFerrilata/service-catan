package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type RoadService interface {
	GetRoadRepository() repository.RoadRepository
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

func (rs roadService) InitRoads() (datamodel.Roads, error) {
	roads := make(datamodel.Roads, 0)

	minQ := 2
	maxQ := 4
	for r := 0; r <= 6; r++ {
		for q := minQ; q <= maxQ; q++ {
			if (r > 0 && q != maxQ) || r > 3 {
				topRoad, err := datamodel.NewRoad(q, r, model.TopLeft)
				if err != nil {
					return nil, errors.Wrap(err, "service.RoadService.InitRoads")
				}
				roads = append(roads, topRoad)
			}

			if r > 0 && r < 6 {
				midRoad, err := datamodel.NewRoad(q, r, model.MiddleLeft)
				if err != nil {
					return nil, errors.Wrap(err, "service.RoadService.InitRoads")
				}
				roads = append(roads, midRoad)
			}

			if (r < 6 && q != maxQ) || r < 3 {
				botRoad, err := datamodel.NewRoad(q, r, model.MiddleLeft)
				if err != nil {
					return nil, errors.Wrap(err, "service.RoadService.InitRoads")
				}
				roads = append(roads, botRoad)
			}
		}

		if r < 3 {
			minQ--
		} else {
			maxQ--
		}
	}

	return roads, nil
}
