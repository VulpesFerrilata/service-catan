package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type ResourceCardService interface {
	GetResourceCardRepository() repository.ResourceCardRepository
}

func NewResourceCardService(resourceCardRepository repository.ResourceCardRepository) ResourceCardService {
	return &resourceCardService{
		resourceCardRepository: resourceCardRepository,
	}
}

type resourceCardService struct {
	resourceCardRepository repository.ResourceCardRepository
}

func (rcs resourceCardService) GetResourceCardRepository() repository.ResourceCardRepository {
	return rcs.resourceCardRepository
}
