package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type HarborRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (datamodel.Harbors, error)
	InsertOrUpdate(ctx context.Context, harbor *datamodel.Harbor) error
}

func NewHarborRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate,
	hexRepository HexRepository) HarborRepository {
	return &harborRepository{
		transactionMiddleware: transactionMiddleware,
		validate:              validate,
		hexRepository:         hexRepository,
	}
}

type harborRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	validate              *validator.Validate
	hexRepository         HexRepository
}

func (h harborRepository) FindByGameId(ctx context.Context, gameId uint) (datamodel.Harbors, error) {
	harborModels := make([]*model.Harbor, 0)
	err := h.transactionMiddleware.Get(ctx).Find(&harborModels, "game_id = ?", gameId).Error
	if err != nil {
		return nil, errors.Wrap(err, "repository.HarborRepository.FindByGameId")
	}

	harbors := make(datamodel.Harbors, 0)
	for _, harborModel := range harborModels {
		harbor, err := datamodel.NewHarborFromHarborModel(harborModel)
		if err != nil {
			return nil, errors.Wrap(err, "repository.HarborRepository.FindByGameId")
		}

		hex, err := h.hexRepository.GetById(ctx, harborModel.HexID)
		if err != nil {
			return nil, errors.Wrap(err, "repository.HarborRepository.FindByGameId")
		}
		harbor.SetHex(hex)

		harbors = append(harbors, harbor)
	}

	return harbors, nil
}

func (h harborRepository) InsertOrUpdate(ctx context.Context, harbor *datamodel.Harbor) error {
	harborModel := harbor.ToModel()

	if err := h.validate.StructCtx(ctx, harborModel); err != nil {
		return errors.Wrap(err, "repository.HarborRepository.InsertOrUpdate")
	}

	if err := h.transactionMiddleware.Get(ctx).Save(harborModel).Error; err != nil {
		return errors.Wrap(err, "repository.HarborRepository.InsertOrUpdate")
	}

	hex := harbor.GetHex()
	err := h.hexRepository.InsertOrUpdate(ctx, hex)
	return errors.Wrap(err, "repository.HarborRepository.InsertOrUpdate")
}
