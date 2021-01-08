package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/grpc/protoc/user"
	"github.com/VulpesFerrilata/library/pkg/app_error"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type SafePlayerRepository interface {
	FindByGameId(ctx context.Context, gameId int) (datamodel.Players, error)
}

type PlayerRepository interface {
	SafePlayerRepository
	Save(ctx context.Context, player *datamodel.Player) error
}

func NewPlayerRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate,
	userService user.UserService) PlayerRepository {
	return &playerRepository{
		transactionMiddleware: transactionMiddleware,
		validate:              validate,
		userService:           userService,
	}
}

type playerRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	validate              *validator.Validate
	userService           user.UserService
}

func (pr playerRepository) FindByGameId(ctx context.Context, gameId int) (datamodel.Players, error) {
	playerModels := make([]*model.Player, 0)
	err := pr.transactionMiddleware.Get(ctx).Find(&playerModels, "game_id = ?", gameId).Error
	if err != nil {
		return nil, errors.Wrap(err, "repository.PlayerRepository.FindByGameId")
	}

	players := make(datamodel.Players, 0)
	for _, playerModel := range playerModels {
		player := datamodel.NewPlayerFromPlayerModel(playerModel)

		userRequestPb := new(user.UserRequest)
		userRequestPb.ID = int64(playerModel.UserID)
		userResponsePb, err := pr.userService.GetUserById(ctx, userRequestPb)
		if err != nil {
			return nil, errors.Wrap(err, "repository.PlayerRepository.FindByGameId")
		}
		user := datamodel.NewUserFromUserPb(userResponsePb)
		player.SetUser(user)

		players = append(players, player)
	}

	return datamodel.NewPlayersFromPlayerModels(playerModels), errors.Wrap(err, "repository.PlayerRepository.FindByGameId")
}

func (pr playerRepository) insertOrUpdate(ctx context.Context, player *datamodel.Player) error {
	return player.Persist(func(playerModel *model.Player) error {
		if err := pr.validate.StructCtx(ctx, playerModel); err != nil {
			if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
				err = app_error.NewEntityValidationError(playerModel, fieldErrors)
			}
			return errors.Wrap(err, "repository.PlayerRepository.InsertOrUpdate")
		}
		err := pr.transactionMiddleware.Get(ctx).Save(playerModel).Error
		return errors.Wrap(err, "repository.PlayerRepository.InsertOrUpdate")
	})
}

func (pr playerRepository) delete(ctx context.Context, player *datamodel.Player) error {
	return player.Persist(func(playerModel *model.Player) error {
		err := pr.transactionMiddleware.Get(ctx).Delete(playerModel).Error
		return errors.Wrap(err, "repository.PlayerRepository.Delete")
	})
}

func (pr playerRepository) Save(ctx context.Context, player *datamodel.Player) error {
	if player.IsRemoved() {
		err := pr.delete(ctx, player)
		return errors.Wrap(err, "service.PlayerRepository.Save")
	}
	if player.IsModified() {
		err := pr.insertOrUpdate(ctx, player)
		return errors.Wrap(err, "service.PlayerRepository.Save")
	}
	return nil
}
