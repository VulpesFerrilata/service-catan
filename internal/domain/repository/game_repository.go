package repository

import (
	"context"

	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/library/pkg/app_error"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"gorm.io/gorm"
)

type SafeGameRepository interface {
	GetById(ctx context.Context, id uint) (*datamodel.Game, error)
}

type GameRepository interface {
	SafeGameRepository
	InsertOrUpdate(ctx context.Context, game *datamodel.Game) error
	Delete(ctx context.Context, game *datamodel.Game) error
}

func NewGameRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate,
	playerRepository PlayerRepository) GameRepository {
	return &gameRepository{
		transactionMiddleware: transactionMiddleware,
		validate:              validate,
		playerRepository:      playerRepository,
	}
}

type gameRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	validate              *validator.Validate
	playerRepository      PlayerRepository
}

func (gr gameRepository) GetById(ctx context.Context, gameId uint) (*datamodel.Game, error) {
	gameModel := new(model.Game)
	err := gr.transactionMiddleware.Get(ctx).First(&gameModel, gameId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, app_error.NewNotFoundError("game")
	}
	if err != nil {
		return nil, errors.Wrap(err, "repository.GameRepository.GetById")
	}

	game := datamodel.NewGameFromGameModel(gameModel)

	players, err := gr.playerRepository.FindByGameId(ctx, gameModel.ID)
	if err != nil {
		return nil, errors.Wrap(err, "repository.GameRepository.GetById")
	}
	game.AddPlayers(players...)

	return game, nil
}

func (gr gameRepository) Save(ctx context.Context, game *datamodel.Game) error {
	err := gr.save(ctx, game)
	if err != nil {
		return errors.Wrap(err, "repository.GameRepository.Save")
	}

	for _, player := range game.GetPlayers() {
		err := gr.playerRepository.Save(ctx, player)
		if err != nil {
			return errors.Wrap(err, "repository.GameRepository.Save")
		}
	}
	return nil
}

func (gr gameRepository) save(ctx context.Context, game *datamodel.Game) error {
	if game.IsRemoved() {
		err := gr.Delete(ctx, game)
		return errors.Wrap(err, "repository.GameRepository.save")
	}
	if game.IsModified() {
		err := gr.InsertOrUpdate(ctx, game)
		return errors.Wrap(err, "repository.GameRepository.save")
	}
	return nil
}

func (gr *gameRepository) InsertOrUpdate(ctx context.Context, game *datamodel.Game) error {
	return game.Persist(func(gameModel *model.Game) error {
		if err := gr.validate.StructCtx(ctx, gameModel); err != nil {
			if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
				err = app_error.NewEntityValidationError(gameModel, fieldErrors)
			}
			return errors.Wrap(err, "repository.GameRepository.InsertOrUpdate")
		}

		err := gr.transactionMiddleware.Get(ctx).Save(gameModel).Error
		return errors.Wrap(err, "repository.GameRepository.InsertOrUpdate")
	})
}

func (gr *gameRepository) Delete(ctx context.Context, game *datamodel.Game) error {
	return game.Persist(func(gameModel *model.Game) error {
		err := gr.transactionMiddleware.Get(ctx).Delete(gameModel).Error
		return errors.Wrap(err, "repository.GameRepository.Delete")
	})
}
