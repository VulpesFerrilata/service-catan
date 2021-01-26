package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type ConstructionService interface {
	GetConstructionRepository() repository.ConstructionRepository
	InitConstructions(terrains datamodel.Terrains) (datamodel.Constructions, error)
}

func NewConstructionService(constructionRepository repository.ConstructionRepository) ConstructionService {
	return &constructionService{
		constructionRepository: constructionRepository,
	}
}

type constructionService struct {
	constructionRepository repository.ConstructionRepository
}

func (c constructionService) GetConstructionRepository() repository.ConstructionRepository {
	return c.constructionRepository
}

func (c constructionService) InitConstructions(terrains datamodel.Terrains) (datamodel.Constructions, error) {
	constructions := make(datamodel.Constructions, 0)

	setHexCorners := make([]*datamodel.HexCorner, 0)
	for _, terrain := range terrains {
		// hex corners may duplicate among terrains
		hexCorners := terrain.GetHex().GetPossibleAdjacentHexCorners()
		for _, hexCorner := range hexCorners {
			isDuplicated := false
			for _, setHexCorner := range setHexCorners {
				if setHexCorner.Equals(hexCorner) {
					isDuplicated = true
					break
				}
			}
			if !isDuplicated {
				setHexCorners = append(setHexCorners, hexCorner)
			}
		}
	}

	for _, setHexCorner := range setHexCorners {
		construction, err := datamodel.NewConstruction(setHexCorner)
		if err != nil {
			return nil, errors.Wrap(err, "service.ConstructionService.InitConstructions")
		}
		constructions = append(constructions, construction)
	}

	return constructions, nil
}
