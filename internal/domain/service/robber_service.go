package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type RobberService interface {
	GetRobberRepository() repository.SafeRobberRepository
	Save(ctx context.Context, robber *datamodel.Robber) error
}

func NewRobberService(robberRepository repository.RobberRepository) RobberService {
	return &robberService{
		robberRepository: robberRepository,
	}
}

type robberService struct {
	robberRepository repository.RobberRepository
}

func (rs robberService) GetRobberRepository() repository.SafeRobberRepository {
	return rs.robberRepository
}

func (rs robberService) Save(ctx context.Context, robber *datamodel.Robber) error {
	if robber.IsRemoved() {
		err := rs.robberRepository.Delete(ctx, robber)
		return errors.Wrap(err, "service.RobberService.Save")
	}
	if robber.IsModified() {
		err := rs.robberRepository.InsertOrUpdate(ctx, robber)
		return errors.Wrap(err, "service.RobberService.Save")
	}
	return nil
}
