package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/grpc/protoc/user"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type PlayerRepository interface {
	FindByGameId(ctx context.Context, gameId uuid.UUID) (datamodel.Players, error)
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

func (pr playerRepository) FindByGameId(ctx context.Context, gameId uuid.UUID) (datamodel.Players, error) {
	playerModels := make([]*model.Player, 0)
	err := pr.transactionMiddleware.Get(ctx).Find(&playerModels, "game_id = ?", gameId).Error
	if err != nil {
		return nil, errors.Wrap(err, "repository.PlayerRepository.FindByGameId")
	}

	players := make(datamodel.Players, 0)
	for _, playerModel := range playerModels {
		player := datamodel.NewPlayerFromPlayerModel(playerModel)

		userRequestPb := new(user.UserRequest)
		userRequestPb.ID = playerModel.UserID.String()
		userResponsePb, err := pr.userService.GetUserById(ctx, userRequestPb)
		if err != nil {
			return nil, errors.Wrap(err, "repository.PlayerRepository.FindByGameId")
		}

		user, err := datamodel.NewUserFromUserPb(userResponsePb)
		if err != nil {
			return nil, errors.Wrap(err, "repository.PlayerRepository.FindByGameId")
		}
		player.SetUser(user)

		players = append(players, player)
	}

	return datamodel.NewPlayersFromPlayerModels(playerModels), errors.Wrap(err, "repository.PlayerRepository.FindByGameId")
}

func (pr playerRepository) insertOrUpdate(ctx context.Context, player *datamodel.Player) error {
	playerModel := player.ToModel()

	if err := pr.validate.StructCtx(ctx, playerModel); err != nil {
		return errors.Wrap(err, "repository.PlayerRepository.insertOrUpdate")
	}

	err := pr.transactionMiddleware.Get(ctx).Save(playerModel).Error
	return errors.Wrap(err, "repository.PlayerRepository.insertOrUpdate")
}

func (pr playerRepository) delete(ctx context.Context, player *datamodel.Player) error {
	playerModel := player.ToModel()
	err := pr.transactionMiddleware.Get(ctx).Delete(playerModel).Error
	return errors.Wrap(err, "repository.PlayerRepository.delete")
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
