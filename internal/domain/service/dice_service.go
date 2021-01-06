package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
	"github.com/pkg/errors"
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

func (ds diceService) Save(ctx context.Context, dice *datamodel.Dice) error {
	if dice.IsRemoved() {
		err := ds.diceRepository.Delete(ctx, dice)
		return errors.Wrap(err, "service.DiceService.Save")
	}
	if dice.IsModified() {
		err := ds.diceRepository.InsertOrUpdate(ctx, dice)
		return errors.Wrap(err, "service.DiceService.Save")
	}
	return nil
}
