package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type DiceRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (datamodel.Dices, error)
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

func (dr diceRepository) FindByGameId(ctx context.Context, gameId uint) (datamodel.Dices, error) {
	diceModels := make([]*model.Dice, 0)
	err := dr.transactionMiddleware.Get(ctx).Find(&diceModels, "game_id = ?", gameId).Error
	return datamodel.NewDicesFromDiceModels(diceModels), errors.Wrap(err, "repository.DiceRepository.FindByGameId")
}

func (dr diceRepository) InsertOrUpdate(ctx context.Context, dice *datamodel.Dice) error {
	diceModel := dice.ToModel()

	if err := dr.validate.StructCtx(ctx, diceModel); err != nil {
		return errors.Wrap(err, "repository.DiceRepository.InsertOrUpdate")
	}

	err := dr.transactionMiddleware.Get(ctx).Save(diceModel).Error
	return errors.Wrap(err, "repository.DiceRepository.InsertOrUpdate")
}
