package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type ConstructionRepository interface {
	FindByGameId(ctx context.Context, gameId uuid.UUID) (datamodel.Constructions, error)
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

func (c constructionRepository) FindByGameId(ctx context.Context, gameId uuid.UUID) (datamodel.Constructions, error) {
	constructionModels := make([]*model.Construction, 0)

	err := c.transactionMiddleware.Get(ctx).Find(&constructionModels, "game_id = ?", gameId).Error
	if err != nil {
		return nil, errors.Wrap(err, "repository.ConstructionRepository.FindByGameId")
	}

	constructions := make(datamodel.Constructions, 0)
	for _, constructionModel := range constructionModels {
		construction, err := datamodel.NewConstructionFromModel(constructionModel)
		if err != nil {
			return nil, errors.Wrap(err, "repository.ConstructionRepository.FindByGameId")
		}

		constructions = append(constructions, construction)
	}

	return constructions, nil
}

func (c constructionRepository) InsertOrUpdate(ctx context.Context, construction *datamodel.Construction) error {
	constructionModel := construction.ToModel()

	if err := c.validate.StructCtx(ctx, constructionModel); err != nil {
		return errors.Wrap(err, "repository.ConstructionRepository.InsertOrUpdate")
	}

	err := c.transactionMiddleware.Get(ctx).Save(constructionModel).Error
	return errors.Wrap(err, "repository.ConstructionRepository.InsertOrUpdate")
}
