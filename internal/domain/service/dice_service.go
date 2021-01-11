package service

import (
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type DiceService interface {
	GetDiceRepository() repository.DiceRepository
}

func NewDiceService(diceRepository repository.DiceRepository) DiceService {
	return &diceService{
		diceRepository: diceRepository,
	}
}

type diceService struct {
	diceRepository repository.DiceRepository
}

func (ds diceService) GetDiceRepository() repository.DiceRepository {
	return ds.diceRepository
}
