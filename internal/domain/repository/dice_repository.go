package repository

import (
	"context"

	model "github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

type ReadOnlyDiceRepository interface {
	FindByGameId(ctx context.Context, gameId uint) ([]*model.Dice, error)
}

type DiceRepository interface {
	ReadOnlyDiceRepository
	Insert(ctx context.Context, dices ...*model.Dice) error
	Save(ctx context.Context, dices ...*model.Dice) error
}
