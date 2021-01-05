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

type SafeConstructionRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (datamodel.Constructions, error)
}

type ConstructionRepository interface {
	SafeConstructionRepository
	InsertOrUpdate(ctx context.Context, construction *datamodel.Construction) error
}

func NewConstructionRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate) ConstructionRepository {
	return &constructionRepository{
		transactionMiddleware: transactionMiddleware,
		validate:              validate,
	}
}

type constructionRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	validate              *validator.Validate
}

func (cr constructionRepository) FindByGameId(ctx context.Context, gameId uint) (datamodel.Constructions, error) {
	constructionModels := make([]*model.Construction, 0)
	err := cr.transactionMiddleware.Get(ctx).Find(&constructionModels, "game_id = ?", gameId).Error
	return datamodel.NewConstructionsFromConstructionModels(constructionModels), errors.Wrap(err, "repository.ConstructionRepository.FindByGameId")
}

func (cr constructionRepository) InsertOrUpdate(ctx context.Context, construction *datamodel.Construction) error {
	return construction.Persist(func(constructionModel *model.Construction) error {
		if err := cr.validate.StructCtx(ctx, constructionModel); err != nil {
			if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
				err = app_error.NewEntityValidationError(constructionModel, fieldErrors)
			}
			return errors.Wrap(err, "repository.ConstructionRepository.InsertOrUpdate")
		}

		err := cr.transactionMiddleware.Get(ctx).Save(constructionModel).Error
		return errors.Wrap(err, "repository.ConstructionRepository.InsertOrUpdate")
	})
}
