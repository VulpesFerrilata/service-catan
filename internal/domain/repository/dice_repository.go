package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/db"
)

type SafeDiceRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Dices, error)
}

type DiceRepository interface {
	SafeDiceRepository
	InsertOrUpdate(ctx context.Context, dice *model.Dice) error
}

func NewDiceRepository(dbContext *db.DbContext) DiceRepository {
	return &diceRepository{
		dbContext: dbContext,
	}
}

type diceRepository struct {
	dbContext *db.DbContext
}

func (dr *diceRepository) FindByGameId(ctx context.Context, gameId uint) (model.Dices, error) {
	var dices model.Dices
	return dices, dr.dbContext.GetDB(ctx).Find(&dices, "game_id = ?", gameId).Error
}

func (dr *diceRepository) InsertOrUpdate(ctx context.Context, dice *model.Dice) error {
	return dr.dbContext.GetDB(ctx).Save(dice).Error
}
