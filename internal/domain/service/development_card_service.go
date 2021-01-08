package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type DevelopmentCardService interface {
	GetDevelopmentCardRepository() repository.SafeDevelopmentCardRepository
	Save(ctx context.Context, developmentCard *datamodel.DevelopmentCard) error
}

func NewDevelopmentCardService(developmentCardRepository repository.DevelopmentCardRepository) DevelopmentCardService {
	return &developmentCardService{
		developmentCardRepository: developmentCardRepository,
	}
}

type developmentCardService struct {
	developmentCardRepository repository.DevelopmentCardRepository
}

func (dcs developmentCardService) GetDevelopmentCardRepository() repository.SafeDevelopmentCardRepository {
	return dcs.developmentCardRepository
}
