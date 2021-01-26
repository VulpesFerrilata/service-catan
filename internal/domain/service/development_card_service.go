package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type DevelopmentCardService interface {
	GetDevelopmentCardRepository() repository.DevelopmentCardRepository
	InitDevelopmentCards() (datamodel.DevelopmentCards, error)
}

func NewDevelopmentCardService(developmentCardRepository repository.DevelopmentCardRepository) DevelopmentCardService {
	return &developmentCardService{
		developmentCardRepository: developmentCardRepository,
	}
}

type developmentCardService struct {
	developmentCardRepository repository.DevelopmentCardRepository
}

func (d developmentCardService) GetDevelopmentCardRepository() repository.DevelopmentCardRepository {
	return d.developmentCardRepository
}

func (d developmentCardService) InitDevelopmentCards() (datamodel.DevelopmentCards, error) {
	developmentCards := make(datamodel.DevelopmentCards, 0)

	developmentTypes := map[model.DevelopmentType]int{
		model.Knight:        14,
		model.VictoryPoints: 5,
		model.RoadBuilding:  2,
		model.YearOfPlenty:  2,
		model.Monopoly:      2,
	}

	for developmentType, quantity := range developmentTypes {
		for i := 1; i <= quantity; i++ {
			developmentCard, err := datamodel.NewDevelopmentCard(developmentType)
			if err != nil {
				return nil, errors.Wrap(err, "service.DevelopmentCardService.InitDevelopmentCards")
			}
			developmentCards = append(developmentCards, developmentCard)
		}
	}

	return developmentCards, nil
}
