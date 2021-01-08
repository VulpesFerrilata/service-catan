package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
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
