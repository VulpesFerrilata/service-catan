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

type SafeRoadRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (datamodel.Roads, error)
}

type RoadRepository interface {
	SafeRoadRepository
	Save(ctx context.Context, road *datamodel.Road) error
}

func NewRoadRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate) RoadRepository {
	return &roadRepository{
		transactionMiddleware: transactionMiddleware,
		validate:              validate,
	}
}

type roadRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	validate              *validator.Validate
}

func (rr roadRepository) FindByGameId(ctx context.Context, gameId uint) (datamodel.Roads, error) {
	roadModels := make([]*model.Road, 0)
	err := rr.transactionMiddleware.Get(ctx).Find(&roadModels, "game_id = ?", gameId).Error
	return datamodel.NewRoadsFromRoadModels(roadModels), errors.Wrap(err, "repository.RoadRepository.FindByGameId")
}

func (rr roadRepository) insertOrUpdate(ctx context.Context, road *datamodel.Road) error {
	return road.Persist(func(roadModel *model.Road) error {
		if err := rr.validate.StructCtx(ctx, roadModel); err != nil {
			if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
				err = app_error.NewEntityValidationError(roadModel, fieldErrors)
			}
			return errors.Wrap(err, "repository.RoadRepository.InsertOrUpdate")
		}
		err := rr.transactionMiddleware.Get(ctx).Save(roadModel).Error
		return errors.Wrap(err, "repository.RoadRepository.InsertOrUpdate")
	})
}

func (rr roadRepository) delete(ctx context.Context, road *datamodel.Road) error {
	return road.Persist(func(roadModel *model.Road) error {
		err := rr.transactionMiddleware.Get(ctx).Delete(roadModel).Error
		return errors.Wrap(err, "repository.RoadRepository.Delete")
	})
}

func (rr roadRepository) Save(ctx context.Context, road *datamodel.Road) error {
	if road.IsRemoved() {
		err := rr.delete(ctx, road)
		return errors.Wrap(err, "service.RoadRepository.Save")
	}
	if road.IsModified() {
		err := rr.insertOrUpdate(ctx, road)
		return errors.Wrap(err, "service.RoadRepository.Save")
	}
	return nil
}
