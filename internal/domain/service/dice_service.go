package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
)

type DiceService interface {
	GetDiceRepository() repository.DiceRepository
	InitDices() (datamodel.Dices, error)
}

func NewDiceService(diceRepository repository.DiceRepository) DiceService {
	return &diceService{
		diceRepository: diceRepository,
	}
}

type diceService struct {
	diceRepository repository.DiceRepository
}

func (d diceService) GetDiceRepository() repository.DiceRepository {
	return d.diceRepository
}

func (d diceService) InitDices() (datamodel.Dices, error) {
	dices := make(datamodel.Dices, 0)

	for i := 1; i <= 2; i++ {
		dice, err := datamodel.NewDice(6)
		if err != nil {
			return nil, errors.Wrap(err, "service.DiceService.InitDices")
		}
		dices = append(dices, dice)
	}

	return dices, nil
}
