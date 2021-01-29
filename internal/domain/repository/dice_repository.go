package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type DiceRepository interface {
	FindByGameId(ctx context.Context, gameId uuid.UUID) (datamodel.Dices, error)
	InsertOrUpdate(ctx context.Context, dice *datamodel.Dice) error
}

func NewDiceRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate) DiceRepository {
	return &diceRepository{
		transactionMiddleware: transactionMiddleware,
		validate:              validate,
	}
}

type diceRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	validate              *validator.Validate
}

func (d diceRepository) FindByGameId(ctx context.Context, gameId uuid.UUID) (datamodel.Dices, error) {
	diceModels := make([]*model.Dice, 0)
	err := d.transactionMiddleware.Get(ctx).Find(&diceModels, "game_id = ?", gameId).Error
	if err != nil {
		return nil, errors.Wrap(err, "repository.DiceRepository.FindByGameId")
	}

	dices := make(datamodel.Dices, 0)
	for _, diceModel := range diceModels {
		dice := datamodel.NewDiceFromModel(diceModel)
		dices = append(dices, dice)
	}

	return dices, errors.Wrap(err, "repository.DiceRepository.FindByGameId")
}

func (d diceRepository) InsertOrUpdate(ctx context.Context, dice *datamodel.Dice) error {
	diceModel := dice.ToModel()

	if err := d.validate.StructCtx(ctx, diceModel); err != nil {
		return errors.Wrap(err, "repository.DiceRepository.InsertOrUpdate")
	}

	err := d.transactionMiddleware.Get(ctx).Save(diceModel).Error
	return errors.Wrap(err, "repository.DiceRepository.InsertOrUpdate")
}
