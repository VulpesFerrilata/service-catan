package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type HarborRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (datamodel.Harbors, error)
	Save(ctx context.Context, harbor *datamodel.Harbor) error
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

func (hr harborRepository) insertOrUpdate(ctx context.Context, harbor *datamodel.Harbor) error {
	return harbor.Persist(func(harborModel *model.Harbor) error {
		if err := hr.validate.StructCtx(ctx, harborModel); err != nil {
			return errors.Wrap(err, "repository.HarborRepository.InsertOrUpdate")
		}

		err := hr.transactionMiddleware.Get(ctx).Save(harborModel).Error
		return errors.Wrap(err, "repository.HarborRepository.InsertOrUpdate")
	})
}

func (hr harborRepository) delete(ctx context.Context, harbor *datamodel.Harbor) error {
	return harbor.Persist(func(harborModel *model.Harbor) error {
		err := hr.transactionMiddleware.Get(ctx).Delete(harborModel).Error
		return errors.Wrap(err, "repository.HarborRepository.Delete")
	})
}

func (hr harborRepository) Save(ctx context.Context, harbor *datamodel.Harbor) error {
	if harbor.IsRemoved() {
		err := hr.delete(ctx, harbor)
		return errors.Wrap(err, "service.HarborRepository.Save")
	}
	if harbor.IsModified() {
		err := hr.insertOrUpdate(ctx, harbor)
		return errors.Wrap(err, "service.HarborRepository.Save")
	}
	return nil
}
