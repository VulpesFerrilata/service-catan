package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type RobberService interface {
	GetRobberRepository() repository.RobberRepository
}

func NewRobberService(robberRepository repository.RobberRepository) RobberService {
	return &robberService{
		robberRepository: robberRepository,
	}
}

type robberService struct {
	robberRepository repository.RobberRepository
}

func (rs robberService) GetRobberRepository() repository.RobberRepository {
	return rs.robberRepository
}
