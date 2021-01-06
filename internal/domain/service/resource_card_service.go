package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type ResourceCardService interface {
	GetResourceCardRepository() repository.SafeResourceCardRepository
	Save(ctx context.Context, resourceCard *datamodel.ResourceCard) error
}

func NewResourceCardService(resourceCardRepository repository.ResourceCardRepository) ResourceCardService {
	return &resourceCardService{
		resourceCardRepository: resourceCardRepository,
	}
}

type resourceCardService struct {
	resourceCardRepository repository.ResourceCardRepository
}

func (rcs resourceCardService) GetResourceCardRepository() repository.SafeResourceCardRepository {
	return rcs.resourceCardRepository
}

func (rcs resourceCardService) Save(ctx context.Context, resourceCard *datamodel.ResourceCard) error {
	if resourceCard.IsRemoved() {
		err := rcs.resourceCardRepository.Delete(ctx, resourceCard)
		return errors.Wrap(err, "service.ResourceCardService.Save")
	}
	if resourceCard.IsModified() {
		err := rcs.resourceCardRepository.InsertOrUpdate(ctx, resourceCard)
		return errors.Wrap(err, "service.ResourceCardService.Save")
	}
	return nil
}
