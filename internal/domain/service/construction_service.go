package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
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

func (cs constructionService) InitConstructions() (datamodel.Constructions, error) {
	constructions := make(datamodel.Constructions, 0)

	minQ := 1
	maxQ := 4
	for r := 0; r <= 6; r++ {
		for q := minQ; q <= maxQ; q++ {
			if (r > 0 && q != minQ && q != maxQ) || r > 3 {
				topConstruction, err := datamodel.NewConstruction(q, r, model.Top, model.Land)
				if err != nil {
					return nil, errors.Wrap(err, "service.ConstructionService.InitConstructions")
				}
				constructions = append(constructions, topConstruction)
			}
			if (r < 6 && q != minQ && q != maxQ) || r < 3 {
				botConstruction, err := datamodel.NewConstruction(q, r, model.Bottom, model.Land)
				if err != nil {
					return nil, errors.Wrap(err, "service.ConstructionService.InitConstructions")
				}
				constructions = append(constructions, botConstruction)
			}
		}

		if r < 3 {
			minQ--
		} else {
			maxQ--
		}
	}

	return constructions, nil
}
