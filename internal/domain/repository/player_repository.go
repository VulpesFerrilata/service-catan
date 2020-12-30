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
	FindByGameId(ctx context.Context, gameId uint) (model.Players, error)
}

type PlayerRepository interface {
	SafePlayerRepository
	Save(ctx context.Context, player *model.Player) error
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

func (pr *playerRepository) FindByGameId(ctx context.Context, gameId uint) (model.Players, error) {
	playerDMs := make([]*datamodel.Player, 0)
	err := pr.transactionMiddleware.Get(ctx).Find(&playerDMs, "game_id = ?", gameId).Error
	if err != nil {
		return nil, errors.Wrap(err, "repository.PlayerRepository.FindByGameId")
	}
	var players model.Players

	for _, playerDM := range playerDMs {
		player := model.EmptyPlayer()
		player.Persist(func(player *datamodel.Player) error {
			player = playerDM
			return nil
		})

		userRequest := new(user.UserRequest)
		userRequest.ID = int64(playerDM.UserID)
		userPb, err := pr.userService.GetUserById(ctx, userRequest)
		if err != nil {
			return nil, errors.Wrap(err, "repository.PlayerRepository.FindByGameId")
		}
		user := model.NewUser(userPb)
		player.SetUser(user)

		players = append(players, player)
	}

	return players, nil
}

func (pr *playerRepository) Save(ctx context.Context, player *model.Player) error {
	if player.IsRemoved() {
		err := pr.delete(ctx, player)
		return errors.Wrap(err, "repository.PlayerRepository.Save")
	}
	if player.IsModified() {
		err := pr.insertOrUpdate(ctx, player)
		return errors.Wrap(err, "repository.PlayerRepository.Save")
	}
	return nil
}

func (pr *playerRepository) insertOrUpdate(ctx context.Context, player *model.Player) error {
	return player.Persist(func(player *datamodel.Player) error {
		if err := pr.validate.StructCtx(ctx, player); err != nil {
			if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
				err = app_error.NewValidationError(app_error.EntityValidation, "player", fieldErrors)
			}
			return errors.Wrap(err, "repository.UserRepository.Insert")
		}
		err := pr.transactionMiddleware.Get(ctx).Save(player).Error
		return errors.Wrap(err, "repository.PlayerRepository.InsertOrUpdate")
	})
}

func (pr *playerRepository) delete(ctx context.Context, player *model.Player) error {
	return player.Persist(func(player *datamodel.Player) error {
		err := pr.transactionMiddleware.Get(ctx).Delete(player).Error
		return errors.Wrap(err, "repository.PlayerRepository.Delete")
	})
}
