package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type ConstructionRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (datamodel.Constructions, error)
	Save(ctx context.Context, construction *datamodel.Construction) error
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

func (cr constructionRepository) insertOrUpdate(ctx context.Context, construction *datamodel.Construction) error {
	constructionModel := construction.ToModel()

	if err := cr.validate.StructCtx(ctx, constructionModel); err != nil {
		return errors.Wrap(err, "repository.ConstructionRepository.insertOrUpdate")
	}

	err := cr.transactionMiddleware.Get(ctx).Save(constructionModel).Error
	return errors.Wrap(err, "repository.ConstructionRepository.insertOrUpdate")
}

func (cr constructionRepository) delete(ctx context.Context, construction *datamodel.Construction) error {
	constructionModel := construction.ToModel()
	err := cr.transactionMiddleware.Get(ctx).Delete(constructionModel).Error
	return errors.Wrap(err, "repository.ConstructionRepository.delete")
}

func (cr constructionRepository) Save(ctx context.Context, construction *datamodel.Construction) error {
	if construction.IsRemoved() {
		err := cr.delete(ctx, construction)
		return errors.Wrap(err, "service.ConstructionRepository.Save")
	}
	if construction.IsModified() {
		err := cr.insertOrUpdate(ctx, construction)
		return errors.Wrap(err, "service.ConstructionRepository.Save")
	}
	return nil
}
