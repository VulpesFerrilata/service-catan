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

type HexRepository interface {
	GetById(ctx context.Context, id uuid.UUID) (*datamodel.Hex, error)
	InsertOrUpdate(ctx context.Context, hex *datamodel.Hex) error
}

func NewHexRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate) HexRepository {
	return &hexRepository{
		transactionMiddleware: transactionMiddleware,
		validate:              validate,
	}
}

type hexRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	validate              *validator.Validate
}

func (h hexRepository) GetById(ctx context.Context, id uuid.UUID) (*datamodel.Hex, error) {
	hexModel := new(model.Hex)
	err := h.transactionMiddleware.Get(ctx).First(hexModel, id).Error
	return datamodel.NewHexFromModel(hexModel), errors.Wrap(err, "repository.HexRepository.GetByID")
}

func (h hexRepository) InsertOrUpdate(ctx context.Context, hex *datamodel.Hex) error {
	hexModel := hex.ToModel()

	if err := h.validate.StructCtx(ctx, hexModel); err != nil {
		return errors.Wrap(err, "repository.HexRepository.InsertOrUpdate")
	}

	err := h.transactionMiddleware.Get(ctx).Save(hexModel).Error
	return errors.Wrap(err, "repository.HexRepository.InsertOrUpdate")
}
