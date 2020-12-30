package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
)

type SafeDiceRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Dices, error)
}

type DiceRepository interface {
	SafeDiceRepository
	InsertOrUpdate(ctx context.Context, dice *model.Dice) error
}

func NewDiceRepository(transactionMiddleware *middleware.TransactionMiddleware) DiceRepository {
	return &diceRepository{
		transactionMiddleware: transactionMiddleware,
	}
}

type diceRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
}

func (dr *diceRepository) FindByGameId(ctx context.Context, gameId uint) (model.Dices, error) {
	var dices model.Dices
	return dices, dr.transactionMiddleware.Get(ctx).Find(&dices, "game_id = ?", gameId).Error
}

func (dr *diceRepository) InsertOrUpdate(ctx context.Context, dice *model.Dice) error {
	return dr.transactionMiddleware.Get(ctx).Save(dice).Error
}
