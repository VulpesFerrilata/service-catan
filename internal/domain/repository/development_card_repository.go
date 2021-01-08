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

type SafeDevelopmentCardRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (datamodel.DevelopmentCards, error)
}

type DevelopmentCardRepository interface {
	SafeDevelopmentCardRepository
	Save(ctx context.Context, developmentCard *datamodel.DevelopmentCard) error
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

func (dcr developmentCardRepository) FindByGameId(ctx context.Context, gameId uint) (datamodel.DevelopmentCards, error) {
	developmentCardModels := make([]*model.DevelopmentCard, 0)
	err := dcr.transactionMiddleware.Get(ctx).Find(&developmentCardModels, "game_id = ?", gameId).Error
	return datamodel.NewDevelopmentCardsFromDevelopmentCardModels(developmentCardModels), errors.Wrap(err, "repository.DevelopmentCardRepository.FindByGameId")
}

func (dcr developmentCardRepository) insertOrUpdate(ctx context.Context, developmentCard *datamodel.DevelopmentCard) error {
	return developmentCard.Persist(func(developmentCardModel *model.DevelopmentCard) error {
		if err := dcr.validate.StructCtx(ctx, developmentCardModel); err != nil {
			if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
				err = app_error.NewEntityValidationError(developmentCardModel, fieldErrors)
			}
			return errors.Wrap(err, "repository.DevelopmentCardRepository.InsertOrUpdate")
		}

		err := dcr.transactionMiddleware.Get(ctx).Save(developmentCardModel).Error
		return errors.Wrap(err, "repository.DevelopmentCardRepository.InsertOrUpdate")
	})
}

func (dcr developmentCardRepository) delete(ctx context.Context, developmentCard *datamodel.DevelopmentCard) error {
	return developmentCard.Persist(func(developmentCardModel *model.DevelopmentCard) error {
		err := dcr.transactionMiddleware.Get(ctx).Delete(developmentCardModel).Error
		return errors.Wrap(err, "repository.DevelopmentCardRepository.Delete")
	})
}

func (dcr developmentCardRepository) Save(ctx context.Context, developmentCard *datamodel.DevelopmentCard) error {
	if developmentCard.IsRemoved() {
		err := dcr.delete(ctx, developmentCard)
		return errors.Wrap(err, "service.DevelopmentCardRepository.Save")
	}
	if developmentCard.IsModified() {
		err := dcr.insertOrUpdate(ctx, developmentCard)
		return errors.Wrap(err, "service.DevelopmentCardRepository.Save")
	}
	return nil
}
