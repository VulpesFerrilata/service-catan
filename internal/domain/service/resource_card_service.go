package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type ResourceCardService interface {
	GetResourceCardRepository() repository.ResourceCardRepository
	InitResourceCards() (datamodel.ResourceCards, error)
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

func (rcs resourceCardService) InitResourceCards() (datamodel.ResourceCards, error) {
	resourceCards := make(datamodel.ResourceCards, 0)

	resourceTypes := map[model.ResourceType]int{
		model.LumberResource: 19,
		model.BrickResource:  19,
		model.WoolResource:   19,
		model.GrainResource:  19,
		model.OreResource:    19,
	}
	for resourceType, quantity := range resourceTypes {
		for i := 1; i <= quantity; i++ {
			resourceCard, err := datamodel.NewResourceCard(resourceType)
			if err != nil {
				return nil, errors.Wrap(err, "service.ResourceCardService.InitResourceCards")
			}
			resourceCards = append(resourceCards, resourceCard)
		}
	}

	return resourceCards, nil
}
