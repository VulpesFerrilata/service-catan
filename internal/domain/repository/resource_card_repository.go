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
	InsertOrUpdate(ctx context.Context, resourceCard *datamodel.ResourceCard) error
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

func (r resourceCardRepository) FindByGameId(ctx context.Context, gameId uint) (datamodel.ResourceCards, error) {
	resourceCardModels := make([]*model.ResourceCard, 0)
	err := r.transactionMiddleware.Get(ctx).Find(&resourceCardModels, "game_id = ?", gameId).Error
	return datamodel.NewResourceCardsFromResourceCardModels(resourceCardModels), errors.Wrap(err, "repository.ResourceCardRepository.FindByGameId")
}

func (r resourceCardRepository) InsertOrUpdate(ctx context.Context, resourceCard *datamodel.ResourceCard) error {
	resourceCardModel := resourceCard.ToModel()

	if err := r.validate.StructCtx(ctx, resourceCardModel); err != nil {
		return errors.Wrap(err, "repository.ResourceCardRepository.InsertOrUpdate")
	}

	err := r.transactionMiddleware.Get(ctx).Save(resourceCardModel).Error
	return errors.Wrap(err, "repository.ResourceCardRepository.InsertOrUpdate")
}
