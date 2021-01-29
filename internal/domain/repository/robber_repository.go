package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/app_error"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"gorm.io/gorm"
)

type RobberRepository interface {
	GetByGameId(ctx context.Context, gameId uuid.UUID) (*datamodel.Robber, error)
	InsertOrUpdate(ctx context.Context, robber *datamodel.Robber) error
}

func NewRobberRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate) RobberRepository {
	return &robberRepository{
		transactionMiddleware: transactionMiddleware,
		validate:              validate,
	}
}

type robberRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	validate              *validator.Validate
}

func (r robberRepository) GetByGameId(ctx context.Context, gameId uuid.UUID) (*datamodel.Robber, error) {
	robberModel := new(model.Robber)
	err := r.transactionMiddleware.Get(ctx).First(&robberModel, "game_id = ?", gameId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, app_error.NewNotFoundError("robber")
	}
	if err != nil {
		return nil, errors.Wrap(err, "repository.RobberRepository.GetByGameId")
	}

	robber := datamodel.NewRobberFromRobberModel(robberModel)
	return robber, nil
}

func (r robberRepository) InsertOrUpdate(ctx context.Context, robber *datamodel.Robber) error {
	robberModel := robber.ToModel()

	if err := r.validate.StructCtx(ctx, robberModel); err != nil {
		return errors.Wrap(err, "repository.RobberRepository.InsertOrUpdate")
	}

	err := r.transactionMiddleware.Get(ctx).Save(robberModel).Error
	return errors.Wrap(err, "repository.RobberRepository.InsertOrUpdate")
}
