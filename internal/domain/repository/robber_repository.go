package repository

import (
	"context"

	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/app_error"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"gorm.io/gorm"
)

type SafeRobberRepository interface {
	GetByGameId(ctx context.Context, gameId uint) (*datamodel.Robber, error)
}

type RobberRepository interface {
	SafeRobberRepository
	InsertOrUpdate(ctx context.Context, robber *datamodel.Robber) error
	Delete(ctx context.Context, robber *datamodel.Robber) error
}

func NewRobberRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate) RobberRepository {
	return &robberRepository{
		transactionMiddleware: transactionMiddleware,
		validate:              validate,
	}
}

type robberRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	validate              *validator.Validate
}

func (rr robberRepository) GetByGameId(ctx context.Context, gameId uint) (*datamodel.Robber, error) {
	robberModel := new(model.Robber)
	err := rr.transactionMiddleware.Get(ctx).First(&robberModel, "game_id = ?", gameId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, app_error.NewNotFoundError("game")
	}
	return datamodel.NewRobberFromRobberModel(robberModel), errors.Wrap(err, "repository.RobberRepository.GetByGameId")
}

func (rr robberRepository) InsertOrUpdate(ctx context.Context, robber *datamodel.Robber) error {
	return robber.Persist(func(robberModel *model.Robber) error {
		if err := rr.validate.StructCtx(ctx, robberModel); err != nil {
			if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
				err = app_error.NewEntityValidationError(robberModel, fieldErrors)
			}
			return errors.Wrap(err, "repository.RobberRepository.InsertOrUpdate")
		}
		err := rr.transactionMiddleware.Get(ctx).Save(robberModel).Error
		return errors.Wrap(err, "repository.RobberRepository.InsertOrUpdate")
	})
}

func (rr robberRepository) Delete(ctx context.Context, robber *datamodel.Robber) error {
	return robber.Persist(func(robberModel *model.Robber) error {
		err := rr.transactionMiddleware.Get(ctx).Delete(robberModel).Error
		return errors.Wrap(err, "repository.RobberRepository.Delete")
	})
}
