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