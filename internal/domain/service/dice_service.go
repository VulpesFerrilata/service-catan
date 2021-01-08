package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type DiceService interface {
	GetDiceRepository() repository.SafeDiceRepository
	Save(ctx context.Context, dice *datamodel.Dice) error
}

func NewDiceService(diceRepository repository.DiceRepository) DiceService {
	return &diceService{
		diceRepository: diceRepository,
	}
}

type diceService struct {
	diceRepository repository.DiceRepository
}

func (ds diceService) GetDiceRepository() repository.SafeDiceRepository {
	return ds.diceRepository
}
