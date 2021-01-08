package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
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
