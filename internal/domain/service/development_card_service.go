package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type DevelopmentCardService interface {
	GetDevelopmentCardRepository() repository.DevelopmentCardRepository
}

func NewDevelopmentCardService(developmentCardRepository repository.DevelopmentCardRepository) DevelopmentCardService {
	return &developmentCardService{
		developmentCardRepository: developmentCardRepository,
	}
}

type developmentCardService struct {
	developmentCardRepository repository.DevelopmentCardRepository
}

func (dcs developmentCardService) GetDevelopmentCardRepository() repository.DevelopmentCardRepository {
	return dcs.developmentCardRepository
}
