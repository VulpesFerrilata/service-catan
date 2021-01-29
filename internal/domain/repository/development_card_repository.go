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

type DevelopmentCardRepository interface {
	FindByGameId(ctx context.Context, gameId uuid.UUID) (datamodel.DevelopmentCards, error)
	InsertOrUpdate(ctx context.Context, developmentCard *datamodel.DevelopmentCard) error
}

func NewDevelopmentCardRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate) DevelopmentCardRepository {
	return &developmentCardRepository{
		transactionMiddleware: transactionMiddleware,
		validate:              validate,
	}
}

type developmentCardRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	validate              *validator.Validate
}

func (d developmentCardRepository) FindByGameId(ctx context.Context, gameId uuid.UUID) (datamodel.DevelopmentCards, error) {
	developmentCardModels := make([]*model.DevelopmentCard, 0)
	err := d.transactionMiddleware.Get(ctx).Find(&developmentCardModels, "game_id = ?", gameId).Error
	if err != nil {
		return nil, errors.Wrap(err, "repository.DevelopmentCardRepository.FindByGameId")
	}

	developmentCards := make(datamodel.DevelopmentCards, 0)
	for _, developmentCardModel := range developmentCardModels {
		developmentCard := datamodel.NewDevelopmentCardFromModel(developmentCardModel)
		developmentCards = append(developmentCards, developmentCard)
	}

	return developmentCards, nil
}

func (d developmentCardRepository) InsertOrUpdate(ctx context.Context, developmentCard *datamodel.DevelopmentCard) error {
	developmentCardModel := developmentCard.ToModel()

	if err := d.validate.StructCtx(ctx, developmentCardModel); err != nil {
		return errors.Wrap(err, "repository.DevelopmentCardRepository.InsertOrUpdate")
	}

	err := d.transactionMiddleware.Get(ctx).Save(developmentCardModel).Error
	return errors.Wrap(err, "repository.DevelopmentCardRepository.InsertOrUpdate")
}
