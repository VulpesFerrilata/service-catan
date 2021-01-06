package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
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

func (rs roadService) Save(ctx context.Context, road *datamodel.Road) error {
	if road.IsRemoved() {
		err := rs.roadRepository.Delete(ctx, road)
		return errors.Wrap(err, "service.RoadService.Save")
	}
	if road.IsModified() {
		err := rs.roadRepository.InsertOrUpdate(ctx, road)
		return errors.Wrap(err, "service.RoadService.Save")
	}
	return nil
}
