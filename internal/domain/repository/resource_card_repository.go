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

type SafeResourceCardRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (datamodel.ResourceCards, error)
}

type ResourceCardRepository interface {
	SafeResourceCardRepository
	InsertOrUpdate(ctx context.Context, resourceCard *datamodel.ResourceCard) error
	Delete(ctx context.Context, resourceCard *datamodel.ResourceCard) error
}

func NewResourceCardRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate) ResourceCardRepository {
	return &resourceCardRepository{
		transactionMiddleware: transactionMiddleware,
		validate:              validate,
	}
}

type resourceCardRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	validate              *validator.Validate
}

func (rcr resourceCardRepository) FindByGameId(ctx context.Context, gameId uint) (datamodel.ResourceCards, error) {
	resourceCardModels := make([]*model.ResourceCard, 0)
	err := rcr.transactionMiddleware.Get(ctx).Find(&resourceCardModels, "game_id = ?", gameId).Error
	return datamodel.NewResourceCardsFromResourceCardModels(resourceCardModels), errors.Wrap(err, "repository.ResourceCardRepository.FindByGameId")
}

func (rcr resourceCardRepository) InsertOrUpdate(ctx context.Context, resourceCard *datamodel.ResourceCard) error {
	return resourceCard.Persist(func(resourceCardModel *model.ResourceCard) error {
		if err := rcr.validate.StructCtx(ctx, resourceCardModel); err != nil {
			if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
				err = app_error.NewEntityValidationError(resourceCardModel, fieldErrors)
			}
			return errors.Wrap(err, "repository.ResourceCardRepository.InsertOrUpdate")
		}
		err := rcr.transactionMiddleware.Get(ctx).Save(resourceCardModel).Error
		return errors.Wrap(err, "repository.ResourceCardRepository.InsertOrUpdate")
	})
}

func (rcr resourceCardRepository) Delete(ctx context.Context, resourceCard *datamodel.ResourceCard) error {
	return resourceCard.Persist(func(resourceCardModel *model.ResourceCard) error {
		err := rcr.transactionMiddleware.Get(ctx).Delete(resourceCardModel).Error
		return errors.Wrap(err, "repository.ResourceCardRepository.Delete")
	})
}
