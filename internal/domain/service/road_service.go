package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type RoadService interface {
	GetRoadRepository() repository.SafeRoadRepository
	Save(ctx context.Context, road *datamodel.Road) error
}

func NewRoadService(roadRepository repository.RoadRepository) RoadService {
	return &roadService{
		roadRepository: roadRepository,
	}
}

type roadService struct {
	roadRepository repository.RoadRepository
}

func (rs roadService) GetRoadRepository() repository.SafeRoadRepository {
	return rs.roadRepository
}
