package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type RoadService interface {
	GetRoadRepository() repository.SafeRoadRepository
	Save(ctx context.Context, road *model.Road) error
}

type roadService struct {
	roadRepository repository.RoadRepository
}

func (rs *roadService) GetRoadRepository() repository.SafeRoadRepository {
	return rs.roadRepository
}

func (rs *roadService) validate(ctx context.Context, road *model.Road) error {
	//TODO: validate road
	return nil
}

func (rs *roadService) Save(ctx context.Context, road *model.Road) error {
	if err := rs.validate(ctx, road); err != nil {
		return err
	}

	return rs.roadRepository.InsertOrUpdate(ctx, road)
}
