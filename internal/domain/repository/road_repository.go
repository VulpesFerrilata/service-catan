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
	InsertOrUpdate(ctx context.Context, road *datamodel.Road) error
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

func (r roadRepository) FindByGameId(ctx context.Context, gameId uint) (datamodel.Roads, error) {
	roadModels := make([]*model.Road, 0)
	err := r.transactionMiddleware.Get(ctx).Find(&roadModels, "game_id = ?", gameId).Error
	if err != nil {
		return nil, errors.Wrap(err, "repository.RoadRepository.FindByGameId")
	}
	roads, err := datamodel.NewRoadsFromRoadModels(roadModels)
	return roads, errors.Wrap(err, "repository.RoadRepository.FindByGameId")
}

func (r roadRepository) InsertOrUpdate(ctx context.Context, road *datamodel.Road) error {
	roadModel := road.ToModel()

	if err := r.validate.StructCtx(ctx, roadModel); err != nil {
		return errors.Wrap(err, "repository.RoadRepository.InsertOrUpdate")
	}

	err := r.transactionMiddleware.Get(ctx).Save(roadModel).Error
	return errors.Wrap(err, "repository.RoadRepository.InsertOrUpdate")
}
