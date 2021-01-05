package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
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

func (dcs developmentCardService) Save(ctx context.Context, developmentCard *datamodel.DevelopmentCard) error {
	if developmentCard.IsRemoved() {
		return nil
	}
	if developmentCard.IsModified() {
		err := dcs.developmentCardRepository.InsertOrUpdate(ctx, developmentCard)
		return errors.Wrap(err, "service.DevelopmentCardService.Save")
	}
	return nil
}
