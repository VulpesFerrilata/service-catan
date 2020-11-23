package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type ConstructionService interface {
	GetConstructionRepository() repository.SafeConstructionRepository
	Save(ctx context.Context, construction *model.Construction) error
}

func NewConstructionService(constructionRepository repository.ConstructionRepository) ConstructionService {
	return &constructionService{
		constructionRepository: constructionRepository,
	}
}

type constructionService struct {
	constructionRepository repository.ConstructionRepository
}

func (cs *constructionService) GetConstructionRepository() repository.SafeConstructionRepository {
	return cs.constructionRepository
}

func (cs *constructionService) validate(ctx context.Context, construction *model.Construction) error {
	//TODO: validate construction
	return nil
}

func (cs *constructionService) Save(ctx context.Context, construction *model.Construction) error {
	if err := cs.validate(ctx, construction); err != nil {
		return err
	}

	return cs.constructionRepository.InsertOrUpdate(ctx, construction)
}
