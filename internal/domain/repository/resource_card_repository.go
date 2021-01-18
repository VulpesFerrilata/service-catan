package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type ResourceCardRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (datamodel.ResourceCards, error)
	Save(ctx context.Context, resourceCard *datamodel.ResourceCard) error
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

func (rcr resourceCardRepository) insertOrUpdate(ctx context.Context, resourceCard *datamodel.ResourceCard) error {
	resourceCardModel := resourceCard.ToModel()

	if err := rcr.validate.StructCtx(ctx, resourceCardModel); err != nil {
		return errors.Wrap(err, "repository.ResourceCardRepository.insertOrUpdate")
	}

	err := rcr.transactionMiddleware.Get(ctx).Save(resourceCardModel).Error
	return errors.Wrap(err, "repository.ResourceCardRepository.insertOrUpdate")
}

func (rcr resourceCardRepository) delete(ctx context.Context, resourceCard *datamodel.ResourceCard) error {
	resourceCardModel := resourceCard.ToModel()
	err := rcr.transactionMiddleware.Get(ctx).Delete(resourceCardModel).Error
	return errors.Wrap(err, "repository.ResourceCardRepository.delete")
}

func (rcr resourceCardRepository) Save(ctx context.Context, resourceCard *datamodel.ResourceCard) error {
	if resourceCard.IsRemoved() {
		err := rcr.delete(ctx, resourceCard)
		return errors.Wrap(err, "service.ResourceCardRepository.Save")
	}
	if resourceCard.IsModified() {
		err := rcr.insertOrUpdate(ctx, resourceCard)
		return errors.Wrap(err, "service.ResourceCardRepository.Save")
	}
	return nil
}
