package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type ConstructionService interface {
	GetConstructionRepository() repository.SafeConstructionRepository
	Save(ctx context.Context, construction *datamodel.Construction) error
}

func NewConstructionService(constructionRepository repository.ConstructionRepository) ConstructionService {
	return &constructionService{
		constructionRepository: constructionRepository,
	}
}

type constructionService struct {
	constructionRepository repository.ConstructionRepository
}

func (cs constructionService) GetConstructionRepository() repository.SafeConstructionRepository {
	return cs.constructionRepository
}

func (cs constructionService) Save(ctx context.Context, construction *datamodel.Construction) error {
	if construction.IsRemoved() {
		return nil
	}
	if construction.IsModified() {
		err := cs.constructionRepository.InsertOrUpdate(ctx, construction)
		return errors.Wrap(err, "service.ConstructionService.Save")
	}
	return nil
}
