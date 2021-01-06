package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/app_error"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type SafeHarborRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (datamodel.Harbors, error)
}

type HarborRepository interface {
	SafeHarborRepository
	InsertOrUpdate(ctx context.Context, harbor *datamodel.Harbor) error
	Delete(ctx context.Context, harbor *datamodel.Harbor) error
}

func NewHarborRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate) HarborRepository {
	return &harborRepository{
		transactionMiddleware: transactionMiddleware,
		validate:              validate,
	}
}

type harborRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	validate              *validator.Validate
}

func (hr harborRepository) FindByGameId(ctx context.Context, gameId uint) (datamodel.Harbors, error) {
	harborModels := make([]*model.Harbor, 0)
	err := hr.transactionMiddleware.Get(ctx).Find(&harborModels, "game_id = ?", gameId).Error
	return datamodel.NewHarborsFromHarborModels(harborModels), errors.Wrap(err, "repository.HarborRepository.FindByGameId")
}

func (hr harborRepository) InsertOrUpdate(ctx context.Context, harbor *datamodel.Harbor) error {
	return harbor.Persist(func(harborModel *model.Harbor) error {
		if err := hr.validate.StructCtx(ctx, harborModel); err != nil {
			if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
				err = app_error.NewEntityValidationError(harborModel, fieldErrors)
			}
			return errors.Wrap(err, "repository.HarborRepository.InsertOrUpdate")
		}

		err := hr.transactionMiddleware.Get(ctx).Save(harborModel).Error
		return errors.Wrap(err, "repository.HarborRepository.InsertOrUpdate")
	})
}

func (hr harborRepository) Delete(ctx context.Context, harbor *datamodel.Harbor) error {
	return harbor.Persist(func(harborModel *model.Harbor) error {
		err := hr.transactionMiddleware.Get(ctx).Delete(harborModel).Error
		return errors.Wrap(err, "repository.HarborRepository.Delete")
	})
}
