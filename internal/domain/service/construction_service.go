package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type ConstructionService interface {
	GetConstructionRepository() repository.ConstructionRepository
}

func NewConstructionService(constructionRepository repository.ConstructionRepository) ConstructionService {
	return &constructionService{
		constructionRepository: constructionRepository,
	}
}

type constructionService struct {
	constructionRepository repository.ConstructionRepository
}

func (cs constructionService) GetConstructionRepository() repository.ConstructionRepository {
	return cs.constructionRepository
}
