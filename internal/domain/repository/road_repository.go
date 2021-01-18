package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type RoadRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (datamodel.Roads, error)
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
	roadModel := road.ToModel()

	if err := rr.validate.StructCtx(ctx, roadModel); err != nil {
		return errors.Wrap(err, "repository.RoadRepository.insertOrUpdate")
	}

	err := rr.transactionMiddleware.Get(ctx).Save(roadModel).Error
	return errors.Wrap(err, "repository.RoadRepository.insertOrUpdate")
}

func (rr roadRepository) delete(ctx context.Context, road *datamodel.Road) error {
	roadModel := road.ToModel()
	err := rr.transactionMiddleware.Get(ctx).Delete(roadModel).Error
	return errors.Wrap(err, "repository.RoadRepository.delete")
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
